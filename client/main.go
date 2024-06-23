package main

import (
	"bufio"
	"errors"
	client2 "file-editor/client/client"
	"file-editor/client/input"
	"fmt"
	"os"
	"strings"
)

func main() {
	client, err := client2.NewMyClient("localhost:8080")
	if err != nil {
		panic("error creating client")
	}
	command, err := client.Parser.ParseFromArgs()
	if err == nil {
		command.Run()
	}
	if !errors.Is(err, input.ErrNoArgs) {
		print("error parsing arguments")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input := ""
	for input != "exit" {
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			return
		}
		input = strings.TrimSpace(input)
		command, err := client.Parser.Parse(input)
		if err != nil {
			fmt.Println("Error parsing command:", err)
			continue
		}
		print(command)
	}
	print("closing client...")
	return
}
