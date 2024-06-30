package handlers_test

//go:generate mockgen -destination=./mocks/mock_files_handler.go -package handlers_mocks -source=files_handler.go

import (
	"context"
	"encoding/json"
	"errors"
	"file-editor/internal/handlers"
	handlers_mocks "file-editor/internal/handlers/mocks"
	"file-editor/pkg/files"
	"file-editor/proto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
	filesService *handlers_mocks.MockFilesService
	filesHandler *handlers.Handler
	ctrl         *gomock.Controller
}

func (suite *HandlerSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())

	suite.filesService = handlers_mocks.NewMockFilesService(suite.ctrl)

	suite.filesHandler = handlers.New(suite.filesService)
}

func (suite *HandlerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (suite *HandlerSuite) TestNewFileSuccess() {
	ctx := context.Background()
	req := &proto.NewFileRequest{
		Filename: "test.txt",
		Content:  "This is a test file content",
	}
	expectedResponse := &proto.NewFileResponse{
		Response: "File created successfully",
	}
	responseJSON, _ := json.Marshal(expectedResponse)

	suite.filesService.EXPECT().Request(files.NewFile, req).Return(string(responseJSON), nil)

	res, err := suite.filesHandler.NewFile(ctx, req)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	suite.Assert().Equal(expectedResponse.Response, res.Response)
}

func (suite *HandlerSuite) TestNewFileRequestError() {
	ctx := context.Background()
	req := &proto.NewFileRequest{
		Filename: "test.txt",
		Content:  "This is a test file content",
	}

	suite.filesService.EXPECT().Request(files.NewFile, req).Return(nil, errors.New("request error"))

	res, err := suite.filesHandler.NewFile(ctx, req)

	suite.Require().Error(err)
	suite.Require().Nil(res)

	suite.Assert().Equal("request error", err.Error())
}

func (suite *HandlerSuite) TestNewFileUnmarshalError() {
	ctx := context.Background()
	req := &proto.NewFileRequest{
		Filename: "test.txt",
		Content:  "This is a test file content",
	}

	invalidJSON := "{invalid json}"

	suite.filesService.EXPECT().Request(files.NewFile, req).Return(invalidJSON, nil)

	res, err := suite.filesHandler.NewFile(ctx, req)

	suite.Require().Error(err)
	suite.Require().Nil(res)

	suite.Assert().Contains(err.Error(), "invalid character")
}

func (suite *HandlerSuite) TestReadFileFileAlreadyExistsError() {
	ctx := context.Background()
	req := &proto.ReadFileRequest{
		Filename: "test.txt",
	}

	suite.filesService.EXPECT().Request(files.Read, req).Return(nil, files.NewNewFileAlreadyExistsError("file already exists in the directory"))

	res, err := suite.filesHandler.ReadFile(ctx, req)

	suite.Require().Error(err)
	suite.Require().Nil(res)

	suite.Assert().Contains(err.Error(), "bad request: file already exists in the directory")
	suite.Assert().ErrorIs(err, files.NewNewFileAlreadyExistsError(""))
}

func (suite *HandlerSuite) TestReadFileSuccess() {
	ctx := context.Background()
	req := &proto.ReadFileRequest{
		Filename: "test.txt",
	}
	expectedResponse := &proto.ReadFileResponse{
		Content: "This is the content of the file",
	}
	responseJSON, _ := json.Marshal(expectedResponse)

	suite.filesService.EXPECT().Request(files.Read, req).Return(string(responseJSON), nil)

	res, err := suite.filesHandler.ReadFile(ctx, req)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.Assert().Equal(expectedResponse.Content, res.Content)
}

func (suite *HandlerSuite) TestReadFileRequestError() {
	ctx := context.Background()
	req := &proto.ReadFileRequest{
		Filename: "test.txt",
	}

	suite.filesService.EXPECT().Request(files.Read, req).Return(nil, errors.New("request error"))

	res, err := suite.filesHandler.ReadFile(ctx, req)
	suite.Require().Error(err)
	suite.Require().Nil(res)

	suite.Assert().Equal("request error", err.Error())
}

func (suite *HandlerSuite) TestReadFileUnmarshalError() {
	ctx := context.Background()
	req := &proto.ReadFileRequest{
		Filename: "test.txt",
	}

	invalidJSON := "{invalid json}"

	suite.filesService.EXPECT().Request(files.Read, req).Return(invalidJSON, nil)

	res, err := suite.filesHandler.ReadFile(ctx, req)
	suite.Require().Error(err)
	suite.Require().Nil(res)

	suite.Assert().Contains(err.Error(), "invalid character")
}
