package c9r

import (
	"testing"
	"time"
)

func TestDeduplicateBook(t *testing.T) {
	b1 := &Book{Title: "b1"}
	b2 := &Book{Title: "b2"}
	b3 := &Book{Title: "b3"}
	for _, tc := range []struct {
		desc string
		in   []*Book
		out  []*Book
	}{
		{
			desc: "EmptyInput",
			in:   []*Book{},
			out:  []*Book{},
		},
		{
			desc: "DistinctInput",
			in:   []*Book{b1, b2, b3},
			out:  []*Book{b1, b2, b3},
		},
		{
			desc: "DuplicatedInput",
			in:   []*Book{b1, b2, b1, b3, b2, b3, b1},
			out:  []*Book{b1, b2, b3},
		},
	} {
		gotBooks := DeduplicateBooks(tc.in)
		if got, want := len(gotBooks), len(tc.out); got != want {
			t.Fatalf("got %d books, want %d", got, want)
		}

		for n := range gotBooks {
			if got, want := gotBooks[n], tc.out[n]; got != want {
				t.Fatalf("gotBooks[%d] = %v, want %v", n, got, want)
			}
		}
	}
}

func TestParseUpdate(t *testing.T) {
	for _, tc := range []struct {
		desc string
		in   string
		out  time.Time
	}{
		{
			desc: "DateHourMinuteInThePast",
			in:   "01-01 12:34",
			out:  time.Date(time.Now().Year(), 1, 1, 4, 34, 0, 0, time.UTC),
		},
		{
			desc: "DateHourMinuteInTheFuture",
			in:   "12-31 23:45",
			out:  time.Date(time.Now().AddDate(-1, 0, 0).Year(), 12, 31, 15, 45, 0, 0, time.UTC),
		},
		{
			desc: "MinutesBack",
			in:   "43分钟前",
			out:  time.Now().Add(-43 * time.Minute).UTC().Truncate(time.Minute),
		},
		{
			desc: "HoursBack",
			in:   "2小时前",
			out:  time.Now().Add(-2 * time.Hour).UTC().Truncate(time.Minute),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			if got, want := ParseDate(tc.in), tc.out; got != want {
				t.Fatalf("got = %q, want %q", got, want)
			}
		})
	}
}
