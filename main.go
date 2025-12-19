package main

import (
	"groupie/handlers"
	"net/http"
)

func main() {
	
	
	http.HandleFunc("/", handlers.HandleHome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// handlers.HandlerErr("Status Internal Server Error", http.StatusInternalServerError)
	}
}
