package handlers

import (
	"context"
	"file-editor/api"
)

type FilesService interface {
	ReadFile(ctx context.Context, filename string) (*api.ReadFileResponse, error)
	SaveFile(ctx context.Context, request *api.SaveFileRequest) (*api.SaveFileResponse, error)
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

func (h *Handler) ReadFile(ctx context.Context, req *api.ReadFileRequest) (*api.ReadFileResponse, error) {
	return h.FilesService.ReadFile(ctx, req.Filename)
}

func (h *Handler) SaveFile(ctx context.Context, req *api.SaveFileRequest) (*api.SaveFileResponse, error) {
	return h.FilesService.SaveFile(ctx, req)
}
