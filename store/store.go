package store

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/loivis/convolvulus-update/update"
)

type Store struct {
	collection string
}

var client *firestore.Client

func init() {
	var err error
	client, err = firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("failed to create firestore client: %v", err)
	}
}

func New() *Store {
	return &Store{
		collection: "books",
	}
}

func (s *Store) Get(ctx context.Context, b *update.Book) (*update.Book, error) {
	snap, err := client.Collection(s.collection).Doc(b.DocID()).Get(ctx)
	if err != nil {
		log.Printf("failed to get book(%+v): %v", b, err)
		return nil, err
	}

	var ret *update.Book
	if err := snap.DataTo(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Store) Put(ctx context.Context, b *update.Book) error {
	id := b.DocID()

	_, err := client.Collection(s.collection).Doc(id).Set(ctx, b)
	if err != nil {
		log.Printf("failed to add/update book(%+v): %v", b, err)
		return err
	}

	log.Printf("book added/updated: %s(%+v)", id, b)

	return nil
}
