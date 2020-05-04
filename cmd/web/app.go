package main

import (
	"snippetbox.org/pkg/models"
)

type App struct {
	HTMLDir string
	StaticDir string
	Database *models.Database
}
