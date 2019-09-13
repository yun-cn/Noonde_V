package mysql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yun313350095/Noonde/api"
	"strings"
	"time"
)

// Service ..
type Service struct {
	Conf api.ConfService
	Log  api.LogService

	writer   *sqlx.DB
	reader   *sqlx.DB
	minTime  *time.Time
	maxTime  *time.Time
	location *time.Location
}

// Begin ..
func (s *Service) Begin(ctx context.Context) (*sqlx.Tx, error) {
	return s.writer.BeginTxx(ctx, nil)
}

// Deadlocked ..
func (s *Service) Deadlocked(err error) bool {
	if strings.Contains(err.Error(), "Lock wait timeout exceeded") {
		return true
	}

	if strings.Contains(err.Error(), "Deadlock found") {
		return true
	}

	return false
}

// Format ..
func (s *Service) Format(t time.Time) string {
	return t.In(s.location).Format("2006-01-02T15:04:05")
}

// MaxTime ..
func (s *Service) MaxTime() *time.Time {
	return s.maxTime
}

// MinTime ..
func (s *Service) MinTime() *time.Time {
	return s.minTime
}

// Reader ..
func (s *Service) Reader() *sqlx.DB {
	return s.reader
}

// SafeTime ..
func (s *Service) SafeTime(t time.Time) time.Time {
	if t == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		return *s.minTime
	}

	return t
}

// SumUp ..
func (s *Service) SumUp(ctx context.Context, tx *sqlx.Tx, ofType string, forType string, forIDs []int64) error {
	return nil
}

// Writer ..
func (s *Service) Writer() *sqlx.DB {
	return s.writer
}

//--------------------------------------------------------------------------------

// Load ..
func (s *Service) Load() error {
	minTime := time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC)
	maxTime := time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC)
	s.minTime = &minTime
	s.maxTime = &maxTime

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		s.Log.Error(err)
		return err
	}
	s.location = location

	writer, err := sqlx.Connect("mysql", s.Conf.String("mysql.writer"))
	if err != nil {
		s.Log.Error(err)
		return err
	}
	writer.SetMaxOpenConns(s.Conf.Int("mysql.max_open_conns"))
	s.writer = writer

	reader, err := sqlx.Connect("mysql", s.Conf.String("mysql.reader"))
	if err != nil {
		s.Log.Error(err)
		return err
	}
	reader.SetMaxOpenConns(s.Conf.Int("mysql.max_open_conns"))
	s.reader = reader
	return nil
}
