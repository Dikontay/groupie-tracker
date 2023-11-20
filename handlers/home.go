package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "page not found", http.StatusNotFound)

		errorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		// http.Error(w, "method not allowed", 405)
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	artistsCh := make(chan []Artist, 1)

	go func() {
		artists := []Artist{}
		err := getElement(artistUrl, &artists)
		if err != nil {
			fmt.Println(err)
			artistsCh <- nil
			return
		} else {
			artistsCh <- artists
		}
	}()

	ts, err := template.ParseFiles("./static/templates/index.html")
	if err != nil {
		// http.Error(w, "internal server error", 500)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	defer close(artistsCh)
	select {
	case artists := <-artistsCh:
		if artists == nil {
			// http.Error(w, "Not found", http.StatusNotFound)
			errorHandler(w, http.StatusNotFound)
		} else {
			err = ts.Execute(w, artists)
			if err != nil {
				errorHandler(w, http.StatusInternalServerError)
				return
			}
		}
	case <-time.After(3 * time.Second):
		errorHandler(w, http.StatusRequestTimeout)
		return

	}
	err = json.NewEncoder(w).Encode(getSuggestionsForSearch())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getSuggestionsForSearch() []string {
	suggestions := []string{}
	artists := []Artist{}
	err := getElement(artistUrl, &artists)
	if err != nil {
		fmt.Println(err)
	}

	relations := []Realtions{}
	for i := range artists {
		relation := Realtions{}
		err = getElement(relationUrl+"/"+strconv.Itoa(i+1), &relation)
		if err != nil {
			fmt.Println(err)
		}
		for index := range relations {
			for location := range relations[index].DatesLocations {
				suggestions = append(suggestions, location)
			}
		}
		relations = append(relations, relation)
		suggestions = append([]string{artists[i].Name + "artist/band", artists[i].FirstAlbumDate, strconv.Itoa(artists[i].CreationDate)})
		for j := range artists[i].Members {
			suggestions = append(suggestions, artists[i].Members[j]+" - member")
		}

	}
	return suggestions
}
