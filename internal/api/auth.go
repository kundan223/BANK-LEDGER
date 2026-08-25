package api

import (
	"encoding/json"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {

	var req loginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	user, err := s.userService.Login(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {

		writeError(
			w,
			http.StatusUnauthorized,
			"invalid email or password",
		)

		return
	}

	response := userResponse{
		ID:    user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Time.Format(
			"2006-01-02T15:04:05Z07:00",
		),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
