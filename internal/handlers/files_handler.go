package handlers

import (
	"context"
	"file-editor/api"
)

type FilesService interface {
	ReadFile(ctx context.Context, filename string) (*api.ReadFileResponse, error)
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
