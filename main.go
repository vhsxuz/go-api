package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	config := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// fmt.Println(config)

	sqlStorage := MyNewSQLStorage(config)
	db, err := sqlStorage.Init()

	if err != nil {
		log.Fatal("[-] ", err)
	}

	store := NewStore(db)
	api := newAPIServer(":8000", store)
	api.Serve()
}
