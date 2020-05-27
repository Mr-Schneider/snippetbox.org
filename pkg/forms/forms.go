package forms

import (
	"strings"
	"unicode/utf8"
)

// NewSnippet models the snippet structure
type NewSnippet struct {
	Title    string
	Content  string
	Expires  string
	Failures map[string]string
}

// Valid makes sure the the fields are correctly formatted
func (f *NewSnippet) Valid() bool {
	f.Failures = make(map[string]string)

	// Check for non-empty title
	if strings.TrimSpace(f.Title) == "" {
		f.Failures["Title"] = "Title is required"
	} else if utf8.RuneCountInString(f.Title) > 100 {
		f.Failures["Title"] = "Title cannot be longer than 100 characters"
	}

	// Check for non-empty content
	if strings.TrimSpace(f.Content) == "" {
		f.Failures["Content"] = "Content is required"
	}

	// Only allow 1 y/d/m expiry
	permitted := map[string]bool{"3600": true, "86400": true, "31536000": true}
	if strings.TrimSpace(f.Expires) == "" {
		f.Failures["Expires"] = "Expiry time is required"
	} else if !permitted[f.Expires] {
		f.Failures["Expires"] = "Expiry time must be 1 year/day/min"
	}

	return len(f.Failures) == 0
}
