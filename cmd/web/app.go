package main

import (
	"snippetbox.org/pkg/models"
	"github.com/alexedwards/scs"
)

type App struct {
	HTMLDir string
	StaticDir string
	Database *models.Database
	Sessions *scs.SessionManager
}
