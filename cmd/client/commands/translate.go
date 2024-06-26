package commands

import (
	"context"
	"file-editor/proto"
	"time"
)

type TranslateCommand struct {
	Name string
}

func (c *TranslateCommand) Run(t proto.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	request := proto.TranslateFileRequest{
		Filename: c.Name,
	}
	r, err := t.TranslateText(ctx, &request)
	if err != nil {
		return "", err

	}
	return r.Content, nil
}

func (c *TranslateCommand) Print() {
	print("translate command", c.Name)
}
