package main

import (
	"errors"
	client2 "file-editor/client/client"
	"file-editor/client/input"
	"fmt"
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
	err = client.Run()
	if err != nil {
		fmt.Print("error running client: ", err)
	}
	return
}
