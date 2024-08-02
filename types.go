package main

import (
	"github.com/google/uuid"
)

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BankID    string `json:"bank_id"`
	Balance   int64  `json:"balance"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		BankID:    uuid.NewString(),
		Balance:   0,
	}
}
