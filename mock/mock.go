package mock

import (
	"github.com/loivis/convolvulus-update/c9r"
)

// Left is a mock of c9r.Left
type Left struct {
	FindFunc func(string) c9r.Source
}

// Find .
func (s *Left) Find(name string) c9r.Source {
	return s.FindFunc(name)
}

// Right is a mock of c9r.Right
type Right struct {
	UpdateFunc func(b *c9r.Book) error
}

// Update .
func (s *Right) Update(b *c9r.Book) error {
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
