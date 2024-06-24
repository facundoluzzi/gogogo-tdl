package commands

import (
	"context"
	"file-editor/api"
	"time"
)

type ReadCommand struct {
	Name string
}

func (c *ReadCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request := api.ReadFileRequest{
		Filename: c.Name,
	}
	r, err := t.ReadFile(ctx, &request)
	if err != nil {
		return "", err

	}
	return r.Content, nil
}

func (c *ReadCommand) Print() {
	print("read command", c.Name)
}
