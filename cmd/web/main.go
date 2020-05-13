package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	"snippetbox.org/pkg/models"

	"github.com/alexedwards/scs"
	_ "github.com/go-sql-driver/mysql"
)

var sessionManager *scs.SessionManager

func main() {
	// Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", "sbx:redpong13@/snippetbox?parseTime=true", "MySQL DSN")

	flag.Parse()

	// Database connection
	db := connect(*dsn)
	defer db.Close()

	// Initalize session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 6 * time.Hour
	//sessionManager.Persist(true)

	app := &App{
		HTMLDir: *htmlDir,
		StaticDir: *staticDir,
		Database: &models.Database{db},
		Sessions: sessionManager,
	}

	//Start server, quit on failure
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, sessionManager.LoadAndSave(app.Routes()))
	log.Fatal(err)
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
