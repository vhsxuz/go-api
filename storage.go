package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(string) error
	UpdateAccount(*Account) error
	GetAccountById(string) (*Account, error)
}

type DatabaseStore struct {
	db *sql.DB
}

func newDatabaseStore() (*DatabaseStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseStore{
		db: db,
	}, nil

}
