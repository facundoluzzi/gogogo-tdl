package commands

import (
	"context"
	"file-editor/api"
)

type CreateCommand struct {
	Name    string
	Content string
}

func (c *CreateCommand) Run(t api.TextEditorClient) (string, error) {
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

func (c *CreateCommand) Print() {
	print("create command", c.Name, c.Content)
}
