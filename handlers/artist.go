package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"groupie-trakcer/models"
	"groupie-trakcer/singleton"
)

type DataToSend struct {
	Artist   models.Artist
	Relation models.Relations
}

func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		// http.Error(w, "Not found", 404)

		errorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	ts, err := template.ParseFiles("./static/templates/artists.html")
	if err != nil {
		// fmt.Println(err)
		log.Printf("Can not parse index.html %e", err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("ID")
	if id == "" {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("Can not get element %e", err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	if id[0] == '0' {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if intId < 1 || intId > 52 {
		errorHandler(w, http.StatusBadRequest)
		// fmt.Println("ERROR")
		return
	}
	if err != nil {
		log.Println(err)
		errorHandler(w, http.StatusBadRequest)
		return
	}

	data, err := singleton.GetAllData()
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	dataToSend := &DataToSend{data.AllArtists[intId-1], data.AllRelations[intId-1]}
	// dataToSend := struct{Artist models.Artist, Relation models.Relations}{}
	// relations := models.Relations{}
	// err = getElement(relationUrl+"/"+strconv.Itoa(intId), &relations)
	if err != nil {
		log.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	// artists[intId-1].Relations = relations
	err = ts.Execute(w, dataToSend)
	if err != nil {
		log.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
}
