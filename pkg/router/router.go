package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveeenvijayan/todo-go/pkg/config"
	"github.com/praveeenvijayan/todo-go/pkg/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	DB := config.Connect()
	h := handlers.New(DB)
	router.HandleFunc("/todo", h.GetAllTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo/{id}", h.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo", h.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", h.UpdatedTodo).Methods(http.MethodPut)
	router.HandleFunc("/todo/{id}", h.DeleteTodo).Methods(http.MethodDelete)

	return router
}
