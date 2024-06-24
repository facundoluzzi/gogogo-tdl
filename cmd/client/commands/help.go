package commands

import (
	"file-editor/api"
)

type HelpCommand struct {
}

func (c *HelpCommand) Run(t api.TextEditorClient) (string, error) {
	response := "\nTo use the client the following commands are available:\n" + "read <filename>\n" + "readall\n" + "save <filename> <content>\n" + "find <filename> <text> \n" + "or write 'exit' to close the client" + "\nEnter command: "
	return response, nil
}

func (c *HelpCommand) Print() {
	print("help command")
}
