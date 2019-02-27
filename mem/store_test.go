package mem

import (
	"fmt"
	"testing"

	"github.com/loivis/convolvulus-update/c9r"
)

func TestStore_Get(t *testing.T) {
	b := c9r.Book{ID: "bar", Site: "foo"}
	s := &Store{
		books: map[string]c9r.Book{
			"foo-bar": b,
		},
	}

	t.Run("Success", func(t *testing.T) {
		if got, want := *s.Get(&b), b; !(&got).Equals(&want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})

	t.Run("NonExisting", func(t *testing.T) {
		wantBook := c9r.Book{}
		if got, want := *s.Get(&c9r.Book{ID: "baz"}), wantBook; !(&got).Equals(&want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})
}

func TestStore_Put(t *testing.T) {
	b := c9r.Book{ID: "bar", Site: "foo"}
	s := NewStore()

	t.Run("Success", func(t *testing.T) {
		s.Put(&b)

		if got, want := *s.Get(&b), b; !(&got).Equals(&want) {
			t.Fatalf("got book = %+v, want %+v", got, want)
		}
	})
}

func BenchmarkGetPut(b *testing.B) {
	s := NewStore()

	for n := 0; n < b.N; n++ {
		go func(n int) {
			s.Get(&c9r.Book{ID: fmt.Sprintf("%d", n)})
		}(n)
		go func(n int) {
			s.Put(&c9r.Book{ID: fmt.Sprintf("%d", n)})
		}(n)
	}
}

func Test_Key(t *testing.T) {
	b := c9r.Book{ID: "bar", Site: "foo"}

	if got, want := key(&b), "foo-bar"; got != want {
		t.Fatalf("got key = %q, want %q", got, want)
	}
}
