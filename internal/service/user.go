package service

import (
	"context"
	"errors"
	"bank-ledger/internal/db"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"service/validation"
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

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return db.User{}, err
	}

	user, err :=  s.store.CreateUser(
		ctx,
		db.CreateUserParams{
			Email:        email,
			PasswordHash: string(passwordHash),

		},
	)

	if err != nil {
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
