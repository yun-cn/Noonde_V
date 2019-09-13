package http

import (
	cont "context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/yun313350095/Noonde/api"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

// Service ..
type Service struct {
	Status map[string]int
	Conf   api.ConfService
	Log    api.LogService
	//WS      api.WSService
	Elastic api.ElasticService
	Job     api.JobService
	WWW     api.WWWService

	//upgrader   *websocket.Upgrader
	router     *httprouter.Router
	negroni    *negroni.Negroni
	server     *http.Server
	inShutdown int32
}

// NewContext .
func (s *Service) NewContext(ctx cont.Context) api.Context {
	return &context{context: ctx}
}

// NewParamError ..
func (s *Service) NewParamError() api.ParamError {
	return &paramError{}
}

// Router ..
func (s *Service) Router() *httprouter.Router {
	return s.router
}

// WriteError ..
func (s *Service) WriteError(w http.ResponseWriter, code string, err error) {
	w.Header().Set("Content-Type", "application/json")

	var body []byte

	switch err2 := err.(type) {
	case api.ParamError:
		var err3 error
		body, err3 = json.Marshal(struct {
			Code    string         `json:"code"`
			Message string         `json:"message"`
			Detail  api.ParamError `json:"detail"`
		}{
			code,
			message[code],
			err2,
		})
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}

	default:
		var err3 error
		body, err3 = json.Marshal(struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Detail  string `json:"detail"`
		}{
			code,
			message[code],
			err.Error(),
		})
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(status[code])
	w.Write(body)
}

//--------------------------------------------------------------------------------

// GracefulShutdown ..
func (s *Service) GracefulShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	s.Job.SetInShutdown()
	atomic.StoreInt32(&s.inShutdown, 1)
	s.Elastic.StopClient()

	timeout := 10 * time.Second
	ctx, cancel := cont.WithTimeout(cont.Background(), timeout)
	defer cancel()

	s.Log.Info("Gracefully shutting down..")
	if err := s.server.Shutdown(ctx); err != nil {
		s.Log.Fatal(err)
		return
	}

	i := 0
	for {
		if s.Job.Count() == 0 || i >= 30 {
			break
		}
		time.Sleep(time.Second)
		i++
	}

	s.Log.Info("Server stopped.")
}

// ListenAndServe ..
func (s *Service) ListenAndServe() {
	go func() {
		s.server = &http.Server{Addr: ":" + s.Conf.String("http.port"), Handler: s.negroni}

		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			s.Log.Fatal(err)
		}
	}()
}

// Load ..
func (s *Service) Load() {
	s.router = httprouter.New()
	s.negroni = negroni.New()
}

// SetMiddlewares ..
func (s *Service) SetMiddlewares() {
	s.negroni.Use(s.loggerHandler())
	s.negroni.Use(s.corsHandler())
	s.router.PanicHandler = s.panicHandler()
}

// SetRoutes ..
func (s *Service) SetRoutes() {
	s.WWW.SetRoutes()
	s.router.ServeFiles("/static/*filepath", http.Dir(s.Conf.String("goroot")+"/static"))
	s.negroni.UseHandler(s.router)
}

//--------------------------------------------------------------------------------
func (s *Service) corsHandler() negroni.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins: s.Conf.StringSlice("cors.origins"),
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// Debug:          true,
	})
	return cors
}

func (s *Service) loggerHandler() negroni.Handler {
	fmt := "{{.StartTime}} - {{.Status}} - {{.Duration}} - {{.Method}} - {{.Hostname}} - {{.Path}} - {{.Request.RemoteAddr}} - {{.Request.UserAgent}}"
	log := negroni.NewLogger()
	log.SetFormat(fmt)
	return log
}

func (s *Service) panicHandler() func(w http.ResponseWriter, r *http.Request, info interface{}) {
	return func(w http.ResponseWriter, r *http.Request, info interface{}) {
		s.Log.ErrorWithStacktrace(info)

		err, ok := info.(error)
		if !ok {
			err = fmt.Errorf("%v", info)
		}
		s.WriteError(w, api.ErrUnknown, err)
	}
}
