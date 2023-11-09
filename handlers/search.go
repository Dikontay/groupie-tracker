package handlers

import (
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var userInput string
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
	}
	if r.URL.Path != "/search" {
		errorHandler(w, http.StatusNotFound)
	}

	if err := r.ParseForm(); err != nil {
		errorHandler(w, http.StatusBadRequest)
	}
	userInput = r.FormValue("query")

	search(userInput)
}

func search(input string) {

	
}
