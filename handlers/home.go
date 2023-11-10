package handlers

import (
	"fmt"
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

	ts, err := template.ParseFiles("./ui/index.html")
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
}
