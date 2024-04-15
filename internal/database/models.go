package database

import "time"

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}
