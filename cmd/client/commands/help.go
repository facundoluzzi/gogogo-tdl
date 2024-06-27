package commands

import (
	"file-editor/api"
)

type HelpCommand struct {
}

func (c *HelpCommand) Run(t api.TextEditorClient) (string, error) {
	response := `
		To use the client the following commands are available:
		---------------------------------------------------------------------------------------
		read <filename>           					- Read a file
		read-all                  					- Read all files
		edit <filename>           					- Edit a file
		save <filename>           					- Save a file
		find <filename> <text>    					- Find text in a file
		delete <filename> <start> <length>        	- Delete text length characters from start
		exit                      					- Close the client
		---------------------------------------------------------------------------------------
		Enter command: 
	`

	return response, nil
}

func (c *HelpCommand) Print() {
	print("help command")
}
