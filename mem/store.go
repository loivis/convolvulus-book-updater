package mem

import (
	"fmt"
	"sync"

	"github.com/loivis/convolvulus-update/c9r"
)

type Store struct {
	booksMu sync.Mutex
	books   map[string]c9r.Book
}

func NewStore() *Store {
	return &Store{
		books: make(map[string]c9r.Book),
	}
}

func (s *Store) Get(b *c9r.Book) *c9r.Book {
	s.booksMu.Lock()
	out := s.books[key(b)]
	s.booksMu.Unlock()
	return &out
}

func (s *Store) Put(b *c9r.Book) error {
	s.booksMu.Lock()
	s.books[key(b)] = *b
	s.booksMu.Unlock()
	return nil
}

func key(b *c9r.Book) string {
	return fmt.Sprintf("%s-%s", b.Site, b.ID)
}