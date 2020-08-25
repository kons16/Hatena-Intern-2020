package grpc

import (
	"context"

	"github.com/hatena/Hatena-Intern-2020/services/fetcher-go/fetcher"
	pb "github.com/hatena/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.FetcherServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetcherは受け取ったURLからタイトルを取得する
func (s *Server) Fetcher(ctx context.Context, in *pb.FetcherRequest) (*pb.FetcherReply, error) {
	title, err := fetcher.Fetcher(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	return &pb.FetcherReply{Title: title}, nil
}
