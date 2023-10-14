package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var (
	artistUrl    = "https://groupietrackers.herokuapp.com/api/artists"
	locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
	datesUrl     = "https://groupietrackers.herokuapp.com/api/dates"
	relationUrl  = "https://groupietrackers.herokuapp.com/api/relation"
)

type Artist struct {
	ID             int       `json:"id"`
	Image          string    `json:"image"`
	Name           string    `json:"name"`
	Members        []string  `json:"members"`
	CreationDate   int       `json:"creationDate"`
	FirstAlbumDate string    `json:"firstAlbum"`
	Locations      string    `json:"locations"`
	ConcertDates   string    `json:"concertDates"`
	Relations      Realtions `json:"omitempty"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Locations struct {
	Index []struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
		DatesUrl string   `json:"dates"`
	} `json:"index"`
}

type Realtions struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		//http.Error(w, "Not found", 404)
		errorHandler(w, http.StatusNotFound)
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	ts, err := template.ParseFiles("./templates/artists.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	id := r.URL.Query().Get("ID")
	if id == "" {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	artists := []Artist{}

	err = getElement(artistUrl, &artists)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}

	relations := Realtions{}
	err = getElement(relationUrl+"/"+strconv.Itoa(intId), &relations)
	if err != nil {
		fmt.Println(err)
		return
	}
	artists[intId-1].Relations = relations
	err = ts.Execute(w, artists[intId-1])
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

}
