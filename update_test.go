package convvls

import (
	"testing"
	"time"

	"github.com/loivis/convolvulus-update/c9r"
	"github.com/loivis/convolvulus-update/mem"
	"github.com/loivis/convolvulus-update/mock"
)

func TestService_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		left := map[string]c9r.Left{
			"l1": &mock.Left{
				UpdateFunc: func(*c9r.Book) {},
			},
			"l2": &mock.Left{
				UpdateFunc: func(*c9r.Book) {},
			},
		}
		right := map[string]c9r.Right{
			"r1": &mock.Right{
				UpdateFunc: func(b *c9r.Book) bool {
					b.Update = time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC)
					return true
				},
			},
			"r2": &mock.Right{UpdateFunc: func(*c9r.Book) bool { return false }},
		}
		store := mem.NewStore()

		s := &service{
			Left:  left,
			Right: right,
			Store: store,
		}

		b := c9r.Book{Site: "r1"}
		err := s.update(b)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		wantBook := c9r.Book{Site: "r1", Update: time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC)}

		if got, want := *s.Store.Get(&b), wantBook; got != want {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})
}
