package conf

import (
	"github.com/spf13/viper"
	"github.com/yun313350095/Noonde/api"
)

// Service ..
type Service struct {
	Viper *viper.Viper
	Log   api.LogService
}

// Bool ..
func (s *Service) Bool(key string) bool {
	return s.Viper.GetBool(key)
}

// Int ..
func (s *Service) Int(key string) int {
	return s.Viper.GetInt(key)
}

// Int64 ..
func (s *Service) Int64(key string) int64 {
	return s.Viper.GetInt64(key)
}

// String ..
func (s *Service) String(key string) string {
	return s.Viper.GetString(key)
}

// StringSlice ..
func (s *Service) StringSlice(key string) []string {
	return s.Viper.GetStringSlice(key)
}

//--------------------------------------------------------------------------------

// Load ..
func (s *Service) Load() error {
	s.Viper = viper.New()

	s.Log.Info(s.Viper.BindEnv("gopath"))
	s.Log.Info(s.Viper.BindEnv("goenv"))
	s.Log.Info(s.Viper.BindEnv("goconf"))
	//s.Log.Info(s.Viper.BindEnv("gogql"))

	s.Viper.SetDefault("gopath", "~/go")
	s.Viper.SetDefault("goenv", "local")

	root := s.Viper.GetString("gopath") + "/src/github.com/yun313350095/Noonde/api"
	s.Viper.Set("goroot", root)
	s.Viper.SetDefault("goconf", root+"/conf")

	s.Viper.SetConfigName(s.Viper.GetString("goenv"))
	s.Viper.AddConfigPath(s.Viper.GetString("goconf"))
	//s.Viper.SetDefault("gogql", root+"/scraper")

	return s.Viper.ReadInConfig()
}
