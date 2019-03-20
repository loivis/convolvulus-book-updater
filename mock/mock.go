package mock

import (
	"github.com/loivis/convolvulus-update/update"
)

// Left is a mock of update.Left
type Left struct {
	FindFunc func(string) update.Source
}

// Find .
func (s *Left) Find(name string) update.Source {
	return s.FindFunc(name)
}

// Right is a mock of update.Right
type Right struct {
	UpdateFunc func(b *update.Book) error
}

// Update .
func (s *Right) Update(b *update.Book) error {
	return s.UpdateFunc(b)
}

type Store struct {
	GetFunc func(*update.Book) *update.Book
	PutFunc func(*update.Book) error
}

func (s *Store) Get(b *update.Book) *update.Book {
	return s.GetFunc(b)
}

func (s *Store) Put(b *update.Book) error {
	return s.PutFunc(b)
}
