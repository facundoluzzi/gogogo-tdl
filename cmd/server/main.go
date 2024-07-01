package main

import (
	"file-editor/internal/handlers"
	filesService "file-editor/pkg/files"
	"file-editor/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(100 * 1024 * 1024), // 100 MB
		grpc.MaxSendMsgSize(100 * 1024 * 1024), // 100 MB
	}
	s := grpc.NewServer(opts...)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}
	directory := fmt.Sprintf("%s/%s", cwd, filesService.FilesFolder)
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create storage directory: %v", err)
	}

	ch := make(chan []byte)

	filesService := filesService.New(ch)

	handler := handlers.New(filesService)

	proto.RegisterTextEditorServer(s, handler)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
