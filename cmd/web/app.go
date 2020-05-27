package main

import (
	"github.com/gorilla/sessions"
	"snippetbox.org/pkg/models"
)

// App structure
type App struct {
	HTMLDir   string
	StaticDir string
	Database  *models.Database
	Sessions  *sessions.CookieStore
}
