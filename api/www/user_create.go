package www

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yun313350095/Noonde/api"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

type userCreateForm struct {
	service *Service
	context api.Context

	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	AvatarData string `json:"avatar_data"`

	avatarBase64 string
	avatarExt    string
	avatar       string
	avatarSum    string
	avatarURL    string
}

func (f *userCreateForm) parseAndValidate() error {
	s := f.service
	c := f.context
	var err error

	// Parse request body.
	r := c.Request()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(f)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	// Convert.
	avatar := strings.Split(f.AvatarData, ",")
	if len(avatar) >= 2 {
		if strings.Contains(avatar[0], "data:image/png;") {
			f.avatarExt = "png"
		} else if strings.Contains(avatar[0], "data:image/jpeg;") {
			f.avatarExt = "jpg"
		}
		f.avatarBase64 = avatar[1]
	}
	f.avatarSum = fmt.Sprintf("%x", sha256.Sum256([]byte(f.AvatarData)))

	// Start validaon.
	perr := s.HTTP.NewParamError()
	tx := c.Tx()

	// Email is not empty?
	if f.Email == "" {
		perr.PushIfNotExists("email", api.ParamEmptyValue)
	}

	// Email is not registered?
	var i int64
	err = tx.GetContext(c.Context(), &i, `select id from users where email=?`, f.Email)
	if err != nil && err != sql.ErrNoRows {
		s.Log.Error(err)
		return err
	}
	if err == nil {
		perr.PushIfNotExists("email", api.ParamDuplicateUser)
	}

	//  Password is not empty ?
	if f.Password == "" {
		perr.PushIfNotExists("password", api.ParamEmptyValue)
	}

	// NickName is not empty ?
	if f.Nickname == "" {
		perr.PushIfNotExists("nickname", api.ParamEmptyValue)
	}

	//  AvatarData is not empty?
	if f.AvatarData == "" {
		perr.PushIfNotExists("avatar_data", api.ParamEmptyValue)
	}

	if perr.IsNotEmpty() {
		return perr
	}
	return nil
}

func (s *Service) userCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	// Setup context
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
	f := &userCreateForm{
		service: s,
		context: c,
	}
	perr := f.parseAndValidate()
	if perr != nil {
		s.Log.Error(perr)
		s.HTTP.WriteError(w, api.ErrInvalidRequest, perr)
		return
	}

	// Encrypt password.
	bytes, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}
	pass := string(bytes)

	s.Log.Info(pass)
	s.Log.Info(f.Email)
	s.Log.Info(f.Nickname)
	s.Log.Info(f.avatarSum)
	initUrlName := f.Email + "+++++noonde.com"

	// Insert user.
	res, err := tx.ExecContext(c.Context(), `
		insert into users (
		  email,
          password,
		  nickname,
	      avatar_sum,
		  max_tokens,
	      profile
		) values (
		   ?,?,?,?,5,''
		)`,
		f.Email,
		pass,
		f.avatarSum,
		initUrlName,
	)
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

	// Get user ID.
	userID, err := res.LastInsertId()
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

	// Update avatar path
	f.avatar = "user/" + strconv.FormatInt(userID, 10) + "." + f.avatarExt
	s.Log.Info(f.avatar)
	_, err = tx.ExecContext(c.Context(), `update users set avatar=? where id=?`, f.avatar, userID)
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
	f.avatarURL, err = s.S3.Presign(f.avatar)

	// Get User
	type bodyUser struct {
		ID             int64  `db:"id"              json:"id,omitempty"`
		Email          string `db:"email"           json:"email,omitempty"`
		NickName       string `db:"nickname"        json:"nickname,omitempty"`
		Avatar         string `db:"avatar"          json:"avatar,omitempty"`
		AvatarURL      string `db:"avatar_url"      json:"avatar_url,omitempty"`
		Profile        string `db:"profile"         json:"profile"`
		NoteCount      int32  `db:"note_count"      json:"note_count"`
		FollowerCount  int32  `db:"follower_count"  json:"follower_count"`
		FollowingCount int32  `db:"following_count" json:"following_count"`
	}
	user := &bodyUser{}
	err = tx.GetContext(c.Context(), user, `select id, email, nickname, avatar, profile from users where id = ?`, userID)
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
	user.AvatarURL, err = s.S3.Presign(user.Avatar)
	if err != nil {
		s.Log.Error(err)
		s.HTTP.WriteError(w, api.ErrUnknown, err)
		return
	}

	user.NoteCount = 0
	user.FollowerCount = 0
	user.FollowingCount = 0
	// Commit transaction.
	err = tx.Commit()
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

	// Build response.
	body, err := json.Marshal(struct {
		Code string    `json:"code,omitempty"`
		User *bodyUser `json:"user,omitempty"`
	}{
		api.OK,
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

	// Run job.
	s.Job.Add(f)
}

func (f *userCreateForm) Perform() {
	s := f.service

	// Write log.
	s.Log.Info("-> userCreateJob")

	// Update avatar image
	bb, err := base64.StdEncoding.DecodeString(f.avatarBase64)
	if err != nil {
		s.Log.Error(err)
		return
	}
	s.S3.Upload(
		f.avatar,
		bytes.NewReader(bb),
	)
}
