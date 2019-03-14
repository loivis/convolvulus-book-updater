package qidian

import (
	"errors"
	"fmt"
	"time"

	"github.com/loivis/convolvulus-update/c9r"
	"github.com/loivis/convolvulus-update/http"
)

// Site .
type Site struct {
	name        string
	chapterLink string
}

// New returns an instance of Site
func New(opts ...func(*Site)) *Site {
	s := &Site{
		name:        "起点中文网",
		chapterLink: "https://book.qidian.com/info/%s#Catalog",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithChapterLink(l string) func(*Site) {
	return func(s *Site) {
		s.chapterLink = l
	}
}

// Update .
func (s *Site) Update(b *c9r.Book) error {
	url := fmt.Sprintf(s.chapterLink, b.ID)
	doc, err := http.GetDoc(url)
	if err != nil {
		return err
	}

	a := doc.Find("div.volume").Last().Find("li").Last().Find("a")
	// 首发时间：2019-03-14 13:12:54 章节字数：2673
	str, _ := a.Attr("title")

	update, err := parseUpdate(str)
	if err != nil {
		return err
	}

	b.Update = update.UTC()

	return nil
}

func parseUpdate(str string) (time.Time, error) {
	// 首发时间：2019-03-14 13:12:54 章节字数：2673
	var update time.Time
	if len(str) < 50 {
		return update, errors.New("invalid update string")
	}

	update, err := time.ParseInLocation("2006-01-02 15:04:05", str[15:34], c9r.Location)
	if err != nil {
		return update, err
	}

	return update, nil
}
