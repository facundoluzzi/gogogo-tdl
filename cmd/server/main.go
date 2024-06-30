package main

import (
	"file-editor/internal/handlers"
	filesService "file-editor/pkg/files"
	"file-editor/proto"
	"log"
	"net"

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

	ch := make(chan []byte)

	filesService := filesService.New(ch)

	handler := handlers.New(filesService)

	proto.RegisterTextEditorServer(s, handler)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
