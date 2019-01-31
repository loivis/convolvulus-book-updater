package store

import (
	"testing"

	"github.com/loivis/convolvulus-update/c9r"
)

func TestStore_New(t *testing.T) {
	bs := 10
	kind := "foo"
	projectID := "bar"

	s := New(bs, kind, projectID)

	if got, want := s.batchSize, bs; got != want {
		t.Errorf("s.kind = %q, want %q", got, want)
	}

	if got, want := s.kind, kind; got != want {
		t.Errorf("s.kind = %q, want %q", got, want)
	}

	if got, want := s.projectID, projectID; got != want {
		t.Errorf("s.projectID = %q, want %q", got, want)
	}
}

func TestStore_SplitBooks(t *testing.T) {
	b1 := &c9r.Book{Title: "b1"}
	b2 := &c9r.Book{Title: "b2"}
	b3 := &c9r.Book{Title: "b3"}

	for _, tc := range []struct {
		desc string
		in   []*c9r.Book
		out  [][]*c9r.Book
	}{
		{
			desc: "Empty",
			in:   []*c9r.Book{},
			out:  [][]*c9r.Book{[]*c9r.Book{}},
		},
		{
			desc: "More",
			in:   []*c9r.Book{b1, b2, b3, b2, b1, b3, b3, b2},
			out: [][]*c9r.Book{
				[]*c9r.Book{b1, b2, b3},
				[]*c9r.Book{b2, b1, b3},
				[]*c9r.Book{b3, b2},
			},
		},
	} {
		s := New(3, "", "")
		split := s.splitBooks(tc.in)

		if got, want := len(split), len(tc.out); got != want {
			t.Fatalf("got batches length = %d, want %d", got, want)
		}

		for i := range split {
			if got, want := len(split[i]), len(tc.out[i]); got != want {
				t.Fatalf("batches[%d] length = %d, want %d", i, got, want)
			}

			for j := range split[i] {
				if got, want := split[i][j], tc.out[i][j]; got != want {
					t.Fatalf("batches[%d][%d] = %v, want %v", i, j, got, want)
				}
			}
		}
	}

}
