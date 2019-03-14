package qidian

import (
	"testing"
	"time"

	"github.com/loivis/convolvulus-update/c9r"
)

func TestNew(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		s := New()

		if s == nil {
			t.Fatal("s is nil")
		}

		if got, want := s.name, "起点中文网"; got != want {
			t.Fatalf("s.name = %q, want %q", got, want)
		}

		if s.chapterLink == "" {
			t.Fatalf("s.chapterLink not set")
		}
	})

	t.Run("WithChapterLink", func(t *testing.T) {
		link := "https://example.org"
		s := New(WithChapterLink(link))

		if got, want := s.chapterLink, link; got != want {
			t.Fatalf("s.chapterLink = %q, want %q", got, want)
		}
	})
}

func TestParseUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		str := "首发时间：2019-03-14 13:12:54 章节字数：2673"

		update, err := parseUpdate(str)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		wantUpdate := time.Date(2019, 3, 14, 13, 12, 54, 0, c9r.Location)

		if got, want := update.String(), wantUpdate.String(); got != want {
			t.Fatalf("got update = %q, want %q", got, want)
		}
	})

	t.Run("ShortTimeString", func(t *testing.T) {
		str := "首发时间：2019-03-14 13:12:54"

		_, err := parseUpdate(str)
		if err == nil {
			t.Fatal("error is nil")
		}

		if got, want := err.Error(), "invalid update string"; got != want {
			t.Fatalf("got update = %q, want %q", got, want)
		}
	})
}
