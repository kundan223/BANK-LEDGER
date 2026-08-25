package api

import (
	"net/http"

	"bank-ledger/internal/db"
	"bank-ledger/internal/service"
)

type Server struct {
	store       *db.Store
	userService *service.UserService
	router      *http.ServeMux
}

func NewServer(
	store *db.Store,
	userService *service.UserService,
) *Server {
	server := &Server{
		store:       store,
		userService: userService,
		router:      http.NewServeMux(),
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("GET /health", s.health)
	s.router.HandleFunc("POST /users", s.createUser)
	s.router.HandleFunc("POST /login", s.login)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) Handler() http.Handler {
	return s.router
}
