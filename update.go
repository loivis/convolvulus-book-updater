package convvls

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/loivis/convolvulus-update/c9r"
)

var s *service

func init() {
	s = &service{}
}

type Message struct {
	Data []byte `json:"data"`
}

func Update(ctx context.Context, m Message) error {
	var books []c9r.Book

	if err := json.Unmarshal(m.Data, &books); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(books))

	for _, book := range books {
		go func(b c9r.Book) {
			s.update(b)
			wg.Done()
		}(book)
	}

	wg.Wait()

	return nil
}

type service struct {
	Left  map[string]c9r.Left
	Right map[string]c9r.Right
	Store c9r.Store
}

func (s *service) update(b c9r.Book) error {
	if _, ok := s.Right[b.Site]; !ok {
		return errors.New(fmt.Sprintf("%q doesn't exist", b.Site))
	}

	hasUpdate := s.Right[b.Site].Update(&b)

	if !hasUpdate {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(len(s.Left))

	for _, site := range s.Left {
		go func(site c9r.Left) {
			site.Update(&b)
			wg.Done()
		}(site)
	}

	wg.Wait()

	if err := s.Store.Put(&b); err != nil {
		return err
	}

	return nil
}

// var logger = log.New(os.Stderr, "", 0)

// func Update(http.ResponseWriter, *http.Request) {
// 	qidian := qidian.New(qidian.WithName("起点中文网"))
// 	zongheng := zongheng.New(zongheng.WithName("纵横中文网"))
// 	right := right.New(
// 		map[string]c9r.Site{"起点中文网": qidian, "纵横中文网": zongheng},
// 		logger.Printf,
// 	)

// 	s := store.New(500, "book", "convvls")

// 	sites := []string{"起点中文网", "纵横中文网"}

// 	var wg sync.WaitGroup
// 	wg.Add(len(sites))

// 	for _, site := range sites {
// 		go func(site string) {
// 			update(s, site, right)
// 			wg.Done()
// 		}(site)
// 	}
// 	wg.Wait()
// }

// func update(s *store.Store, site string, right *right.Searcher) {
// 	ctx := context.Background()

// 	updated, err := s.LatestSiteUpdate(ctx, site)
// 	if err != nil {
// 		logger.Println(err)
// 		return
// 	}
// 	logger.Printf("%v was updated at %v\n", site, updated)

// 	books := c9r.DeduplicateBooks(right.UpdateSite(ctx, site, updated))
// 	logger.Printf("updating %d books for %s", len(books), site)
// 	if err := s.PutAll(ctx, books); err != nil {
// 		logger.Println(err)
// 	}
// }
