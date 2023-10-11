package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorData struct {
	ErrorCode int
	ErrorDesc string
}

func errorHandler(w http.ResponseWriter, status int) {
	errm := ErrorData{ErrorCode: status, ErrorDesc: http.StatusText(status)}
	fmt.Println(errm)
	temp, err := template.ParseFiles("./templates/errors.html")
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