package piaotian

import "testing"

func TestParseID(t *testing.T) {
	s := &Site{chapterLink: "https://foo.bar/baz/"}

	for _, tc := range []struct {
		desc string
		in   string
		out  string
	}{
		{
			desc: "FromEmpty",
			in:   "",
			out:  "",
		},
		{
			desc: "FromBookInfo",
			in:   "https://example.org/bookinfo/5/5623.html",
			out:  "https://foo.bar/baz/5/5623/",
		},
		{
			desc: "FromChapterList",
			in:   "https://example.org/html/5/5623/",
			out:  "https://foo.bar/baz/5/5623/",
		},
		{
			desc: "FromChapter",
			in:   "https://example.org/html/5/5623/6894017.html",
			out:  "https://foo.bar/baz/5/5623/",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			if got, want := s.getChapterLink(tc.in), tc.out; got != want {
				t.Fatalf("got id = %q, want %q", got, want)
			}
		})
	}
}
