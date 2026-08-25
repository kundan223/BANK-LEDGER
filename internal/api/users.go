package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"bank-ledger/internal/service"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	user, err := s.userService.CreateUser(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {

		if errors.Is(err, service.ErrEmailAlreadyExists) {
			writeError(
				w,
				http.StatusConflict,
				"email already exists",
			)
			return
		}

		writeError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := userResponse{
		ID:    user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Time.Format(
			"2006-01-02T15:04:05Z07:00",
		),
	}

	json.NewEncoder(w).Encode(response)
}
