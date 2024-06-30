package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"file-editor/pkg/files"
	"file-editor/proto"
	"fmt"
)

const (
	apiURL = "https://7fbf693ca6eb4dccbc232dc858de3b94.proto.mockbin.io/"
)

type FilesService interface {
	Request(operationType files.OperationType, request interface{}) (response interface{}, err error)
}

type Handler struct {
	FilesService FilesService
	proto.UnimplementedTextEditorServer
}

func New(filesService FilesService) *Handler {
	return &Handler{
		FilesService: filesService,
	}
}

func (h *Handler) NewFile(ctx context.Context, req *proto.NewFileRequest) (*proto.NewFileResponse, error) {
	response, err := h.FilesService.Request(files.NewFile, req)
	if err != nil {
		if errors.Is(err, &files.NewFileAlreadyExistsError{}) {
			return nil, fmt.Errorf("bad request")
		}

		return nil, err
	}

	res := &proto.NewFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) SaveFile(ctx context.Context, req *proto.SaveFileRequest) (*proto.SaveFileResponse, error) {
	response, err := h.FilesService.Request(files.Save, req)
	if err != nil {
		return nil, err
	}

	res := &proto.SaveFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) FindText(ctx context.Context, req *proto.FindTextRequest) (*proto.FindTextResponse, error) {
	response, err := h.FilesService.Request(files.Find, req)
	if err != nil {
		return nil, err
	}

	res := &proto.FindTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ReadAllFiles(ctx context.Context, req *proto.Empty) (*proto.ReadAllFilesResponse, error) {
	response, err := h.FilesService.Request(files.ReadAll, req)
	if err != nil {
		return nil, err
	}

	res := &proto.ReadAllFilesResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ReadFile(ctx context.Context, req *proto.ReadFileRequest) (*proto.ReadFileResponse, error) {
	response, err := h.FilesService.Request(files.Read, req)
	if err != nil {
		var fileExistsErr *files.NewFileAlreadyExistsError
		if errors.As(err, &fileExistsErr) {
			return nil, fmt.Errorf("bad request: %w", fileExistsErr)
		}

		return nil, err
	}

	res := &proto.ReadFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) DeleteText(ctx context.Context, req *proto.DeleteTextRequest) (*proto.DeleteTextResponse, error) {
	response, err := h.FilesService.Request(files.Delete, req)
	if err != nil {
		return nil, err
	}

	res := &proto.DeleteTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) FindAndReplace(ctx context.Context, req *proto.FindAndReplaceRequest) (*proto.FindAndReplaceResponse, error) {
	response, err := h.FilesService.Request(files.FindAndReplace, req)
	if err != nil {
		return nil, err
	}

	res := &proto.FindAndReplaceResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) AppendText(ctx context.Context, req *proto.AppendTextRequest) (*proto.AppendTextResponse, error) {
	response, err := h.FilesService.Request(files.Append, req)
	if err != nil {
		return nil, err
	}

	res := &proto.AppendTextResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest) (*proto.DeleteFileResponse, error) {
	response, err := h.FilesService.Request(files.DeleteFile, req)
	if err != nil {
		return nil, err
	}

	res := &proto.DeleteFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) TranslateText(ctx context.Context, req *proto.TranslateFileRequest) (*proto.TranslateFileResponse, error) {
	response, err := h.FilesService.Request(files.Translate, req)
	if err != nil {
		return nil, err
	}

	res := &proto.TranslateFileResponse{}
	if err := json.Unmarshal([]byte(response.(string)), res); err != nil {
		return nil, err
	}

	return res, nil
}
