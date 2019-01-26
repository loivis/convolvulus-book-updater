package right

import (
	"context"
	"time"

	"github.com/loivis/convolvulus-book-updater/c9r"
)

// Searcher .
type Searcher struct {
	logf  func(s string, v ...interface{})
	sites map[string]c9r.Site
}

// New returns an instance of Searcher
func New(sites map[string]c9r.Site, logf func(s string, v ...interface{})) *Searcher {
	return &Searcher{
		sites: sites,
		logf:  logf,
	}
}

func (s *Searcher) hasSite(site string) bool {
	_, ok := s.sites[site]
	return ok
}

// UpdateSite .
func (s *Searcher) UpdateSite(ctx context.Context, site string, t time.Time) []*c9r.Book {
	if s.hasSite(site) {
		return s.sites[site].Update(t)
	}

	return nil
}
