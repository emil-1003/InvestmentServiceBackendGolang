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

	s.Path("/hello").Handler(handlers.Hello()).Methods("GET") // http://localhost:8585/api/v1/hello

	return &Server{name, version, s, port}, nil
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Port, s.Router)
}
