package qidian

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
		url := fmt.Sprintf("https://www.qidian.com/all?orderId=5&style=2&pageSize=50&siteid=1&pubflag=0&hiddenField=0&page=%d", i)
		fmt.Println(url)
		resp, err := httpx.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		newBooks, isLast := s.findNewBooks(resp.Body, t)

		books = append(books, newBooks...)

		resp.Body.Close()

		if len(newBooks) == 0 {
			fmt.Printf("failed to find books on page %d for %s", i, s.name)
			break
		}

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

	doc.Find("div.all-book-list").Find("tbody").Find("tr").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		nameSel := sel.Find("a.name")
		title := strings.TrimSpace(nameSel.Text())
		id, _ := nameSel.Attr("data-bid")
		author := sel.Find("a.author").Text()
		updateText := sel.Find("td.date").Text()
		update := c9r.ParseDate(updateText)
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

	if len(books) == 0 {
		isLast = true
		return
	}

	return
}
