package main

import (
	"fmt"
	"groupie-trakcer/handlers"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	// This line creates a new file server that serves files from the ./static directory.
	// The neuteredFileSystem wrapper is used around the file system to modify its behavior.
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	// Registers a handler that returns a 404 Not Found response for the exact path /static.
	mux.Handle("/static", http.NotFoundHandler())
	// Registers a handler for paths starting with /static/. It strips the /static prefix before the request reaches the fileServer.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.ArtistHandle)
	log.Print("Starting server on : http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Defines a new struct type neuteredFileSystem with a single field fs of type http.FileSystem. This custom file system will be used to modify the behavior of the file server.

type neuteredFileSystem struct {
	fs http.FileSystem
}

//This is a method on neuteredFileSystem that overrides the Open method of http.FileSystem. It's used to control access to the file system.

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	//Tries to open the file or directory at the given path.
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	//Retrieves file information.
	s, err := f.Stat()

	//Checks if the opened path is a directory.
	if s.IsDir() {
		//If it's a directory, constructs a path to an "index.html" file inside this directory.
		index := filepath.Join(path, "index.html")
		//Tries to open the "index.html" file.
		if _, err := nfs.fs.Open(index); err != nil {
			//If "index.html" does not exist, closes the directory file.
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			//If "index.html" does not exist, return the original error.
			return nil, err
		}
	}
	//If everything is fine (either it's not a directory, or "index.html" exists in the directory), return the file.
	return f, nil
}
