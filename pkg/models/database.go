package models

import (
	"database/sql"
)

// Database holds the db connection
type Database struct {
	*sql.DB
}

// GetSnippet retrives a snippet from the db
func (db *Database) GetSnippet(id int) (*Snippet, error) {
	// Query statement
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > timezone('utc', now()) AND id = $1`

	// Execute query
	row := db.QueryRow(stmt, id)
	s := &Snippet{}

	// Pull data into snippet
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// LatestSnippets grabs the latest 10 valid snippets
func (db *Database) LatestSnippets() (Snippets, error) {
	// Query statement
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > timezone('utc', now()) ORDER BY created DESC LIMIT 10`

	// Execute query
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := Snippets{}

	// Get all the matching snippets
	for rows.Next() {
		s := &Snippet{}

		// Pull data into snippet
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

// InsertSnippet adds a new snippet to the db
func (db *Database) InsertSnippet(title, content, expires string) (int, error) {
	// Save stored snippet
	var userid int

	// Query statement
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, timezone('utc', now()), timezone('utc', now()) + interval '3600' second) RETURNING id`

	err := db.QueryRow(stmt, title, content).Scan(&userid)
	if err != nil {
		return 0, err
	}

	return userid, nil
}
