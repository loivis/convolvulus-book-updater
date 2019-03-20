package mem

import (
	"context"
	"fmt"
	"testing"

	"github.com/loivis/convolvulus-update/update"
)

func TestStore_Get(t *testing.T) {
	ctx := context.Background()
	b := &update.Book{ID: "bar", Site: "foo"}
	s := &Store{
		books: map[string]update.Book{
			"foo-bar": *b,
		},
	}

	t.Run("Success", func(t *testing.T) {
		gotBook, _ := s.Get(ctx, b)
		if got, want := gotBook, b; !(got).Equals(want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})

	t.Run("NonExisting", func(t *testing.T) {
		gotBook, _ := s.Get(ctx, &update.Book{ID: "baz"})
		wantBook := &update.Book{}
		if got, want := gotBook, wantBook; !(got).Equals(want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})
}

func TestStore_Put(t *testing.T) {
	ctx := context.Background()
	b := &update.Book{ID: "bar", Site: "foo"}
	s := NewStore()

	t.Run("Success", func(t *testing.T) {
		s.Put(ctx, b)

		gotBook, _ := s.Get(ctx, b)

		if got, want := gotBook, b; !(got).Equals(want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})
}

func BenchmarkGetPut(b *testing.B) {
	ctx := context.Background()
	s := NewStore()

	for n := 0; n < b.N; n++ {
		go func(n int) {
			s.Get(ctx, &update.Book{ID: fmt.Sprintf("%d", n)})
		}(n)
		go func(n int) {
			s.Put(ctx, &update.Book{ID: fmt.Sprintf("%d", n)})
		}(n)
	}
}

func Test_Key(t *testing.T) {
	b := update.Book{ID: "bar", Site: "foo"}

	if got, want := key(&b), "foo-bar"; got != want {
		t.Fatalf("got key = %q, want %q", got, want)
	}
}
