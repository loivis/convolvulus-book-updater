package convvls

import (
	"context"
	"testing"
	"time"

	"github.com/loivis/convolvulus-update/update"
	"github.com/loivis/convolvulus-update/mem"
	"github.com/loivis/convolvulus-update/mock"
)

func TestService_Update(t *testing.T) {
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		left := map[string]update.Left{
			"l1": &mock.Left{
				FindFunc: func(string) update.Source { return update.Source{} },
			},
			"l2": &mock.Left{
				FindFunc: func(string) update.Source { return update.Source{} },
			},
		}
		right := map[string]update.Right{
			"r1": &mock.Right{
				UpdateFunc: func(b *update.Book) error {
					b.Update = time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC)
					return nil
				},
			},
			"r2": &mock.Right{UpdateFunc: func(*update.Book) error { return nil }},
		}

		store := mem.NewStore()

		s := &service{
			Left:  left,
			Right: right,
			Store: store,
		}

		b := &update.Book{Site: "r1"}
		err := s.update(b)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		gotBook, _ := s.Store.Get(ctx, b)
		wantBook := &update.Book{Site: "r1", Update: time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC)}

		if got, want := gotBook, wantBook; !(got).Equals(want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})

	t.Run("NonExistingSite", func(t *testing.T) {
		s := &service{Right: make(map[string]update.Right)}
		b := &update.Book{Site: "foo"}

		err := s.update(b)
		if err == nil {
			t.Fatalf("err is nil")
		}

		if got, want := err.Error(), "\"foo\" doesn't exist"; got != want {
			t.Fatalf("got err = %q, want %q", got, want)
		}
	})
}
