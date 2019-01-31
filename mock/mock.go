package mock

import (
	"github.com/loivis/convolvulus-update/c9r"
)

// Left is a mock of c9r.Left
type Left struct {
	UpdateFunc func(b *c9r.Book)
}

// Update .
func (s *Left) Update(b *c9r.Book) {
	s.UpdateFunc(b)
}

// Right is a mock of c9r.Right
type Right struct {
	UpdateFunc func(b *c9r.Book) bool
}

// Update .
func (s *Right) Update(b *c9r.Book) bool {
	return s.UpdateFunc(b)
}

type Store struct {
	GetFunc func(*c9r.Book) *c9r.Book
	PutFunc func(*c9r.Book) error
}

func (s *Store) Get(b *c9r.Book) *c9r.Book {
	return s.GetFunc(b)
}

func (s *Store) Put(b *c9r.Book) error {
	return s.PutFunc(b)
}
