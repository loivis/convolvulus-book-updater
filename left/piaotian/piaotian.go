package piaotian

import (
	"strings"

	"github.com/loivis/convolvulus-update/update"
	"github.com/loivis/convolvulus-update/http"
)

type Site struct {
	home        string
	name        string
	chapterLink string
}

// New returns an instance of Site
func New(opts ...func(*Site)) *Site {
	s := &Site{
		home:        "https://www.ptwxz.com",
		name:        "飘天文学网",
		chapterLink: "https://www.ptwxz.com/html/",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Site) Find(name string) update.Source {
	source := update.Source{Site: s.name}

	q := name + " site:" + s.home
	source.ChapterLink = s.getChapterLink(http.Search(q))

	return source
}

func (s *Site) getChapterLink(link string) string {
	if link == "" {
		return ""
	}

	link = strings.Trim(link, ".html")
	ss := strings.Split(link, "/")

	if len(ss) < 6 {
		return ""
	}

	return s.chapterLink + ss[4] + "/" + ss[5] + "/"
}
