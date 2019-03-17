package c9r

import (
	"context"
	"crypto/sha1"
	"fmt"
	"sync"
	"time"
)

// Left .
type Left interface {
	Find(name string) Source
}

// Right
type Right interface {
	Update(b *Book) error
}

type Store interface {
	Get(ctx context.Context, b *Book) (*Book, error)
	Put(ctx context.Context, b *Book) error
}

// Book includes more infomation of a book
type Book struct {
	Author string    `json:"author,omitempty" firestore:"author"`
	ID     string    `json:"id,omitempty" firestore:"id"`
	Site   string    `json:"site,omitempty" firestore:"site"`
	Title  string    `json:"title,omitempty" firestore:"title"`
	Update time.Time `json:"update,omitempty" firestore:"update"`

	SourcesMu sync.Mutex `json:"-" firestore:"-"`
	Sources   []*Source  `json:"sources,omitempty" firestore:"sources"`
}

type Source struct {
	Site        string `json:"site,omitempty" firestore:"site"`
	ChapterLink string `json:"chapter_link,omitempty" firestore:"chapterLink"`
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

func (b *Book) DocID() string {
	h := sha1.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%s-%s-%s", b.Site, b.Author, b.Title)))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
