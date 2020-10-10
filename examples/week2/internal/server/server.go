package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Storage interface {
	GetUser(id string) string
	SetUser(id, name string)
}

// Server is a main application server and hancler structure
type Server struct {
	r  *mux.Router
	st Storage
}

// New creates a new server
func New(st Storage) *Server {
	srv := &Server{
		st: st,
	}
	srv.setupRouter()
	return srv
}

func (s *Server) setupRouter() {
	s.r = mux.NewRouter()

	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")

	s.r.HandleFunc("/user/{id}", s.getUser).Methods("GET")
	s.r.HandleFunc("/user", s.setUser).Methods("POST")
	s.r.HandleFunc("/user/{id}", s.setUser).Methods("PUT")
}

// Run server
func (s *Server) Run(addres string) error {
	srv := &http.Server{
		Handler: s.r,
		Addr:    addres,
	}

	return srv.ListenAndServe()
}
