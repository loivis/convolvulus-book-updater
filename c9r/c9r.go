package c9r

import (
	"time"
)

// Site .
type Site interface {
	Update(t time.Time) []*Book
}

// Book includes more infomation of a book
type Book struct {
	Author string    `json:"author,omitempty"`
	ID     string    `json:"id,omitempty"`
	Site   string    `json:"site,omitempty"`
	Title  string    `json:"title,omitempty"`
	Update time.Time `json:"update,omitempty"`
}
