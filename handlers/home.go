package handlers

import (
	"net/http"
	"text/template"

	
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", 404)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}

	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	err = ts.Execute(w, nil)
}