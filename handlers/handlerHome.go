package handlers

import (
	"net/http"
	"html/template"
)

var Temp, Err = template.ParseGlob("./templates/*.html")

func HandleHome(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		HandelError(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		HandelError(w, "Page Not Found", http.StatusNotFound)
		return
	}
	data := FetchData()
	err := 	Temp.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		HandelError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}