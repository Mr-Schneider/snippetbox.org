package main

import (
	"fmt"
	"net/http"
	"strconv"
	"snippetbox.org/pkg/forms"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	// 404 if not truly root
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	// Get the latest snippets
	snippets, err := app.Database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})
	//w.Write([]byte("Hello from Snippetbox"))
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Get requested snippet id
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	// Get snippet
	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if snippet == nil {
		app.NotFound(w)
		return
	}

	// Return requested snippet
	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet,
	})
	//fmt.Fprint(w, snippet)
	//w.Write([]byte("Display a specific snippet..."))
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "new.page.html", &HTMLData{
		Form : &forms.NewSnippet{},
	})
	//w.Write([]byte("Display the new snippet form..."))
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// Parse the post data
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewSnippet{
		Title: r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}

	// Validate form
	if !form.Valid() {
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}

	// Insert the new snippet
	id, err := app.Database.InsertSnippet(form.Title, form.Content, form.Expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	w.Write([]byte("Create a new snippet..."))
}