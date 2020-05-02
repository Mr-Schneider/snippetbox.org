package main

import (
	"log"
	"net/http"
)

func main() {
	// Initalize server mux
	mux := http.NewServeMux()

	// Paths
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// File server for html assets
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Start server, quit on failure
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


