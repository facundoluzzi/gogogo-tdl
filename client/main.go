package main

import (
	"errors"
	client2 "file-editor/client/client"
	"file-editor/client/input"
	"fmt"
)

func main() {
	textEditor, err := client2.NewTextEditorClient("localhost:8080")
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
