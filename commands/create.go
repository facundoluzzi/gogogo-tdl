package commands

import (
	"fmt"
	"os"
)

type CreateCommand struct {
	Name    string
	Content string
}

func (c *CreateCommand) Run() error {
	file, err := os.Create(c.Name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(c.Content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func (c *CreateCommand) Print() {
	print("create command", c.Name, c.Content)
}
