package renderer

import (
	"bytes"
	"context"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"html/template"
	"log"
	"regexp"
)

type RenderApp struct {
	fetcherClient	pb_fetcher.FetcherClient
}

// NewRenderApp は RenderApp を作成する
func NewRenderApp(
	fetcherClient	pb_fetcher.FetcherClient,
) *RenderApp {
	return &RenderApp{fetcherClient}
}

var urlRE = regexp.MustCompile(`https?://[^\s]+`)
var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

var md = goldmark.New(
			goldmark.WithExtensions(extension.Linkify),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		)

// Render は受け取った文書を HTML に変換する
func (ra *RenderApp) Render(ctx context.Context, src string) (string, error) {
	// [](URL) からURLのみを正規表現で抽出
	r := regexp.MustCompile(`\[\]\((.+?)\)`)
	results := r.FindAllStringSubmatch(src, -1)
	if len(results) != 0 && results[0][0] != "" {
		url := results[0][1]

		// fetcherCli.Fetcherより、urlからtitleを取得
		reply, err := ra.fetcherClient.Fetcher(ctx, &pb_fetcher.FetcherRequest{Url: url})
		if err != nil {
			return "", err
		}

		inputTitle := "[" + reply.Title + "]"
		r2 := regexp.MustCompile(`\[\]`)
		src = r2.ReplaceAllString(src, inputTitle)
	}

	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
	  log.Fatal(err)
	}

	return buf.String(), nil
}
