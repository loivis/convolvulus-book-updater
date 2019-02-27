package c9r

import (
	"time"
)

// Left .
type Left interface {
	Update(*Book) error
}

// Right
type Right interface {
	Update(*Book) error
}

type Store interface {
	Get(b *Book) *Book
	Put(b *Book) error
}

// Book includes more infomation of a book
type Book struct {
	Author  string    `json:"author,omitempty"`
	ID      string    `json:"id,omitempty"`
	Site    string    `json:"site,omitempty"`
	Title   string    `json:"title,omitempty"`
	Update  time.Time `json:"update,omitempty"`
	Sources []Source  `json:"sources,omitempty"`
}

func (b1 *Book) Equals(b2 *Book) bool {
	authorEq := b1.Author == b2.Author
	idEq := b1.ID == b2.ID
	siteEq := b1.Site == b2.Site
	titleEq := b1.Title == b2.Title
	updateEq := b1.Update == b2.Update

	sourcesEq := true
	if len(b1.Sources) != len(b2.Sources) {
		sourcesEq = false
	}

	// TODO: sort sources before range compare
	for i := range b1.Sources {
		if b1.Sources[i] != b2.Sources[i] {
			sourcesEq = false
			break
		}
	}

	return (authorEq && idEq && siteEq && titleEq && updateEq && sourcesEq)
}

type Source struct {
	Site string
	ID   string
}
