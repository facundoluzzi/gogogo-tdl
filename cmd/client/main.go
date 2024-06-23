package main

import (
	"context"
	"file-editor/api"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(100*1024*1024), // 100 MB
			grpc.MaxCallSendMsgSize(100*1024*1024), // 100 MB
		),
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := api.NewTextEditorClient(conn)

	ctx := context.Background()

	data, err := c.ReadAllFiles(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("could read all files: %v", err)
	}

	fmt.Println(data)
}
