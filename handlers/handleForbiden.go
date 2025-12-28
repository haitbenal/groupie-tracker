package handlers

import (
	"net/http"
	"os"
	"strings"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	// make the access to files directory forbidden
	file, err := os.Stat(strings.TrimPrefix(r.URL.Path, "/"))
	if file.IsDir() {
		HandelError(w, "Forbidden", http.StatusForbidden)
		return
	}
	if err != nil {
		HandelError(w, "Page Not found", http.StatusNotFound)
		return	
	}
	http.ServeFile(w, r, "./assets/"+r.URL.Path[len("/assets/"):])
}