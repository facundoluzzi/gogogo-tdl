package commands

import (
	"context"
	"file-editor/proto"
	"fmt"
	"strings"
	"time"
)

type FindAndReplaceCommand struct {
	Name    string
	Find    string
	Replace string
}

func (c *FindAndReplaceCommand) Run(t proto.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	request := proto.FindAndReplaceRequest{
		Filename:    c.Name,
		FindText:    c.Find,
		ReplaceText: c.Replace,
	}

	r, err := t.FindAndReplace(ctx, &request)
	if err != nil {
		return "", err
	}

	var response strings.Builder
	response.WriteString("============================= Find and Replace Results ==============================\n")
	response.WriteString(fmt.Sprintf("File: %s\n", c.Name))
	response.WriteString(fmt.Sprintf("Found '%s' and replaced by '%s' %d times at positions: ", c.Find, c.Replace, r.Count))
	for _, pos := range r.Positions {
		response.WriteString(fmt.Sprintf("%d, ", pos))
	}
	response.WriteString("\n")
	response.WriteString("-------------------------------------------------------------------------------------\n")

	return response.String(), nil
}

func (c *FindAndReplaceCommand) Print() {
	print("replace command", c.Name, c.Find, c.Replace)
}
