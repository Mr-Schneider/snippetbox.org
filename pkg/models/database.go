package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

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

func (db *Database) InsertSnippet(title, content, expires string) (int, error) {
	// Save stored snippet
	var userid int

	// Query statement
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, timezone('utc', now()), timezone('utc', now()) + interval '3600' second) RETURNING id`

	err := db.QueryRow(stmt, title, content, expires).Scan(&userid)
	if err != nil {
		return 0, err
	}

	return userid, nil
}
