package client

import (
	"bufio"
	"file-editor/api"
	"file-editor/cmd/client/input"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
)

type TextEditor struct {
	parser *input.Parser
	conn   *grpc.ClientConn
	client api.TextEditorClient
}

func NewTextEditorClient(address string) (*TextEditor, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(100*1024*1024), // 100 MB
			grpc.MaxCallSendMsgSize(100*1024*1024), // 100 MB
		),
	}
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := api.NewTextEditorClient(conn)
	parser := input.Parser{}
	return &TextEditor{
		conn:   conn,
		parser: &parser,
		client: c,
	}, nil
}

// Close closes the gRPC connection
func (c *TextEditor) Close() error {
	return c.conn.Close()
}

func (c *TextEditor) Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	for input, err := reader.ReadString('\n'); input != "exit"; {
		if err != nil {
			return fmt.Errorf("error reading command: %w", err)
		}
		input = strings.TrimSpace(input)
		command, err := c.parser.Parse(input)
		if err != nil {
			return fmt.Errorf("error parsing command: %w", err)
		}
		response, err := command.Run(c.client)
		if err != nil {
			return fmt.Errorf("error running command: %w", err)

		}
		fmt.Println(response)
		input, err = reader.ReadString('\n')
	}
	print("closing client...")
	c.conn.Close()
	return nil
}

func (c *TextEditor) ParseFromArgs() (input.Command, error) {
	return c.parser.ParseFromArgs()
}

func (c *TextEditor) ExecuteCommand(command input.Command) {
	response, err := command.Run(c.client)
	if err != nil {
		fmt.Println("error running command:", err)
		return
	}
	fmt.Println(response)
	c.conn.Close()
}
