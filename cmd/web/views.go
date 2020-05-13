package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
	"snippetbox.org/pkg/models"
)

type HTMLData struct {
	Snippet 	*models.Snippet
	Snippets 	[]*models.Snippet
	Path 		string
	Form		interface{}
	Flash		string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {
	if data == nil {
		data = &HTMLData{}
	}

	// Add the current path to the data
	data.Path = r.URL.Path
	
	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	// Map for custome template functions
	fm := template.FuncMap{
		"humanDate": humanDate,
	}

	// Pull the html files together
	ts, err := template.New("").Funcs(fm).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Write template to buffer, then send buffer
	buf := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
