package handlers

import (
	"context"
	"encoding/json"
	"file-editor/api"
	"file-editor/pkg/files"
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
