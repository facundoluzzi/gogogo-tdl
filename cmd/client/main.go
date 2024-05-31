package main

import (
	"context"
	"file-editor/api"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := api.NewTextEditorClient(conn)

	if len(os.Args) < 2 {
		panic("filename is required")
	}

	filename := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, err := c.ReadFile(ctx, &api.ReadFileRequest{Filename: filename})
	if err != nil {
		log.Fatalf("could not open file %s: %v", filename, err)
	}

	fmt.Println(file.Content)
}
