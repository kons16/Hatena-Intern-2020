package grpc

import (
	"context"

	pb "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/renderer"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedRendererServer
	healthpb.UnimplementedHealthServer
	ra *renderer.RenderApp
}

// NewServer は gRPC サーバーを作成する
func NewServer(ra *renderer.RenderApp) *Server {
	return &Server{ra:ra}
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {
	html, err := s.ra.Render(ctx, in.Src)
	if err != nil {
		return nil, err
	}
	return &pb.RenderReply{Html: html}, nil
}
