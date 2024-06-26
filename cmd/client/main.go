package main

import (
	"file-editor/cmd/client/client"
	"fmt"
)

const (
	address = "localhost:50051"
)

func main() {
	textEditor, err := client.NewTextEditorClient(address)
	if err != nil {
		panic("Error al crear el cliente: no se pudo establecer conexión")
	}

	err = textEditor.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el cliente:", err)
	}
}
