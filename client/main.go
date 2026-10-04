package main

import (
	pb "JRPC/internal/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	client := pb.NewCalculateClient(c)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sumRes, err := client.Sum(
		ctx,
		&pb.Request{
			A: 10,
			B: 20,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sum:", sumRes.Value)

	minusRes, err := client.Minus(
		ctx,
		&pb.Request{
			A: 45,
			B: 27,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Minus:", minusRes.Value)

}
