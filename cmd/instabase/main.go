package main

import (
	"context"
	"github.com/yun313350095/Noonde/api"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// Version 0.0.1

func main() {
	var err error

	s := newService()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)

	stop := make(chan bool, 1)
	done := make(chan bool, 1)

	//Instabase search ---------------------------------------
	//go func(stop <-chan bool, done chan<- bool) {
	//loop:
	//	for {
	//		err = s.doSearchSpaceList()
	//		if err != nil {
	//			s.Log.Error(err)
	//		}
	//		select {
	//		case <-stop:
	//			break loop
	//		default:
	//
	//		}
	//
	//		// TODO: debug
	//		// break loop
	//
	//		time.Sleep(time.Minute * 1)
	//	}
	//
	//	s.Log.Info("Search room stop.")
	//	done <- true
	//}(stop, done)

	// Instabase get space details --------------------------
	go func(stop <-chan bool, done chan<- bool) {
	loop:
		for {
			err = s.doGetSpaceDetails()
			if err != nil {
				s.Log.Error(err)
			}
			select {
			case <-stop:
				break loop
			default:

			}

			// TODO: debug
			// break loop
			time.Sleep(time.Minute * 1)
		}
	}(stop, done)

	<-sig
	s.Log.Info("Stopping..")

	stop <- true
	//stop <- true

	<-done
	//<-done
}

func (s *service) doSearchSpaceList() error {
	var err error

	defer func() {
		if info := recover(); info != nil {
			s.Log.ErrorWithStacktrace(info)
		}
	}()

	// Get instabase event_type and start page
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	var eventTypes []*api.InstabaseEventType
	err = s.MySQL.Reader().SelectContext(ctx, &eventTypes,
		`select id, event_type, start_page, hourly_at, event_type_text, event_type_en, created, updated from instabase_event_types where state = 1 order by hourly_at is null desc, hourly_at asc limit 1`)

	if err != nil {
		s.Log.Error(err)
		return err
	}

	done := make(chan bool, len(eventTypes))
	var wg sync.WaitGroup

	for _, oneType := range eventTypes {
		wg.Add(1)
		go func(oneType api.InstabaseEventType) {
			defer func() {
				wg.Done()
				done <- true
			}()

			// Begin transaction.
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel()

			tx, err := s.MySQL.Begin(ctx)
			if err != nil {
				s.Log.Error(err)
				return
			}
			client, err := s.Scraper.NewInstabaseClient()
			if err != nil {
				s.Log.Error(err)
				return
			}

			httpRetried := 0
		again1:
			briefSpace, err := s.searchSpaceList(client, oneType.StartPage, oneType.EventType)
			s.Log.Info("One Type:  --->" + oneType.EventTypeText)
			if err != nil {
				if strings.Contains(err.Error(), "Proxy Error") && httpRetried <= 10 {
					s.Log.Info("Retrying HTTP.  " + strconv.Itoa(httpRetried))
					httpRetried++
					time.Sleep(time.Second)
					goto again1
				}
				s.Log.Error(err)
				return
			}

		again2:
			err = s.saveSpaces(ctx, tx, briefSpace, oneType)
			if err != nil {
				if s.MySQL.Deadlocked(err) {
					s.Log.Info("=== Trying again ===")
					if err2 := tx.Rollback(); err2 != nil {
						s.Log.Error(err2)
						return
					}
					goto again2
				}
				if err2 := tx.Rollback(); err2 != nil {
					s.Log.Error(err2)
					return
				}
			}

			// Commit transaction .
			err = tx.Commit()
			if err != nil {
				if err2 := tx.Rollback(); err2 != nil {
					s.Log.Error(err2)
					return
				}
				s.Log.Error(err)
				return
			}

		}(*oneType)

		wg.Wait()
	}
	return nil
}

//  doGetSpaceDetails
func (s *service) doGetSpaceDetails() error {
	var err error

	defer func() {
		if info := recover(); info != nil {
			s.Log.ErrorWithStacktrace(info)
		}
	}()

	// Get space details
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var spaces []*api.Space

	err = s.MySQL.Reader().SelectContext(ctx, &spaces,
		`select * from spaces where platform_id = 2 order by hash_at is null desc, hash_at asc limit 1`)

	if err != nil {
		s.Log.Error(err)
		return err
	}

	done := make(chan bool, len(spaces))
	var wg sync.WaitGroup

	for _, space := range spaces {
		wg.Add(1)

		go func(space api.Space) {
			defer func() {
				wg.Done()
				done <- true
			}()

			// Begin transaction.
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel()
			tx, err := s.MySQL.Begin(ctx)
			if err != nil {
				s.Log.Error(err)
				return
			}

			client, err := s.Scraper.NewInstabaseClient()
			if err != nil {
				s.Log.Error(err)
				return
			}

			httpRetried := 0
		again1:
			roomDetails, err := s.getSpaceDetails(client, space.Iop)
			if err != nil {
				if strings.Contains(err.Error(), "Proxy Error") && httpRetried <= 10 {
					s.Log.Info("Retrying HTTP. " + strconv.Itoa(httpRetried))
					httpRetried++
					time.Sleep(3 * time.Second)
					goto again1
				}
				s.Log.Error(err)
				return
			}

		again2:
			err = s.saveSpaceDetails(ctx, tx, roomDetails)
			if err != nil {
				if s.MySQL.Deadlocked(err) {
					s.Log.Info("=== Trying again ===")
					if err2 := tx.Rollback(); err2 != nil {
						s.Log.Error(err2)
						return
					}
					goto again2
				}
				if err2 := tx.Rollback(); err2 != nil {
					s.Log.Error(err)
					return
				}
			}

			// Commit transaction.
			err = tx.Commit()
			if err != nil {
				if err2 := tx.Rollback(); err2 != nil {
					s.Log.Error(err2)
					return
				}
				s.Log.Error(err)
				return
			}
		}(*space)
	}

	return nil
}
