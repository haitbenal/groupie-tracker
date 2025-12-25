package groupietracker

import (
	"groupietracker/handlers"
	"net/http"
)

func Run()  {
	http.HandleFunc("/", handlers.HandleHome)
	http.ListenAndServe(":8080", nil)
}