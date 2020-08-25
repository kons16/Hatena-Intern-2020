package renderer

import (
	"bytes"
	"context"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/config"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"regexp"
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
	conf, _ := config.Load()

	// [](URL) からURLのみを正規表現で抽出
	r := regexp.MustCompile(`\[\]\((.+?)\)`)
	results := r.FindAllStringSubmatch(src, -1)
	if len(results) != 0 && results[0][0] != "" {
		// Fetcher(URLからタイトル取得) サービスに接続
		fetcherConn, _ := grpc.Dial(conf.FetcherAddr, grpc.WithInsecure(), grpc.WithBlock())
		defer fetcherConn.Close()
		fetcherCli := pb_fetcher.NewFetcherClient(fetcherConn)

		// fetcherより、urlからtitleを取得
		url := results[0][1]
		reply, err := fetcherCli.Fetcher(ctx, &pb_fetcher.FetcherRequest{Url: url})
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
