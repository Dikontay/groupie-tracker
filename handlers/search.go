package handlers

import (
	"fmt"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var userInput string
	switch r.Method {
	case http.MethodGet:
		userInput = r.FormValue("query")
		fmt.Println(userInput)

	case http.MethodPost:
		_, err := w.Write([]byte(userInput))
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
		}
		fmt.Println("we should post something")
	default:
		errorHandler(w, http.StatusMethodNotAllowed)
	}

}

// func contains(dataset Artist, substr string) []Artist {
// 	opt := make(map[string]string)
// 	for _, val := range dataset {
// 		opt[val.Name] = " -> name"

// 		for _, mems := range val.Members {
// 			opt[mems] = " -> member"
// 		}
// 		opt[strconv.Itoa(val.CreationDate)] = " -> creation date"
// 		opt[val.FirstAlbum] = " -> first album"

// 		for key, element := range val.DatesLocations {
// 			opt[key] = " -> concert location"
// 			for _, vals := range element {
// 				opt[vals] = " -> concert date"
// 			}
// 		}
// 	}
// 	return opt
// }
