package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"time"
)

type DeleteFileCommand struct {
	Name string
}

func (c *DeleteFileCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := api.DeleteFileRequest{
		Filename: c.Name,
	}

	_, err := t.DeleteFile(ctx, &request)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("File %s deleted", c.Name), nil
}

func (c *DeleteFileCommand) Print() {
	print("delete file command", c.Name)
}
