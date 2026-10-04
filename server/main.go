package main

import (
	"context"
	"log"
	"net"

	pb "JRPC/internal/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Sum(
	ctx context.Context,
	req *pb.Request,
) (*pb.Response, error) {

	result := req.A + req.B

	return &pb.Response{
		Value: result,
	}, nil
}

func (s *server) Minus(
	ctx context.Context,
	req *pb.Request,
) (*pb.Response, error) {

	result := req.A - req.B

	return &pb.Response{
		Value: result,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(
		grpcServer,
		&server{},
	)

	log.Println("Server running on :8080")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
