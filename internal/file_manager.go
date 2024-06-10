package internal

import (
	"os"
	"sync"
)

type Command struct {
	Name string
	Filename string
	ResponseChan chan interface{}
}

type FileManager struct {
	fileChans map[string]chan Command
	mutex sync.Mutex
}

func NewFileManager() *FileManager {
	return &FileManager{
		fileChans: make(map[string]chan Command),
	}
}

func (fm *FileManager) GetFileChan(filename string) chan Command {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()	
	fileChan , exists := fm.fileChans[filename]
	if !exists {
		fileChan = make(chan Command)
		fm.fileChans[filename] = fileChan
		go fm.HandleFileCommands(fileChan, filename)
	}
	return fileChan
}

func (fm *FileManager) HandleFileCommands(fileChan chan Command, filename string) {
	var file *os.File
	var handler CommandHandler

	for command := range fileChan {
		switch command.Name {
		case "read":
			handler = ReadCommandHandler{}
		// the new handlers should go here
		default:
			// handle error
		}
		handler.handle(command, file)
	}
}



