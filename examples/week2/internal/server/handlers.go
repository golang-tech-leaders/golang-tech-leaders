package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "parameter user is mandatory", http.StatusBadRequest)
		return
	}

	userName := s.st.GetUser(id)
	if userName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "user id %q name %q", id, userName)
}

func (s *Server) setUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "parameter user is mandatory: "+err.Error(), http.StatusBadRequest)
		return
	}

	s.st.SetUser(req.User, req.Name)
	fmt.Fprintf(w, "user id %q name %q", req.User, req.Name)
}

type UpdateUserRequest struct {
	User string `json:"user"`
	Name string `json:"name"`
}
