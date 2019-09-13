package www

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yun313350095/Noonde/api"
	"strings"
	"time"
)

const (
	jwtIssuer  = "api"
	jwtSubject = "token"
)

type claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// Service ..
type Service struct {
	HTTP    api.HTTPService
	Conf    api.ConfService
	Log     api.LogService
	Mail    api.MailService
	S3      api.S3Service
	MySQL   api.MySQLService
	Elastic api.ElasticService
	Job     api.JobService
}

func (s *Service) checkAuth(c api.Context) error {
	var err error

	r := c.Request()

	var tk string
	tt, ok := r.URL.Query()["token"]
	if ok && len(tt) > 0 {
		tk = tt[0]
	} else {
		aa, ok := r.Header["Authorization"]
		if !ok || len(aa) < 1 {
			err = errors.New("No authorization headers")
			s.Log.Error(err)
			return err
		}

		a := aa[0]
		tk = strings.TrimPrefix(a, "Bearer ")
	}

	token, err := jwt.ParseWithClaims(tk, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Conf.String("auth.key")), nil
	})
	if err != nil {
		s.Log.Error(err)
		return err
	}
	if !token.Valid {
		err = errors.New("Token is not valid")
		s.Log.Error(err)
		return err
	}
	if token.Claims.(*claims).StandardClaims.ExpiresAt < time.Now().Unix() {
		err = errors.New("Token expired")
		s.Log.Error(err)
		return err
	}

	user := &api.User{}
	err = s.MySQL.Reader().Get(user, `select * from users where id = ?`, token.Claims.(*claims).UserID)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	c.SetCurUser(user)

	atoken := api.UserToken{}
	err = s.MySQL.Reader().Get(&atoken, `select * from user_tokens where user_id =? and token=?`,
		user.ID,
		tk)
	if err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}
