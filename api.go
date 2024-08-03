package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	store   Store
}

func newAPIServer(address string, store Store) *APIServer {
	return &APIServer{
		address: address,
		store:   store,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("api/v1/").Subrouter()

	log.Println("[+] Listening on PORT")
	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
