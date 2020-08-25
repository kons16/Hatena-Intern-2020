package renderer

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

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
func Render(ctx context.Context, src string) (string, error) {
	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
	  log.Fatal(err)
	}

	return buf.String(), nil
}
