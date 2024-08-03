package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func MyNewSQLStorage(config mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal("[-] ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("[-] ", err)
	}

	log.Println("[+] MySQL Connection Successful")
	return &MySQLStorage{
		db: db,
	}
}

func (storage *MySQLStorage) Init() (*sql.DB, error) {
	// initialize the tables

	return storage.db, nil
}
