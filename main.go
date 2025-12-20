package main

import (
	"fmt"
	"groupie/handlers"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/statics/", handlers.HandleForbiden)
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path // مثال: /artist1
		idStr := strings.TrimPrefix(path, "/artist")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		// هنا تستدعي HandlerInfo وتمرير الرقم
		fmt.Println(id)
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		// handlers.HandlerErr("Status Internal Server Error", http.StatusInternalServerError)
	}
}
