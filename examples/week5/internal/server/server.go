package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type API interface {
	GetTime() (time.Time, error)
}

type Server struct {
	api API
	srv http.Server
}

func New(addr string, api API) *Server {
	s := Server{
		api: api,
	}
	r := mux.NewRouter()
	r.HandleFunc("/", s.hello)
	s.srv = http.Server{
		Addr:    addr,
		Handler: r,
	}
	return &s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	t, err := s.api.GetTime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "current time is %s", t.Format(time.RFC822))
}
