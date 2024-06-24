package commands

import (
	"context"
	"file-editor/api"
	"time"
)

type SaveFileCommand struct {
	Name    string
	Content string
}

func (c *SaveFileCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	request := api.SaveFileRequest{
		Filename: c.Name,
		Content:  []byte(c.Content),
	}
	r, err := t.SaveFile(ctx, &request)
	if err != nil {
		return "", err

	}
	return r.Response, nil
}

func (c *SaveFileCommand) Print() {
	print("create command", c.Name, c.Content)
}
