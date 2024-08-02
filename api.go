package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

type APIError struct {
	Error string
}

func newAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) run() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/accounts", makeHTTPHandler(s.handleAccount))
	router.HandleFunc("/api/v1/accounts/{id}", makeHTTPHandler(s.handleGetAccount))

	log.Println("[+] Server running on port", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("[-] Method Not Allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := "4d63095c-f51c-4b76-a2fe-e17da67ba436"
	fmt.Println(id)
	return writeJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	account := newAccount("Andreas", "Alexander")
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("[-] JSON Encoding Error: ", err)
		return err
	}
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
