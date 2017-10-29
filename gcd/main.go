package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	context "golang.org/x/net/context"

	"github.com/daniel-dsouza/microtest/pb"
)

type server struct{}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	fmt.Println("hi there")
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}

	return &pb.GCDResponse{Result: a}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
