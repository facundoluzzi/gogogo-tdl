package main

import (
	"context"
	"file-editor/api"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := api.NewTextEditorClient(conn)

	if len(os.Args) < 2 {
		panic("filename is required")
	}

	filename := os.Args[1]

	ctx := context.Background()

	data, err := c.FindText(ctx, &api.FindTextRequest{Filename: filename, SearchText: os.Args[2]})
	if err != nil {
		log.Fatalf("could not create file %s: %v", filename, err)
	}

	fmt.Println(data)
}
