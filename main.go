package main

import (
	"fmt"
	"net/http"

	"groupietracker/handlers"
)

func main() {
	http.HandleFunc("/assets/", handlers.HandleForbiden)
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/{id}", handlers.HandleInfo)
	fmt.Println("Server On : http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
