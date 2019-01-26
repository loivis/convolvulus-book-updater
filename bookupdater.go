package bookupdater

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/loivis/convolvulus-book-updater/c9r"
	"github.com/loivis/convolvulus-book-updater/right"
	"github.com/loivis/convolvulus-book-updater/right/qidian"
	"github.com/loivis/convolvulus-book-updater/right/zongheng"
	"github.com/loivis/convolvulus-book-updater/store"
)

var logger = log.New(os.Stderr, "", 0)

func Update(http.ResponseWriter, *http.Request) {
	qidian := qidian.New(qidian.WithName("起点中文网"))
	zongheng := zongheng.New(zongheng.WithName("纵横中文网"))
	right := right.New(
		map[string]c9r.Site{"起点中文网": qidian, "纵横中文网": zongheng},
		logger.Printf,
	)

	s := store.New(500, "book", "convvls")

	sites := []string{"起点中文网", "纵横中文网"}

	var wg sync.WaitGroup
	wg.Add(len(sites))

	for _, site := range sites {
		go func(site string) {
			update(s, site, right)
			wg.Done()
		}(site)
	}
	wg.Wait()
}

func update(s *store.Store, site string, right *right.Searcher) {
	ctx := context.Background()

	updated, err := s.LatestSiteUpdate(ctx, site)
	if err != nil {
		logger.Println(err)
		return
	}
	logger.Printf("%v was updated at %v\n", site, updated)

	books := c9r.DeduplicateBooks(right.UpdateSite(ctx, site, updated))
	logger.Printf("updating %d books for %s", len(books), site)
	if err := s.PutAll(ctx, books); err != nil {
		logger.Println(err)
	}
}
