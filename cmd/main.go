package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveeenvijayan/todo-go/pkg/config"
	"github.com/praveeenvijayan/todo-go/pkg/handlers"
)

func main() {
	DB := config.Connect()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/todo", h.GetAllTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo/{id}", h.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo", h.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", h.UpdatedTodo).Methods(http.MethodPut)
	router.HandleFunc("/todo/{id}", h.DeleteTodo).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
