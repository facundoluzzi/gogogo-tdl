package commands

import (
	"context"
	"file-editor/api"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type SaveFileCommand struct {
	Name string
}

func (c *SaveFileCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error al obtener el directorio de trabajo actual: %w", err)
	}

	filePath := filepath.Join(cwd, c.Name)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error al leer el archivo %s: %w", c.Name, err)
	}

	request := api.SaveFileRequest{
		Filename: c.Name,
		Content:  content,
	}

	r, err := t.SaveFile(ctx, &request)
	if err != nil {
		return "", err
	}

	return r.Response, nil
}

func (c *SaveFileCommand) Print() {
	print("create command", c.Name)
}
