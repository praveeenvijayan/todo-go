package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/praveeenvijayan/todo-go/pkg/router"
)

func main() {

	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Println("API is running!")
	http.ListenAndServe(":4000", r)
}
