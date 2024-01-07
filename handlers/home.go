package handlers

import (
	"fmt"
	"groupie-trakcer/singleton"

	"log"
	"net/http"
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

	dataCh := make(chan *singleton.Data)

	//artistsCh := make(chan []Artist, 1)

	go func() {
		data := singleton.GetAllData()
		//artists := []Artist{}
		//err := getElement(artistUrl, &artists)
		if data != nil {
			//fmt.Println(err)
			//artistsCh <- nil
			dataCh <- data
		} else {
			dataCh <- nil
		}
		return
	}()

	ts, err := template.ParseFiles("./static/templates/index.html")
	if err != nil {
		// http.Error(w, "internal server error", 500)
		fmt.Println(err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	select {
	case info := <-dataCh:
		if info == nil {
			// http.Error(w, "Not found", http.StatusNotFound)
			log.Println("Can't get artists")
			errorHandler(w, http.StatusNotFound)
		} else {
			err = ts.Execute(w, info.AllArtists)
			if err != nil {
				errorHandler(w, http.StatusInternalServerError)
				return
			}
		}
	case <-time.After(3 * time.Second):
		log.Println("Can't get artists from channel")
		errorHandler(w, http.StatusRequestTimeout)
		return

	}
	defer close(dataCh)
}
