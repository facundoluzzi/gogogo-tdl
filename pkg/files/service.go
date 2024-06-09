package files

import (
	"context"
	"encoding/binary"
	"file-editor/api"
	"fmt"
	"os"
	"time"
)

type Service struct {
	Producer chan []byte
	Consumer chan []byte
}

func New(ch chan []byte) *Service {
	service := &Service{
		Producer: ch,
		Consumer: ch,
	}

	go service.RunConsumer()

	return service
}

func (s *Service) ReadFile(ctx context.Context, filename string) (*api.ReadFileResponse, error) {
	a, _ := os.Getwd()
	content, err := os.ReadFile(fmt.Sprintf(a + "/pkg/files/" + filename))
	if err != nil {
		return nil, err
	}

	return &api.ReadFileResponse{Content: string(content)}, nil
}

// SaveFile creates a new file using the content provided in the request
func (s *Service) SaveFile(ctx context.Context, request *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	// TODO lock file by name
	fileBytes := request.Content
	filename := request.Filename

	// create a header indicating the filename length and the operation type (create)
	header := make([]byte, headerSize)
	header[lastMessageHeaderIndex] = byte(notLastMessage)
	binary.LittleEndian.PutUint32(header[filenameHeaderStartIndex:filenameHeaderEndIndex], uint32(len(filename)))

	// create the new file, overwriting if it exists
	_, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new file: %w", err)
	}

	for i := 0; i < len(fileBytes); i += bytesSize {
		end := i + bytesSize

		if end >= len(fileBytes) {
			header[lastMessageHeaderIndex] = byte(lastMessage)
			end = len(fileBytes)
		}

		bytesToSend := append(header, append([]byte(filename), fileBytes[i:end]...)...)

		s.Producer <- bytesToSend
	}

	return &api.SaveFileResponse{Response: "success"}, nil
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

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
			// TODO release lock
		}

		time.Sleep(time.Second)
	}
}
