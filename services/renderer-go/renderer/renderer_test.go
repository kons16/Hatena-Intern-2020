package renderer

import (
	"context"
	"github.com/golang/mock/gomock"
	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/renderer/mock_fetcher"
	"github.com/stretchr/testify/assert"
	"testing"
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
			in:		"[Google](https://www.google.com/)",
			out:	"<p><a href=\"https://www.google.com/\">Google</a></p>\n",
		},
		{
			in:		"[](https://www.google.com/)",
			out:	"<p><a href=\"https://www.google.com/\">Google</a></p>\n",
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFetcher := mock_fetcher.NewMockFetcherClient(ctrl)
	reply := &pb_fetcher.FetcherReply{Title: "Google"}
	mockFetcher.EXPECT().Fetcher(context.Background(), &pb_fetcher.FetcherRequest{Url: "https://www.google.com/"}).Return(reply, nil)

	ra := &RenderApp{fetcherClient: mockFetcher}
	for _, testCase := range testCases {
		html, err := ra.Render(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}
