package mock

import (
	"time"

	"github.com/loivis/convolvulus-update/c9r"
)

// Site is a mock of c9r.Site
type Site struct {
	UpdateFunc func(t time.Time) []*c9r.Book
}

// Update .
func (s *Site) Update(t time.Time) []*c9r.Book {
	return s.UpdateFunc(t)
}
