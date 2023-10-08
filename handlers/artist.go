package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	artistUrl = "https://groupietrackers.herokuapp.com/api/artists"
	locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
	datesUrl = "https://groupietrackers.herokuapp.com/api/dates"
	relationUrl = "https://groupietrackers.herokuapp.com/api/relation"

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

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}



type Locations struct {
	Index []struct {
		ID    int      `json:"id"`
		Location []string `json:"locations"`
		DatesUrl string `json:"dates"`
	} `json:"index"`
}

type Realtions struct {
	Index[]struct {
		ID int `json:"id"`
		DatesLocations map[string][]string `json"datesLocations"`
	}`json"index"`
}



func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		http.Error(w, "Not found", 404)
	}
	if r.Method != http.MethodGet {
		fmt.Println("not right method")
		return
	}

	ts,err := template.ParseFiles("./templates/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}	

	// data := []Artist{}
	//data := Dates{}
//	data:= Locations{}
	data:= Realtions{}

	err = getElement(relationUrl, &data)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	ts.Execute(w, data)

}
