package api

import (
	"net/http"

	"bank-ledger/internal/db"
)

type Server struct {
	store  *db.Store
	router *http.ServeMux
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: http.NewServeMux(),
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("GET /health", s.health)
	s.router.HandleFunc("POST /users", s.createUser)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) Handler() http.Handler {
	return s.router
}