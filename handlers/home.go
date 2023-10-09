package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	Data := []Artist{}
	err := getElement(artistUrl, &Data)
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}

	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	err = ts.Execute(w, Data)
}
