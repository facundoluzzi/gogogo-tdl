package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"strings"
	"time"
)

type ReadAllCommand struct {
}

func (c *ReadAllCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	r, err := t.ReadAllFiles(ctx, &api.Empty{})
	if err != nil {
		return "", err

	}

	var response strings.Builder
	response.WriteString("==== Files Content ====\n\n")

	for _, file := range r.Content {
		response.WriteString(fmt.Sprintf("File: %s\n", file.Name))
		response.WriteString(fmt.Sprintf("Content:\n%s\n", file.Content))
		response.WriteString("----------------------\n")
	}

	response.WriteString("======================\n")
	return response.String(), nil
}

func (c *ReadAllCommand) Print() {
	print("read all command")
}
