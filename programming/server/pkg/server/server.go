package server

import (
	"myrest/pkg/handler"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	handler *handler.Handler
	router  *http.ServeMux
}

func InitAPIServer(h *handler.Handler, rt *http.ServeMux) *APIServer {
	return &APIServer{
		handler: h,
		router:  rt,
	}
}

func (s *APIServer) StartServer(port string) error {
	configureRouter := s.configureRouter()
	return http.ListenAndServe(":"+port, configureRouter)
}

func (s *APIServer) configureRouter() http.Handler {
	s.router.HandleFunc("/events_for_day", s.handler.GetEventsForDay)
	s.router.HandleFunc("/events_for_week", s.handler.GetEventsForWeek)
	s.router.HandleFunc("/events_for_month", s.handler.GetEventsForMonth)
	s.router.HandleFunc("/create_event", s.handler.CreateEvent)
	s.router.HandleFunc("/update_event", s.handler.UpdateEvent)
	s.router.HandleFunc("/delete_event", s.handler.DeleteEvent)
	return s.loggingMiddleware(s.router)
}

func (s *APIServer) loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"url":         r.RequestURI,
			"method":      r.Method,
		}).Info("started")

		rw := &responseWriter{w, http.StatusOK}
		start := time.Now()

		h.ServeHTTP(rw, r)

		logger := logrus.WithFields(logrus.Fields{
			"code":      rw.code,
			"text code": http.StatusText(rw.code),
			"duration":  time.Now().Sub(start),
		})

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}

		logger.Log(level, "return response")
	})
}
