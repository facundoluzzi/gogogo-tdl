package commands

import (
	"context"
	"file-editor/api"
	"time"
)

type ReadAllCommand struct {
}

func (c *ReadAllCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := t.ReadAllFiles(ctx, &api.Empty{})
	if err != nil {
		return "", err

	}
	response := ""
	for _, file := range r.Content {
		name := file.Name
		content := file.Content

		response += "file: " + name + " content: " + content + "\n"
	}
	return response, nil
}

func (c *ReadAllCommand) Print() {
	print("read all command")
}
