package models

import (
	"time"
)

// Snippet describes the snippet structure
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Snippets holds multiple snippets
type Snippets []*Snippet
