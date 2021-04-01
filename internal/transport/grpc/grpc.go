package grpc

import (
	"FBSTestTask/internal/service"
	pb "FBSTestTask/internal/transport/grpc/proto"
	"context"
)

type Server struct {
	svc service.FibonacciService
	pb.UnimplementedFibonacciServer
}

func NewServer(fibSvc service.FibonacciService) *Server {
	return &Server{
		svc:                          fibSvc,
		UnimplementedFibonacciServer: pb.UnimplementedFibonacciServer{},
	}
}

func (s *Server) GetFibSlice(ctx context.Context, in *pb.FibonacciRequest) (*pb.FibonacciResponse, error) {
	result, err := s.svc.GetFibSlice(ctx, int(in.Start), int(in.End))
	if err != nil {
		return nil, err
	}
	return &pb.FibonacciResponse{Message: result}, nil
}
