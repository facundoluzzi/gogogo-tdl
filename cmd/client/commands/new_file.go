package commands

import (
	"context"
	"file-editor/proto"
	"fmt"
	"strings"
	"time"
)

type NewFileCommand struct {
	Name    string
	Content string
}

func (c *NewFileCommand) Run(t proto.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := proto.NewFileRequest{
		Filename: c.Name,
		Content:  c.Content,
	}

	r, err := t.NewFile(ctx, &request)
	if err != nil {
		return "", err
	}
	var response strings.Builder
	response.WriteString("======== New File ========\n")
	response.WriteString(fmt.Sprintf(r.Response))
	if len(c.Content) > 0 {
		response.WriteString("Content:\n")
		response.WriteString(c.Content)
		response.WriteString("\n")
	}
	response.WriteString("\n========================\n")

	return response.String(), nil
}

func (c *NewFileCommand) Print() {
	print("new file command", c.Name, c.Content)
}
