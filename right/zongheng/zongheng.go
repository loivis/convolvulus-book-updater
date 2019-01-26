package zongheng

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/loivis/convolvulus-book-updater/c9r"
	"github.com/loivis/convolvulus-book-updater/httpx"
)

// Site .
type Site struct {
	name string
}

// New returns an instance of Site
func New(opts ...func(*Site)) *Site {
	s := &Site{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithName .
func WithName(n string) func(*Site) {
	return func(s *Site) {
		s.name = n
	}
}

// Update .
func (s *Site) Update(t time.Time) []*c9r.Book {
	var books []*c9r.Book

	for i := 1; i < 33; i++ {
		url := fmt.Sprintf("http://book.zongheng.com/store/c0/c0/b0/u0/p%d/v9/s9/t0/u0/i0/ALL.html", i)
		fmt.Println(url)
		resp, err := httpx.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		newBooks, isLast := s.findNewBooks(resp.Body, t)
		books = append(books, newBooks...)

		resp.Body.Close()

		if isLast {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return books
}

func (s *Site) findNewBooks(reader io.Reader, t time.Time) (books []*c9r.Book, isLast bool) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	doc.Find("div.store_booklist").Find("ul.main_con").Find("li").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		nameSel := sel.Find("span.bookname").Find("a")
		title := strings.TrimSpace(nameSel.Text())
		link, _ := nameSel.Attr("href")
		id := strings.Trim(strings.Split(link, "/")[4], ".html")
		author := strings.TrimSpace(sel.Find("span.author").Text())
		update := c9r.ParseDate(sel.Find("span.time").Text())

		b := &c9r.Book{
			Title:  title,
			ID:     id,
			Site:   s.name,
			Author: author,
			Update: update,
		}
		fmt.Printf("%v", b)

		books = append(books, b)

		if update.Before(t.Add(-time.Minute)) {
			isLast = true
			return false
		}

		return true
	})

	return
}
