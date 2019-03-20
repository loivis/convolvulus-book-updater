package convvls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/loivis/convolvulus-update/left/piaotian"
	"github.com/loivis/convolvulus-update/store"

	"github.com/loivis/convolvulus-update/update"
	"github.com/loivis/convolvulus-update/right/qidian"
)

var svc *service

func init() {
	svc = &service{
		Right: map[string]update.Right{
			"起点中文网": qidian.New(),
		},
		Left: map[string]update.Left{
			"飘天文学网": piaotian.New(),
		},
		Store: store.New(),
	}
}

type Message struct {
	Data []byte `json:"data"`
}

func Update(ctx context.Context, m Message) error {
	var books []*update.Book

	if err := json.Unmarshal(m.Data, &books); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(books))

	for _, book := range books {
		go func(b *update.Book) {
			svc.update(b)
			wg.Done()
		}(book)
	}

	wg.Wait()

	return nil
}

type service struct {
	Left  map[string]update.Left
	Right map[string]update.Right
	Store update.Store
}

func (svc *service) update(b *update.Book) error {
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
		go func(site update.Left) {
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
