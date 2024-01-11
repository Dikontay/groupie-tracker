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
	err := http.ListenAndServe(":8080", mux)
	log.Print("Starting server on : http://localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	

}
