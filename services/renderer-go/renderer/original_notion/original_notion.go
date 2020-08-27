package original_notion

import (
	"context"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"regexp"
	"strings"
)

type OriginalNotion struct {
	fc pb_fetcher.FetcherClient
	ctx context.Context
}

func NewOriginalNotion(fc pb_fetcher.FetcherClient, ctx context.Context) *OriginalNotion {
	return &OriginalNotion{
		fc:	fc,
		ctx: ctx,
	}
}

// [](url) で url からタイトルをセット
func (on *OriginalNotion) SetTitle(src string) (string, error) {
	r := regexp.MustCompile(`\[\]\((.+?)\)`)
	results := r.FindAllStringSubmatch(src, -1)

	for _, result := range results {
		// fetcherCli.Fetcherより、urlからtitleを取得
		url := result[1]
		reply, err := on.fc.Fetcher(on.ctx, &pb_fetcher.FetcherRequest{Url: url})
		if err != nil {
			return url, err
		}

		set := "[" + reply.Title + "]" + "(" + url + ")"
		target := "[](" + url + ")"
		src = strings.Replace(src, target, set, 1)
	}

	return src, nil
}

// {色指定}(msg) で msg を指定の色に変更
func (on *OriginalNotion) SetColor(src string) (string, error) {
	r2 := regexp.MustCompile(`\{(.+?)\}\((.+?)\)`)
	resultOriginals := r2.FindAllStringSubmatch(src, -1)

	for _, resultOriginal := range resultOriginals {
		color := resultOriginal[1]
		msg := resultOriginal[2]
		colorTagMsg := "<span style=\"color:" + color + "\">" + msg + "</span>"

		target := "{" + color + "}" + "(" + msg + ")"
		src = strings.Replace(src, target, colorTagMsg, 1)
	}

	return src, nil
}

// %msg% で msg を msgのwiki へリンクさせる
func (on *OriginalNotion) SetWikiLink(src string) (string, error) {
	r2 := regexp.MustCompile(`\%(.+?)\%`)
	resultOriginals := r2.FindAllStringSubmatch(src, -1)

	wikiBase := "https://ja.wikipedia.org/wiki/"

	for _, resultOriginal := range resultOriginals {
		msg := resultOriginal[1]
		setUrl := "[" + msg + "]" + "(" + wikiBase + msg + ")"

		target := "%" + msg + "%"
		src = strings.Replace(src, target, setUrl, 1)
	}

	return src, nil
}
