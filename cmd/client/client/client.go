package client

import (
	"bufio"
	"file-editor/cmd/client/input"
	"file-editor/proto"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TextEditor struct {
	parser *input.Parser
	conn   *grpc.ClientConn
	client proto.TextEditorClient
	reader *bufio.Reader
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

	c := proto.NewTextEditorClient(conn)

	parser := input.Parser{}

	reader := bufio.NewReader(os.Stdin)

	return &TextEditor{
		conn:   conn,
		parser: &parser,
		client: c,
		reader: reader,
	}, nil
}

// Close closes the gRPC connection
func (c *TextEditor) Close() error {
	return c.conn.Close()
}

func (c *TextEditor) Run() error {
	for {
		clearScreen()

		fmt.Println("==== Text Editor ====")
		fmt.Println("Enter command or write 'help' to see available commands")
		fmt.Println("=====================")

		input, err := c.reader.ReadString('\n')
		if err != nil {
			c.handleCommandError(err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		command, err := c.parser.Parse(input)
		if err != nil {
			c.handleCommandError(err)
			continue
		}

		fmt.Print("Running command...\n")

		response, err := command.Run(c.client)
		if err != nil {
			c.handleCommandError(err)
			continue
		}

		fmt.Println(response)

		fmt.Print("Press Enter to continue...\n\n\n")

		_, _ = c.reader.ReadString('\n')
	}

	fmt.Println("Closing client...")

	if err := c.conn.Close(); err != nil {
		log.Printf("error closing client: %v", err)
	}

	return nil
}

// func (c *TextEditor) ParseFromArgs() (input.Command, error) {
// 	return c.parser.ParseFromArgs()
// }

func (c *TextEditor) ExecuteCommand(command input.Command) {
	response, err := command.Run(c.client)
	if err != nil {
		fmt.Println("error running command:", err)
		return
	}

	fmt.Println(response)

	c.conn.Close()
}

func (c *TextEditor) handleCommandError(err error) {
	log.Printf("No se pudo ejecutar el comando solicitado, intente nuevamente: \nERROR: %v\n", err)
	fmt.Println("=====================")
	fmt.Print("Press Enter to continue...\n\n\n")

	_, _ = c.reader.ReadString('\n')
}

// clearScreen cleans the terminal
func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	_ = cmd.Run()
}
