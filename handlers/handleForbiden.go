package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil || info.IsDir() {
		HandlerErr(w, "Forbidden", http.StatusForbidden)
		return
	}
	
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}
