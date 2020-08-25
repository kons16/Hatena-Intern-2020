package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in string
	out string
}

func Test_Render(t *testing.T) {
	testCases := []TestCase{
		{
			in:		"# text",
			out:	"<h1>text</h1>\n",
		},
		{
			in:		"## text2",
			out:	"<h2>text2</h2>\n",
		},
		{
			in:		"[Google](https://www.google.co.jp/)",
			out:	"<p><a href=\"https://www.google.co.jp/\">Google</a></p>\n",
		},
		{
			in:		"- list",
			out:	"<ul>\n<li>list</li>\n</ul>\n",
		},
		{
			in:		"[](https://qiita.com/gold-kou/items/a1cc2be6045723e242eb/)",
			out:	"<p><a href=\"https://example.com/\">いまさらだけどgRPCに入門したので分かりやすくまとめてみた - Qiita</a></p>\n",
		},
	}

	for _, testCase := range testCases {
		html, err := Render(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}
