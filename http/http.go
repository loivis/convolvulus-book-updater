package http

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/loivis/gs-google-search"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36"

var gsClient *gs.Client

func init() {
	gsClient = gs.NewClient(gs.WithUserAgent(userAgent))
}

// GetDoc performs http.MethodGet with custom User-Agent and returns goquery document from response body
func GetDoc(url string) (*goquery.Document, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to get %q: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return doc, nil
}

// Search returns the first link from google search
func Search(q string) string {
	query := &url.Values{
		"q":   {q},
		"num": {"5"},
		"hl":  {"zh-CN"},
	}

	res, err := gsClient.Search(query)
	if err != nil {
		log.Println(err)
		return ""
	}

	if len(res) == 0 {
		log.Printf("no search result for %q", q)
		return ""
	}

	return res[0].Link
}
