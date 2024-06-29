package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"file-editor/api"
	"file-editor/pkg/files"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	apiURL             = "https://7fbf693ca6eb4dccbc232dc858de3b94.api.mockbin.io/"
)

type FilesService interface {
	Request(operationType files.OperationType, request interface{}) (response interface{}, err error)
}

type Handler struct {
	FilesService FilesService
	api.UnimplementedTextEditorServer
}

func New(filesService FilesService) *Handler {
	return &Handler{
		FilesService: filesService,
	}
}

func (h *Handler) NewFile(ctx context.Context, req *api.NewFileRequest) (*api.NewFileResponse, error) {
	response, err := h.FilesService.Request(files.NewFile, req)
	if err != nil {
		return nil, err
	}

	res := &api.NewFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) SaveFile(ctx context.Context, req *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	response, err := h.FilesService.Request(files.Save, req)
	if err != nil {
		return nil, err
	}

	res := &api.SaveFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) FindText(ctx context.Context, req *api.FindTextRequest) (*api.FindTextResponse, error) {
	response, err := h.FilesService.Request(files.Find, req)
	if err != nil {
		return nil, err
	}

	res := &api.FindTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ReadAllFiles(ctx context.Context, req *api.Empty) (*api.ReadAllFilesResponse, error) {
	response, err := h.FilesService.Request(files.ReadAll, req)
	if err != nil {
		return nil, err
	}

	res := &api.ReadAllFilesResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ReadFile(ctx context.Context, req *api.ReadFileRequest) (*api.ReadFileResponse, error) {
	response, err := h.FilesService.Request(files.Read, req)
	if err != nil {
		return nil, err
	}

	res := &api.ReadFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) DeleteText(ctx context.Context, req *api.DeleteTextRequest) (*api.DeleteTextResponse, error) {
	response, err := h.FilesService.Request(files.Delete, req)
	if err != nil {
		return nil, err
	}

	res := &api.DeleteTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) FindAndReplace(ctx context.Context, req *api.FindAndReplaceRequest) (*api.FindAndReplaceResponse, error) {
	response, err := h.FilesService.Request(files.FindAndReplace, req)
	if err != nil {
		return nil, err
	}

	res := &api.FindAndReplaceResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) AppendText(ctx context.Context, req *api.AppendTextRequest) (*api.AppendTextResponse, error) {
	response, err := h.FilesService.Request(files.Append, req)
	if err != nil {
		return nil, err
	}

	res := &api.AppendTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) DeleteFile(ctx context.Context, req *api.DeleteFileRequest) (*api.DeleteFileResponse, error) {
	response, err := h.FilesService.Request(files.DeleteFile, req)
	if err != nil {
		return nil, err
	}

	res := &api.DeleteFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) TranslateFile(ctx context.Context, req *api.ReadFileRequest) (*api.ReadFileResponse, error) {
	response, err := h.FilesService.Request(files.Read, req)
	if err != nil {
		return nil, err
	}

	text := &api.ReadFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), text); err != nil {
		return nil, err
	}

	payload := []map[string]string{
		{"Text": text.Content},
	}
	// Translate text
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling payload: %v\n", err)
		os.Exit(1)
	}

	// Create the HTTP request
	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}
	text.Content = string(respBody)

	return text, nil
}
