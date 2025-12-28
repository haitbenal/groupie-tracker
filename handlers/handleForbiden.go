package handlers

import (
	"net/http"
	"os"
	"strings"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	// make the access to files directory forbidden
	file, err := os.Stat(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		HandelError(w, "Page Not found", http.StatusNotFound)
		return
	}
	if file.IsDir() {
		HandelError(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
