package log

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yun313350095/Noonde/api"
	"runtime"
	"strings"
)

// Service ..
type Service struct {
	Conf api.ConfService
}

// Debug ..
func (s *Service) Debug(args ...interface{}) {
	logrus.WithField(s.caller()).Debug(args...)
}

// Info ..
func (s *Service) Info(args ...interface{}) {
	logrus.WithField(s.caller()).Info(args...)
}

// Warn ..
func (s *Service) Warn(args ...interface{}) {
	logrus.WithField(s.caller()).Warn(args...)
}

// Error ..
func (s *Service) Error(args ...interface{}) {
	logrus.WithField(s.caller()).Error(args...)
}

// Fatal ..
func (s *Service) Fatal(args ...interface{}) {
	logrus.WithField(s.caller()).Fatal(args...)
}

// Panic ..
func (s *Service) Panic(args ...interface{}) {
	logrus.WithField(s.caller()).Panic(args...)
}

// ErrorWithStacktrace ..
func (s *Service) ErrorWithStacktrace(info interface{}) {
	err, ok := info.(error)
	if !ok {
		err = fmt.Errorf("%v", info)
	}

	errs := err.Error() + "\n"
	for d := 0; ; d++ {
		pc, src, li, ok := runtime.Caller(d)
		if !ok {
			break
		}
		errs += fmt.Sprintf("  -> %02d: %s: %s (%d)\n", d, runtime.FuncForPC(pc).Name(), src, li)
	}
	s.Error(errors.New(errs))
}

// Load ..
func (s Service) Load() {
	logrus.SetLevel(logrus.InfoLevel)
	if s.Conf.String("goenv") != "prod" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

//--------------------------------------------------------------------------------

func (s *Service) caller() (string, int) {
	_, f, l, _ := runtime.Caller(2)
	f = strings.Replace(f, s.Conf.String("goroot")+"/", "", 1)
	return f, l
}
