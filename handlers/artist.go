package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	artistUrl = "https://groupietrackers.herokuapp.com/api/artists"
	locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"


)

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbumDate string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	Relations      string   `json:"relations"`
}

func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
	}
	if r.Method != http.MethodPost {
		return
	}

	ts,err := template.ParseFiles("./templates/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}	

	data := []Artist{}

	err = getElement(artistUrl, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	ts.Execute(w, data)

}
