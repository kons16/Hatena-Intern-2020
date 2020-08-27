package grpc

import (
	"context"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/renderer"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Render(t *testing.T) {
	ra := renderer.NewRenderApp(nil)
	s := NewServer(ra)
	src := `foo https://google.com/ bar`
	reply, err := s.Render(context.Background(), &pb.RenderRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, "<p>foo <a href=\"https://google.com/\">https://google.com/</a> bar</p>\n", reply.Html)
}
