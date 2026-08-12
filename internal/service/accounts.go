package service

import (
	"context" // what is this where are we importign it from 

	"bank-ledger/internal/db"
)

type AccountService struct { // one more service all though being used somewhere but for what // maybe its the connection to the store for this file idk 
	store *db.Store
}

func NewAccountService(store *db.Store) *AccountService {   // what does this ervice really do i can not see it being used in this file 
	return &AccountService{
		store: store,
	}
}

func (s *AccountService) CreateAccount(  // so this is the main service then what is this s ? it's name and why are we using a * with it 
	ctx context.Context, // is this context given when the service is called 
	userID int64,
) (db.Account, error) {  
	return s.store.CreateAccount(ctx, userID)
} 