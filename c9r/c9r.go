package c9r

import (
	"time"
)

// Left .
type Left interface {
	Update(*Book)
}

// Right
type Right interface {
	Update(*Book) (hasUpdate bool)
}

type Store interface {
	Get(b *Book) *Book
	Put(b *Book) error
}

// Book includes more infomation of a book
type Book struct {
	Author string    `json:"author,omitempty"`
	ID     string    `json:"id,omitempty"`
	Site   string    `json:"site,omitempty"`
	Title  string    `json:"title,omitempty"`
	Update time.Time `json:"update,omitempty"`
}
