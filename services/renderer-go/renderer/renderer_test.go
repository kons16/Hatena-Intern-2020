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
			in:		"aaa {red}(赤色) aaa",
			out: 	"<p>aaa <a style=\"color:red\">赤色</a> aaa</p>\n",
		},

	}

	ra := &RenderApp{nil}
	for _, testCase := range testCases {
		html, err := ra.Render(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}
