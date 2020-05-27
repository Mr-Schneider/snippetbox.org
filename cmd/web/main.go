package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"snippetbox.org/pkg/models"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var sessionStore *sessions.CookieStore

func main() {
	// Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", "postgres://", "Postgres DSN")

	flag.Parse()

	// Database connection
	db := connect(*dsn)
	defer db.Close()

	// Initalize session manager
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	// Application instance
	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
		Database:  &models.Database{db},
		Sessions:  sessionStore,
	}

	//Start server, quit on failure
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, LogRequest(SecureHeaders(app.Routes())))
	log.Fatal(err)
}

// connect DB connection setup
func connect(dsn string) *sql.DB {
	// Postgres
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
