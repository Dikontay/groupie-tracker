package handlers

import (
	"fmt"
	"groupie-trakcer/singleton"
	"log"
	"net/http"
	"html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./static/templates/index.html")
	if err != nil {

		fmt.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	Data, err := singleton.GetAllData()
	if err != nil {
		log.Printf("Getting all data error %e", err)
		errorHandler(w, http.StatusInternalServerError)
	}

	err = ts.Execute(w, Data.AllArtists)
	if err != nil {
		fmt.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
}
