package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTaskService(store Store) *TasksService {
	return &TasksService{
		store: store,
	}
}

func (service *TasksService) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/tasks", service.HandleGetAllTask).Methods("GET")
	router.HandleFunc("/tasks", service.HandleCreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", service.HandleGetTask).Methods("GET")
}

func (service *TasksService) HandleGetTask(write http.ResponseWriter, read *http.Request) {

}

func (service *TasksService) HandleCreateTask(write http.ResponseWriter, read *http.Request) {

}
