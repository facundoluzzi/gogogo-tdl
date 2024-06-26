package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"strings"
	"time"
)

type ReadCommand struct {
	Name string
}

func (c *ReadCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	request := api.ReadFileRequest{
		Filename: c.Name,
	}
	r, err := t.ReadFile(ctx, &request)
	if err != nil {
		return "", err

	}

	var response strings.Builder
	response.WriteString("==== File Content ====\n")
	response.WriteString(fmt.Sprintf("File: %s\n", c.Name))
	response.WriteString("Content:\n")
	response.WriteString(r.Content)
	response.WriteString("\n======================\n")

	return response.String(), nil
}

func (c *ReadCommand) Print() {
	print("read command", c.Name)
}
