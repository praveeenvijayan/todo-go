package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/praveeenvijayan/todo-go/pkg/models"
)

func (h handler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	var todo []models.ToDo

	if result := h.DB.Find(&todo); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}
