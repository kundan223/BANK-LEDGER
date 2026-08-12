package api

import (
	"encoding/json"
	"net/http"

	"bank-ledger/internal/db"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := s.store.CreateUser(
		r.Context(),
		db.CreateUserParams{
			Email:        req.Email,
			PasswordHash: req.Password,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}