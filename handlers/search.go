package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
 // Check if the method is GET
 if r.Method != http.MethodGet {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}

// Get the query parameter
query := r.URL.Query().Get("query")
if query == "" {
	http.Error(w, "Query not provided", http.StatusBadRequest)
	return
}
artists := []Artist{}
err := getElement(artistUrl,&artists)
if err != nil {
	errorHandler(w, http.StatusInternalServerError)
}
// Perform the search and get suggestions
artist := search(query, artists)

// Convert the suggestions to JSON and send them back
jsonResponse, err := json.Marshal(artist)
if err != nil {
	http.Error(w, "Failed to marshal suggestions", http.StatusInternalServerError)
	return
}

w.Header().Set("Content-Type", "application/json")
w.Write(jsonResponse)
}

func search(input string, artisits[]Artist) Artist {
	for _, artist := range artisits {
		if input == artist.Name{
			return artist
		} else {
			for _, member := range artist.Members{
				if input==member{
					return artist
				}
			}
		}
	} 
	return Artist{}
}
