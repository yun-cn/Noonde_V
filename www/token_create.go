package www

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/yun313350095/Noonde/api"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type tokenCreateForm struct {
	service *Service
	context api.Context

	Email    string `json:"email"`
	Password string `json:"password"`
}

func (f *tokenCreateForm) parseAndValidate() error {
	s := f.service
	c := f.context
	var err error

	r := c.Request()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(f)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	// Start validation.
	perr := s.HTTP.NewParamError()
	tx := c.Tx()

	// Email is not empty?
	if f.Email == "" {
		perr.PushIfNotExists("email", api.ParamEmptyValue)
	}

	// Password is not empty?
	if f.Password == "" {
		perr.PushIfNotExists("password", api.ParamEmptyValue)
	}

	// User exists?
	user := &api.User{}
	err = tx.GetContext(c.Context(), user, `select * from users where email =?`, f.Email)
	if err != nil {
		s.Log.Error(err)
		return err
	}
	c.SetCurUser(user)

	// Valid password ?
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(f.Password))
	if err != nil {
		s.Log.Error(err)
		perr.PushIfNotExists("password", api.ParamInvalidFormat)
	}

	if perr.IsNotEmpty() {
		return perr
	}

	return nil
}

func (s *Service) tokenCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	// Setup context.
	c := s.HTTP.NewContext(r.Context())
	c.SetRequest(r)
	c.SetParams(p)

	// Start transaction.
	tx, err := s.MySQL.Begin(c.Context())
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	c.SetTx(tx)

	// Parse and validate.
	f := &tokenCreateForm{
		service: s,
		context: c,
	}
	perr := f.parseAndValidate()
	if perr != nil {
		s.Log.Error(perr)
		s.HTTP.WriteError(w, api.ErrInvalidRequest, perr)
		return
	}

	// Create jwt
	curuser := c.CurUser()
	exp := time.Now().Add(time.Hour * 24 * 30)
	clm := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		UserID: curuser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtIssuer,
			Subject:   jwtSubject,
		},
	})
	jwt, err := clm.SignedString([]byte(s.Conf.String("auth.key")))
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	// Get overflowed tokens.
	ii := []int64{}
	err = tx.SelectContext(c.Context(), &ii,
		`select id from user_tokens where user_id = ? order by expired desc limit ?`,
		curuser.ID,
		curuser.MaxTokens-1,
	)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	// Delete overflowed tokens
	if len(ii) > 0 {
		query, aa, err := sqlx.In(`delete from user_tokens where user_id = ? and (id not in (?) or expired<now())`,
			curuser.ID, ii)
		if err != nil {
			s.Log.Error(err)
			s.HTTP.WriteError(w, api.ErrUnknown, err)
			return
		}
		_, err = tx.ExecContext(c.Context(), query, aa...)
		if err != nil {
			if err2 := tx.Rollback(); err2 != nil {
				s.Log.Error(err2)
				s.HTTP.WriteError(w, api.ErrUnknown, err2)
				return
			}
			s.Log.Error(err)
			s.HTTP.WriteError(w, api.ErrUnknown, err)
			return
		}
	}

	// Insert user token.
	_, err = tx.ExecContext(c.Context(),
		`insert into user_tokens (user_id, token, expired) values (?, ?, ?)`,
		curuser.ID, jwt, s.MySQL.Format(exp))

	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			s.Log.Error(err2)
			s.HTTP.WriteError(w, api.ErrUnknown, err2)
			return
		}
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	//  Get User .
	type bodyUser struct {
		ID             int64  `db:"id"              json:"id,omitempty"`
		Email          string `db:"email"           json:"email,omitempty"`
		Nickname       string `db:"nickname"        json:"nickname,omitempty"`
		Avatar         string `db:"avatar"          json:"avatar,omitempty"`
		AvatarURL      string `db:"avatar_url"      json:"avatar_url,omitempty"`
		Profile        string `db:"profile"         json:"profile"`
		NoteCount      int32  `db:"note_count"      json:"note_count"`
		FollowerCount  int32  `db:"follower_count"  json:"follower_count"`  // 关注我的
		FollowingCount int32  `db:"following_count" json:"following_count"` // 我关注的
	}

	user := &bodyUser{}
	err = tx.GetContext(c.Context(), user,
		`select id, email, nickname, avatar, profile from users where id = ?`, curuser.ID)

	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			s.Log.Error(err2)
			s.HTTP.WriteError(w, api.ErrUnknown, err2)
			return
		}
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	// Get Note count
	err = tx.QueryRow(`select count(*) from notes where id = ?`, curuser.ID).Scan(&user.NoteCount)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	user.NoteCount = 0

	// 我关注的用户（following）和 关注我的用户（followers）
	// 获取关于我的用户数量
	err = tx.QueryRow(`select count(*) from relationships where followed_id = ?`, curuser.ID).Scan(&user.FollowerCount)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	s.Log.Info(user.FollowerCount)

	// Get following user's count 或取我关于用户
	err = tx.QueryRow(`select count(*) from relationships where follower_id = ?`, curuser.ID).Scan(&user.FollowingCount)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	s.Log.Info(user.FollowingCount)

	user.AvatarURL, err = s.S3.Presign(user.Avatar)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	// Commit transaction.
	if err := tx.Commit(); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			s.Log.Error(err2)
			s.HTTP.WriteError(w, api.ErrUnknown, err2)
			return
		}
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	// Build response.
	body, err := json.Marshal(struct {
		Code  string    `json:"code,omitempty"`
		Token string    `json:"token,omitempty"`
		User  *bodyUser `json:"user,omitempty"`
	}{
		api.OK,
		jwt,
		user,
	})
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	// Write response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
