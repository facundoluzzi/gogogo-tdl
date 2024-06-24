package commands

import (
	"context"
	"file-editor/api"
)

type SaveFileCommand struct {
	Name    string
	Content string
}

func (c *SaveFileCommand) Run(t api.TextEditorClient) (string, error) {
	ctx := context.Background()
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
