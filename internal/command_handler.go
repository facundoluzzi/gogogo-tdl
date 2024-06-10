package internal

import (
	"os"
	"fmt"
)

type CommandHandler interface {
	handle(command Command,file *os.File)
}

type ReadCommandHandler struct {
}

func (handler ReadCommandHandler) handle(command Command,file *os.File) {
	if file != nil {
		content, err := os.ReadFile(command.Filename)
		if err != nil {
			command.ResponseChan <- err
			return
		} else {
			command.ResponseChan <- string(content)
		}
	} else {
		command.ResponseChan <- fmt.Errorf("file %s not found", command.Filename)
	}
}



