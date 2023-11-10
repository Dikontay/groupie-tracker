package main

import (
	"fmt"
	"groupie-trakcer/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.ArtistHandle)
	mux.HandleFunc("/search", handlers.HandleSearch)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println(err)
	}

}
