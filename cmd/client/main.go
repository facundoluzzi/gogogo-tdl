package main

import (
	"context"
	"file-editor/api"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
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

	a, _ := os.Getwd()
	content, err := os.ReadFile(fmt.Sprintf(a + "/cmd/client/" + filename))
	if err != nil {
		log.Fatalf("could not open file %s: %v", filename, err)
	}

	ctx := context.Background()

	_, err = c.SaveFile(ctx, &api.SaveFileRequest{Filename: filename, Content: content})
	if err != nil {
		log.Fatalf("could not create file %s: %v", filename, err)
	}
}
