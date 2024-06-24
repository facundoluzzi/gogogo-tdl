package client

import (
	"bufio"
	"file-editor/client/input"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

type FileEditorClient struct {
	Parser *input.Parser
	conn   *grpc.ClientConn
}

func NewMyClient(address string) (*FileEditorClient, error) {
	// Set up a connection to the server
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Create a new client from the connection
	//client := pb.NewYourServiceClient(conn)

	//Create Parser
	parser := input.Parser{}

	return &FileEditorClient{
		conn:   conn,
		Parser: &parser,
		//client: client,
	}, nil
}

// Close closes the gRPC connection
func (c *FileEditorClient) Close() error {
	return c.conn.Close()
}

// Example method to call a service method
/*func (c *FileEditorClient) DoSomething(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
	return c.client.YourMethod(ctx, req)
}*/

/*func main() {
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

	fmt.Println(file.Body)
}*/

func (c *FileEditorClient) Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	for input, err := reader.ReadString('\n'); input != "exit"; {
		if err != nil {
			return fmt.Errorf("error reading command: %w", err)
		}
		input = strings.TrimSpace(input)
		command, err := c.Parser.Parse(input)
		if err != nil {
			return fmt.Errorf("error parsing command: %w", err)
		}
		command.Print()
		fmt.Print("Enter command: ")
		input, err = reader.ReadString('\n')
	}
	print("closing client...")
	return nil
}
