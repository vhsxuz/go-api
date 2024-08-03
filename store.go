package main

import "database/sql"

type Store interface {
	// Users
	Createuser() error
}

type Storage struct {
	db *sql.DB
}

// Createuser implements Store.
func (storage *Storage) Createuser() error {
	panic("unimplemented")
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) CreateUser() error {
	return nil
}
