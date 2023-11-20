package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	// Check if the method is GET

	if r.Method == http.MethodOptions{
		w.Header().Set("Access-Contro;-Allow-Origin", "*")
		//the options method is used as a preflight rewuest in cors. it is sent by the browther automatically before the actual request (like get and post) in
		// certain situations particularly when the request is more complex
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return 
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the query parameter
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query not provided", http.StatusBadRequest)
		return
	}
	artists := []Artist{}
	err := getElement(artistUrl, &artists)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)

	}
	// Perform the search and get suggestions
	suggestions := getSuggestions(query)
	json.NewEncoder(w).Encode(suggestions)

	// // Convert the suggestions to JSON and send them back
	// jsonResponse, err := json.Marshal(artist)
	// if err != nil {
	// 	http.Error(w, "Failed to marshal suggestions", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonResponse)
}

func search(input string, artisits []Artist) Artist {
	for _, artist := range artisits {
		if input == artist.Name {
			return artist
		} else {
			for _, member := range artist.Members {
				if input == member {
					return artist
				}
			}
		}
	}
	return Artist{}
}

func getSuggestions(query string)[]string{
   artists := []Artist{}
   err := getElement(artistUrl, &artists)
   if err != nil {
	fmt.Println(err)
	return nil
   }
   relations := []Realtions{}
   for i := range artists {
	relation := Realtions{}
	err = getElement(relationUrl+"/"+strconv.Itoa(i+1), &relation)
	if err != nil {
		fmt.Println(err)
		return nil
	   }
	relations=append(relations, relation)
   }
   
 
   matches := []string{}
   query = strings.ToLower(query)
   for i := range artists {
		if strings.Contains(strings.ToLower(query), strings.ToLower(artists[i].Name)){
			matches=append(matches, artists[i].Name)
		} else if strings.Contains(strings.ToLower(query), strconv.Itoa(artists[i].CreationDate)){
			matches = append(matches, strconv.Itoa(artists[i].CreationDate))
		} else if strings.Contains(strings.ToLower(query), artists[i].FirstAlbumDate){
			matches = append(matches, artists[i].FirstAlbumDate)
		}

		for _, member := range artists[i].Members {
			if strings.Contains(query, strings.ToLower(member)) {
				matches = append(matches, member)
			}
		}
		
   }

   for i := range relations {
	for j := range relations[i].DatesLocations{
		if strings.Contains(query, strings.ToLower(j)) {
			matches=append(matches, j)
		}
	}
   }

   return matches
}
