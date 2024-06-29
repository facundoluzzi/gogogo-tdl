package commands

import (
	"context"
	"file-editor/api"
	"time"
)

type AppendTextCommand struct {
	Name    string
	Content string
}

func (c *AppendTextCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := api.AppendTextRequest{
		Filename: c.Name,
		Content:  c.Content,
	}

	r, err := t.AppendText(ctx, &request)
	if err != nil {
		return "", err
	}

	return r.Message, nil
}

func (c *AppendTextCommand) Print() {
	print("append command", c.Name, c.Content)
}
