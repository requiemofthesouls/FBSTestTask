package grpc

import (
	pb "FBSTestTask/internal/transport/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "127.0.0.1:5300"
)

func TestGRPCFibonacciDial(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFibonacciClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetFibSlice(ctx, &pb.FibonacciRequest{Start: 0, End: 100001})
	if err != nil {
		log.Fatalf("fail to call fibonacci service: %v", err)
	}
	log.Printf("Fibonacci response: %s", r.GetMessage())

}
