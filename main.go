package main

import (
	"fmt"
	"groupie-trakcer/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ArtistHandle)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
