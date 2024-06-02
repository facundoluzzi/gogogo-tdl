package client

import (
	"gogogo-tdl/client/input"
	"google.golang.org/grpc"
	"log"
)

type MyClient struct {
	parser *input.Parser
	conn   *grpc.ClientConn
}

func NewMyClient(address string) (*MyClient, error) {
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

	return &MyClient{
		conn:   conn,
		parser: &parser,
		//client: client,
	}, nil
}

// Close closes the gRPC connection
func (c *MyClient) Close() error {
	return c.conn.Close()
}

// Example method to call a service method
/*func (c *MyClient) DoSomething(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
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

	fmt.Println(file.Content)
}*/
