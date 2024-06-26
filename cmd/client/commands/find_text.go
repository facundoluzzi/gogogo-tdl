package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"strings"
	"time"
)

type FindCommand struct {
	Name string
	Text string
}

func (c *FindCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := api.FindTextRequest{
		Filename:   c.Name,
		SearchText: c.Text,
	}

	r, err := t.FindText(ctx, &request)
	if err != nil {
		return "", err

	}

	var response strings.Builder
	response.WriteString("==== Search Results ====\n")
	response.WriteString(fmt.Sprintf("File: %s\n", c.Name))
	response.WriteString(fmt.Sprintf("Search Text: %s\n", c.Text))
	response.WriteString(fmt.Sprintf("Total '%s' Word Count: %d\n", c.Text, r.Count))
	response.WriteString("----------------------\n")

	for _, line := range r.Lines {
		response.WriteString(fmt.Sprintf("%s\n", line))
	}

	response.WriteString("======================\n")

	return response.String(), nil
}

func (c *FindCommand) Print() {
	print("Find command", c.Name, c.Text)
}
