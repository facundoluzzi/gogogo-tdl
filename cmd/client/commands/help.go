package commands

import (
	"file-editor/api"
)

type HelpCommand struct {
}

func (c *HelpCommand) Run(t api.TextEditorClient) (string, error) {
	response := `
		To use the client the following commands are available:
		-------------------------------------------------------
		read <filename>           - Read a file
		read-all                  - Read all files
		save <filename>           - Save a file
		find <filename> <text>    - Find text in a file
		exit                      - Close the client
		-------------------------------------------------------
		Enter command: 
	`

	return response, nil
}

func (c *HelpCommand) Print() {
	print("help command")
}
