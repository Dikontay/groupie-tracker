package main

import (
	"fmt"
	"groupie-trakcer/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.ArtistHandle)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Print("Starting server on : http://localhost:4000")

}
