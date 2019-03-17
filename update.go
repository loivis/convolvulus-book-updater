package convvls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/loivis/convolvulus-update/left/piaotian"
	"github.com/loivis/convolvulus-update/store"

	"github.com/loivis/convolvulus-update/c9r"
	"github.com/loivis/convolvulus-update/right/qidian"
)

var svc *service

func init() {
	svc = &service{
		Right: map[string]c9r.Right{
			"起点中文网": qidian.New(),
		},
		Left: map[string]c9r.Left{
			"飘天文学网": piaotian.New(),
		},
		Store: store.New(),
	}
}

type Message struct {
	Data []byte `json:"data"`
}

func Update(ctx context.Context, m Message) error {
	var books []*c9r.Book

	if err := json.Unmarshal(m.Data, &books); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(books))

	for _, book := range books {
		go func(b *c9r.Book) {
			svc.update(b)
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

func (svc *service) update(b *c9r.Book) error {
	ctx := context.Background()

	if _, ok := svc.Right[b.Site]; !ok {
		return fmt.Errorf("%q doesn't exist", b.Site)
	}

	err := svc.Right[b.Site].Update(b)
	if err != nil {
		return err
	}

	// TODO: update missing and remove non-existing
	var wg sync.WaitGroup
	wg.Add(len(svc.Left))

	for _, site := range svc.Left {
		go func(site c9r.Left) {
			defer wg.Done()

			source := site.Find(b.Title)
			if source.ChapterLink == "" {
				return
			}

			b.SourcesMu.Lock()
			b.Sources = append(b.Sources, &source)
			b.SourcesMu.Unlock()
		}(site)
	}

	wg.Wait()

	if err := svc.Store.Put(ctx, b); err != nil {
		return err
	}

	b, err = svc.Store.Get(ctx, b)
	if err != nil {
		return nil
	}

	for _, source := range b.Sources {
		log.Println(source)
	}

	return nil
}
