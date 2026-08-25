package service

import (
	"context"
	"errors"

	"bank-ledger/internal/db"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store *db.Store
}

func NewUserService(store *db.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	email string,
	password string,
) (db.User, error) {

	// Validate email
	if err := validateEmail(email); err != nil {
		return db.User{}, err
	}

	// Validate password
	if err := validatePassword(password); err != nil {
		return db.User{}, err
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return db.User{}, err
	}

	// Create user in database
	user, err := s.store.CreateUser(
		ctx,
		db.CreateUserParams{
			Email:        email,
			PasswordHash: string(passwordHash),
		},
	)

	if err != nil {

		// Check whether PostgreSQL returned a duplicate-key error
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return db.User{}, ErrEmailAlreadyExists
			}
		}

		return db.User{}, err
	}

	return user, nil
}
