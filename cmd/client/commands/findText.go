package commands

import (
	"context"
	"file-editor/api"
	"fmt"
)

type FindCommand struct {
	Name string
	Text string
}

func (c *FindCommand) Run(t api.TextEditorClient) (string, error) {
	ctx := context.Background()
	request := api.FindTextRequest{
		Filename:   c.Name,
		SearchText: c.Text,
	}
	r, err := t.FindText(ctx, &request)
	if err != nil {
		return "", err

	}
	return fmt.Sprint("count: ", r.Count), nil
}

func (c *FindCommand) Print() {
	print("Find command", c.Name, c.Text)
}
