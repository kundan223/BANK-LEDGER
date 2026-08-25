package service

import (
	"context"

	"bank-ledger/internal/db"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Login(
	ctx context.Context,
	email string,
	password string,
) (db.User, error) {

	user, err := s.store.GetUserByEmail(ctx, email)

	if err != nil {
		return db.User{}, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
