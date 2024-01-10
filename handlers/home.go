package handlers

import (
	"fmt"
	"groupie-trakcer/singleton"
	"log"
	"net/http"
	"text/template"
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

	// dataCh := make(chan *singleton.Data)

	// go func() {
	// 	defer close(dataCh)
	// 	data, err := singleton.GetAllData()
	// 	fmt.Println(data)
	// 	if err != nil {
	//         log.Printf("Error fetching data: %v", err)
	//         dataCh <- nil
	//         return
	//     }
	//     dataCh <- data
	// }()

	ts, err := template.ParseFiles("./static/templates/index.html")
	if err != nil {

		fmt.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	data, err := singleton.GetAllData()
	if err != nil {
		log.Printf("Getting all data error %e", err)
		errorHandler(w, http.StatusInternalServerError)
	}

	ts.Execute(w, data.AllArtists)

	// select {
	// case info := <-dataCh:
	// 	if info == nil {
	// 		log.Println("Can't get artists")
	// 		errorHandler(w, http.StatusNotFound)
	// 	} else {
	// 		err = ts.Execute(w, info.AllArtists)
	// 		if err != nil {
	// 			errorHandler(w, http.StatusInternalServerError)
	// 			return
	// 		}
	// 	}
	// case <-time.After(3 * time.Second):
	// 	log.Println("Can't get artists from channel")
	// 	errorHandler(w, http.StatusRequestTimeout)
	// 	return

	// }
}
