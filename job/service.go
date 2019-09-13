package job

import (
	"github.com/yun313350095/Noonde/api"
	"sync/atomic"
)

// Service ..
type Service struct {
	Channel chan api.Job
	Log     api.LogService
	Conf    api.ConfService

	inShutdown int32
	count      int32
}

// Add ..
func (s *Service) Add(job api.Job) {
	if atomic.LoadInt32(&s.inShutdown) != 0 {
		job.Perform()
		return
	}

	go func() {
		s.Channel <- job
	}()
}

// Count ..
func (s *Service) Count() int32 {
	return s.count
}

// SetInShutdown ..
func (s *Service) SetInShutdown() {
	atomic.StoreInt32(&s.inShutdown, 1)
}

// StartWorkers ..
func (s *Service) StartWorkers() {
	for i := 1; i <= s.Conf.Int("job.worker_num"); i++ {
		go func() {
			defer func() {
				if info := recover(); info != nil {
					atomic.AddInt32(&s.count, -1)
					s.Log.ErrorWithStacktrace(info)
				}
			}()

			s.Log.Info("-> StartWorkers")
			for job := range s.Channel {
				s.Log.Info("-> job.Perform")
				atomic.AddInt32(&s.count, 1)
				job.Perform()
				atomic.AddInt32(&s.count, -1)
			}
		}()
	}
}

//--------------------------------------------------------------------------------

// Load ..
func (s *Service) Load() {
	s.Channel = make(chan api.Job, s.Conf.Int("job.queue_size"))
}
