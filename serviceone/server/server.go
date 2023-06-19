package server

import (
	"context"
	"golang-bazel-demo-app/serviceone/pb"
)

type Server struct {
	pb.UnimplementedServiceOneServer // consciously opt-out of forward compatibility
	// other fields
}

func (s *Server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: "Hello, " + req.Name,
	}, nil
}
