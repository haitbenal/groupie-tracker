package handlers

import (
	"net/http"
)

type info struct {
	Code        int
	Description string
}

func HandlerErr(w http.ResponseWriter, description string, code int) {
	pageErr := info{Code: code, Description: description}
	if ErrParse != nil {
		w.WriteHeader(code)
		Temp.ExecuteTemplate(w, "errorePage.html", pageErr)
		return
	}
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "errorePage.html", pageErr)

}
