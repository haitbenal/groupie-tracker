package handlers

import (
	"net/http"
	"text/template"
)

var Temp, ErrParse = template.ParseGlob("./templates/*.html")

func HandleHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		if r.Method != http.MethodGet {
			HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		HandleArtist(w, r)
	} else {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}

}
