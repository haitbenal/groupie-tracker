package main

import (
	"net/http"

	"groupie/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/statics/", handlers.HandleForbiden)
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/artist/{id}", handlers.HandleInfo)

	http.ListenAndServe(":8080", mux)
}
