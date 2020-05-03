package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	// 404 if not truly root
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, "home.page.html")
	//w.Write([]byte("Hello from Snippetbox"))
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Get requested snippet id
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	// Return requested snippet
	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
	//w.Write([]byte("Display a specific snippet..."))
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
