package main

import (
	"file-editor/api"
	"file-editor/internal/handlers"
	filesService "file-editor/pkg/files"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	ch := make(chan []byte)

	filesService := filesService.New(ch)

	handler := handlers.New(filesService)

	api.RegisterTextEditorServer(s, handler)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
