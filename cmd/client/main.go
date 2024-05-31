package main

import (
	"context"
	"errors"
	"file-editor/api"
	"file-editor/commands"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Command interface {
	Run() error
}

type CommandLineArgs struct {
	Command string
	Name    string
	Content string
}

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

func parseArguments() (*CommandLineArgs, error) {
	var args CommandLineArgs

	flag.StringVar(&args.Command, "c", "", "Command")
	flag.StringVar(&args.Name, "n", "", "Name")
	flag.StringVar(&args.Content, "c", "", "Content")

	flag.Parse()

	if len(args.Command) == 0 || len(args.Name) == 0 || len(args.Content) == 0 {
		fmt.Println("input must be -c command -n name -c content")
		return nil, fmt.Errorf("invalid input")
	}

	fmt.Printf("Received Command: %s\n", args.Command)
	fmt.Printf("Received Name: %s\n", args.Name)
	fmt.Printf("Received Content: %s\n", args.Content)

	return &args, nil
}

func getCommandFromArgs(args *CommandLineArgs) (Command, error) {
	switch args.Command {
	case "create":
		return &commands.CreateCommand{
			Name:    args.Name,
			Content: args.Content,
		}, nil
	case "translate":
		return &commands.TranslateCommand{
			Name: args.Name,
		}, nil
	}
	return nil, errors.New("invalid command")
}
