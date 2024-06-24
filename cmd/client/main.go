package main

import (
	"errors"
	client2 "file-editor/cmd/client/client"
	"file-editor/cmd/client/input"
	"fmt"
)

const (
	address = "localhost:50051"
)

func main() {
	textEditor, err := client2.NewTextEditorClient(address)
	if err != nil {
		panic("error creating client")
	}
	command, err := textEditor.ParseFromArgs()
	if err == nil {
		textEditor.ExecuteCommand(command)
		return
	}
	if !errors.Is(err, input.ErrNoArgs) {
		print("error parsing arguments")
		return
	}
	err = textEditor.Run()
	if err != nil {
		fmt.Print("error running client: ", err)
	}
	return
}
