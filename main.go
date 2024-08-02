package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := newDatabaseStore()
	if err != nil {
		log.Fatal("[-] ERROR: ", err)
	}

	fmt.Printf("%+v\n", store)
	// server := newAPIServer(":8000")
	// server.run()
}
