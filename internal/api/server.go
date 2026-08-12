package api

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	server := &Server{
		router: http.NewServeMux(),
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("GET /health", s.health)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) Handler() http.Handler {
	return s.router
}