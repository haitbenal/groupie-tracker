package handlers

import (
	"net/http"
)

type Data struct {
	Code  int
	Error string
}

func HandelError(w http.ResponseWriter, des string, code int) {
	info := Data{Code: code, Error: des}
	if Err != nil {
		info = Data{Code: http.StatusInternalServerError, Error: "Internal Server Error"}
		w.WriteHeader(http.StatusInternalServerError)
		Temp.ExecuteTemplate(w, "errorPage.html", info)
		return
	}
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "errorPage.html", info)
}
