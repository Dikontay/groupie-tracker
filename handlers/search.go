package handlers

import (
	"encoding/json"
	"groupie-trakcer/singleton"
	"net/http"
	"strings"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	// Check if the method is GET

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// the options method is used as a preflight rewuest in cors. it is sent by the browther automatically before the actual request (like get and post) in
		// certain situations particularly when the request is more complex
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query not provided", http.StatusBadRequest)
		return
	}

	data, err := singleton.GetAllData()

	var filteredData []string

	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}

	for _, artist := range data.AllArtists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filteredData = append(filteredData, artist.Name+"- artist/band")
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				filteredData = append(filteredData, member+" - member")
			}
		}

		if strings.Contains(strings.ToLower(artist.ConcertDates), strings.ToLower(query)) {
			filteredData = append(filteredData, artist.ConcertDates)
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbumDate), strings.ToLower(query)) {
			filteredData = append(filteredData, artist.FirstAlbumDate)
		}
	}

	for i := range data.AllLocations {
		for j := range data.AllLocations[i].Index {
			for _, location := range data.AllLocations[i].Index[j].Location {
				if strings.Contains(strings.ToLower(location), strings.ToLower(query)) {
					filteredData = append(filteredData, location)
				}
			}
		}
	}

	// Convert the suggestions to JSON and send them back
	jsonResponse, err := json.Marshal(filteredData)
	if err != nil {
		http.Error(w, "Failed to marshal suggestions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(jsonResponse)
}
