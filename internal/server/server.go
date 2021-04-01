package server

import (
	grpcLocal "FBSTestTask/internal/transport/grpc"
	pb "FBSTestTask/internal/transport/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	grpcServer *grpc.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		grpcServer: grpc.NewServer([]grpc.ServerOption{}...),
	}
}

func (s *Server) RunHTTP() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) RunGRPC(server grpcLocal.Server) error {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		return err
	}

	pb.RegisterFibonacciServer(s.grpcServer, &server)
	return s.grpcServer.Serve(listener)
}

func (s *Server) StopHTTP(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
