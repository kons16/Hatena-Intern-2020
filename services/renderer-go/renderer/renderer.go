package renderer

import (
	"bytes"
	"context"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/renderer/original_notion"
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
				html.WithUnsafe(),
			),
		)

// Render は受け取った文書を HTML に変換する
func (ra *RenderApp) Render(ctx context.Context, src string) (string, error) {
	on := original_notion.NewOriginalNotion(ra.fetcherClient, ctx)

	// title から url をセット
	src, err := on.SetTitle(src)
	if err != nil {
		return "", err
	}

	// 独自記法で色のセット
	src, err = on.SetColor(src)
	if err != nil {
		return "", err
	}

	// src を markdown に
	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
	  log.Fatal(err)
	}

	return buf.String(), nil
}
