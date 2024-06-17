package files

import (
	"context"
	"encoding/binary"
	"file-editor/api"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Service struct {
	Producer   chan []byte
	Consumer   chan []byte
	waitGroups map[string]*sync.WaitGroup
}

func New(ch chan []byte) *Service {
	service := &Service{
		Producer:   ch,
		Consumer:   ch,
		waitGroups: make(map[string]*sync.WaitGroup),
	}

	go service.RunConsumer()

	return service
}

func (s *Service) ReadFile(ctx context.Context, filename string) (*api.ReadFileResponse, error) {
	err := s.existsFile(filename)
	if err != nil {
		return nil, err
	}

	path, err := s.getFilePath(filename)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &api.ReadFileResponse{Content: string(content)}, nil
}

// SaveFile creates a new file using the content provided in the request
func (s *Service) SaveFile(ctx context.Context, request *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	fileBytes := request.Content
	filename := request.Filename

	// create a header indicating the filename length and the operation type (create)
	header := make([]byte, headerSize)
	header[lastMessageHeaderIndex] = byte(notLastMessage)
	binary.LittleEndian.PutUint32(header[filenameHeaderStartIndex:filenameHeaderEndIndex], uint32(len(filename)))

	path, err := s.getFilePath(filename)
	if err != nil {
		return nil, err
	}

	// create the new file, overwriting if it exists
	_, err = os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new file: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	s.waitGroups[filename] = wg

	for i := 0; i < len(fileBytes); i += bytesSize {
		end := i + bytesSize

		if end >= len(fileBytes) {
			header[lastMessageHeaderIndex] = byte(lastMessage)
			end = len(fileBytes)
		}

		bytesToSend := append(header, append([]byte(filename), fileBytes[i:end]...)...)

		s.Producer <- bytesToSend
	}

	wg.Wait()
	delete(s.waitGroups, filename)

	return &api.SaveFileResponse{Response: fmt.Sprintf("File '%s' has been uploaded successfully", filename)}, nil
}

func (s *Service) RunConsumer() {
	for bytes := range s.Consumer {
		if len(bytes) < headerSize {
			fmt.Println("error: invalid byte slice")
			continue
		}

		filenameLen := int(binary.LittleEndian.Uint32(bytes[filenameHeaderStartIndex:filenameHeaderEndIndex]))

		if len(bytes) < headerSize+filenameLen {
			fmt.Println("error: invalid byte slice")
			continue
		}

		filename := string(bytes[headerSize : headerSize+filenameLen])
		data := bytes[headerSize+filenameLen:]

		path, err := s.getFilePath(filename)
		if err != nil {
			fmt.Printf("error getting file path: %s\n", err.Error())
			continue
		}

		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("error opening file: %s\n", err.Error())
			continue
		}

		_, err = file.Write(data)
		if err != nil {
			fmt.Printf("error writing file: %s\n", err.Error())
			file.Close()
			continue
		}

		err = file.Close()
		if err != nil {
			fmt.Printf("error closing file: %s\n", err.Error())
			continue
		}

		if bytes[lastMessageHeaderIndex] == lastMessage {
			if wg, exists := s.waitGroups[filename]; exists {
				wg.Done()
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) getFilePath(filename string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}

	return filepath.Join(cwd, filesFolder, filename), nil
}

func (s *Service) existsFile(filename string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}

	directories, err := os.ReadDir(cwd + filesFolder)
	for _, dir := range directories {
		if dir.Name() == filename {
			return nil
		}
	}

	return NewFileNotFoundError("invalid request, file doesn't exists")
}
