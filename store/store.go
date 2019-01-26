package store

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/loivis/convolvulus-book-updater/c9r"
)

type Store struct {
	batchSize int
	kind      string
	projectID string
}

func New(batchSize int, kind, projectID string) *Store {
	return &Store{
		batchSize: batchSize,
		kind:      kind,
		projectID: projectID,
	}
}

func (s *Store) LatestSiteUpdate(ctx context.Context, site string) (time.Time, error) {
	client, err := datastore.NewClient(ctx, s.projectID)
	if err != nil {
		return time.Time{}, err
	}

	query := datastore.NewQuery(s.kind).Filter("Site =", site).Order("-Update").Limit(1)
	var books []*c9r.Book
	_, err = client.GetAll(ctx, query, &books)
	if err != nil {
		return time.Time{}, err
	}

	if len(books) == 0 {
		return time.Now().Add(-3 * time.Hour).UTC().Truncate(time.Second), nil
	}

	if time.Since(books[0].Update) > 3*time.Hour {
		return time.Now().Add(-3 * time.Hour).UTC().Truncate(time.Second), nil
	}

	return books[0].Update.UTC().Truncate(time.Second), nil
}

func (s *Store) PutAll(ctx context.Context, books []*c9r.Book) error {
	client, err := datastore.NewClient(ctx, s.projectID)
	if err != nil {
		return err
	}

	for _, books := range s.splitBooks(books) {
		var keys []*datastore.Key
		for _, book := range books {
			keys = append(keys, datastore.NameKey(s.kind, book.Site+book.Author+book.Title, nil))
		}

		if _, err := client.PutMulti(ctx, keys, books); err != nil {
			return err
		}
	}

	return nil
}

// https://github.com/golang/go/wiki/SliceTricks#batching-with-minimal-allocation
func (s *Store) splitBooks(books []*c9r.Book) [][]*c9r.Book {
	var batches [][]*c9r.Book
	for s.batchSize < len(books) {
		books, batches = books[s.batchSize:], append(batches, books[0:s.batchSize:s.batchSize])
	}
	batches = append(batches, books)

	return batches
}
