package files

import (
	"context"
	"file-editor/api"
	"fmt"
	"os"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (m *Service) ReadFile(ctx context.Context, filename string) (*api.ReadFileResponse, error) {
	a, _ := os.Getwd()
	content, err := os.ReadFile(fmt.Sprintf(a + "/pkg/files/" + filename))
	if err != nil {
		return nil, err
	}

	return &api.ReadFileResponse{Content: string(content)}, nil
}
