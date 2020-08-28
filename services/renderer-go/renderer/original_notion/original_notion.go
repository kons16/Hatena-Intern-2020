package original_notion

import (
	"context"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/patrickmn/go-cache"
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
func (on *OriginalNotion) SetTitle(src string, c *cache.Cache) (string, error) {
	r := regexp.MustCompile(`\[\]\((.+?)\)`)
	results := r.FindAllStringSubmatch(src, -1)

	for _, result := range results {
		url := result[1]
		getCacheTitle, found := c.Get(url)

		if found {
			set := "[" + getCacheTitle.(string) + "]" + "(" + url + ")"
			target := "[](" + url + ")"
			src = strings.Replace(src, target, set, 1)
		} else {
			// cacheにない場合、fetcherCli.Fetcherより、urlからtitleを取得
			reply, err := on.fc.Fetcher(on.ctx, &pb_fetcher.FetcherRequest{Url: url})
			if err != nil {
				return url, err
			}
			c.Set(url, reply.Title, cache.DefaultExpiration)

			set := "[" + reply.Title + "]" + "(" + url + ")"
			target := "[](" + url + ")"
			src = strings.Replace(src, target, set, 1)
		}
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
