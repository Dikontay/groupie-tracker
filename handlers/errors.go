package handlers

import (
	"net/http"
	"text/template"
)

type ErrorData struct {
	ErrorCode int
	ErrorDesc string
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errm := ErrorData{ErrorCode: status, ErrorDesc: http.StatusText(status)}
	
	temp, err := template.ParseFiles("./ui/errors.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, errm)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
