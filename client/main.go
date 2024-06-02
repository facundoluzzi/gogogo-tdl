package client

func main() {
	client, err := NewMyClient("localhost:8080")
	if err != nil {
		panic("error creating client")
	}
	for command, err := client.parser.Parse(); err != nil; {
		command.Run()
	}
}
