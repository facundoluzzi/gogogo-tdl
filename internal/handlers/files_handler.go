package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"file-editor/api"
)

type FilesService interface {
	ReadAllFiles(ctx context.Context, request *api.Empty) (*api.ReadAllFilesResponse, error)
	ReadFile(filename string) (*api.ReadFileResponse, error)
	SaveFile(ctx context.Context, request *api.SaveFileRequest) (*api.SaveFileResponse, error)
	FindText(ctx context.Context, request *api.FindTextRequest) (*api.FindTextResponse, error)
	DoRequest(command string, filename string) (response interface{}, err error)
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

func (h *Handler) SaveFile(ctx context.Context, req *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	return h.FilesService.SaveFile(ctx, req)
}

func (h *Handler) FindText(ctx context.Context, req *api.FindTextRequest) (*api.FindTextResponse, error) {
	return h.FilesService.FindText(ctx, req)
}

func (h *Handler) ReadAllFiles(ctx context.Context, req *api.Empty) (*api.ReadAllFilesResponse, error) {
	return h.FilesService.ReadAllFiles(ctx, req)
}

func (h *Handler) ReadFile(ctx context.Context, req *api.ReadFileRequest) (*api.ReadFileResponse, error) {
	res := &api.ReadFileResponse{}
	response, err := h.FilesService.DoRequest("read", req.Filename)
	if err != nil {
		return nil, err
	}

	b, ok := response.(string)
	if !ok {
		return nil, errors.New("response type assertion to string failed")
	}

	err = json.Unmarshal([]byte(b), res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
