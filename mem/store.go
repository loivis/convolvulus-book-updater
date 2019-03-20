package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/loivis/convolvulus-update/update"
)

type Store struct {
	booksMu sync.Mutex
	books   map[string]update.Book
}

func NewStore() *Store {
	return &Store{
		books: make(map[string]update.Book),
	}
}

func (s *Store) Get(ctx context.Context, b *update.Book) (*update.Book, error) {
	s.booksMu.Lock()
	out := s.books[key(b)]
	s.booksMu.Unlock()
	return &out, nil
}

func (s *Store) Put(ctx context.Context, b *update.Book) error {
	s.booksMu.Lock()
	s.books[key(b)] = *b
	s.booksMu.Unlock()
	return nil
}

func key(b *update.Book) string {
	return fmt.Sprintf("%s-%s", b.Site, b.ID)
}
