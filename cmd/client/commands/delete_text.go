package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"time"
)

type DeleteTextCommand struct {
	Name   string
	Start  int
	Length int
}

func (c *DeleteTextCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := api.DeleteTextRequest{
		Filename:      c.Name,
		StartPosition: int32(c.Start),
		Length:        int32(c.Length),
	}

	_, err := t.DeleteText(ctx, &request)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Text deleted from %s starting at %d for %d characters", c.Name, c.Start, c.Length), nil
}

func (c *DeleteTextCommand) Print() {
	print("Delete command", c.Name, c.Start, c.Length)
}
