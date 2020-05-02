package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// 404 if not truly root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// define html templates
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.page.html",
	}

	// parse html templates
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// render template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//w.Write([]byte("Hello from Snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Get requested snippet id
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Return requested snippet
	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
	//w.Write([]byte("Display a specific snippet..."))
}

func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
