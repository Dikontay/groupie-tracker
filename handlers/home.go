package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	
	artistsCh := make(chan []Artist, 1)
	defer close(artistsCh)
	go func(){
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
	
	

	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	select {
	case artists := <-artistsCh:
		if artists == nil {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			err = ts.Execute(w, artists)
			if err != nil {
				http.Error(w, "internal server error", 500)
				return
			}
		}
	case <-time.After(3*time.Second):
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
		return

	}
	
}
