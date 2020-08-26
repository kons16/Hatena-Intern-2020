package fetcher

import (
	"context"
	"net/http"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Fetcher は受け取ったURLからタイトルを取得する
func Fetcher(ctx context.Context, url string) (string, error) {
	fmt.Println(url)
	res, err := http.Get(url)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	t := doc.Find("title").Text()
	return t, nil
}
