package server

import (
	"fmt"
	"net/http"

	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Name    string
	Version string
	Router  *mux.Router
	Port    string
}

func New(name, version, port, path string) (*Server, error) {
	r := mux.NewRouter()

	s := r.PathPrefix(fmt.Sprintf("/%s/%s", path, version)).Subrouter()

	s.Path("/signup").Handler(handlers.Signup()).Methods("POST")

	return &Server{name, version, s, port}, nil
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Port, s.Router)
}
