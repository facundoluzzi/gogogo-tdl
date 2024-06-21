package files_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"file-editor/api"
	"file-editor/pkg/files"
)

type FileServiceSuite struct {
	suite.Suite
	service            *files.Service
	temporaryDirectory string
}

func (suite *FileServiceSuite) SetupTest() {
	ch := make(chan []byte)
	suite.service = files.New(ch)

	temporaryDirectory := filepath.Join(".", "filesstorage")
	err := os.Mkdir(temporaryDirectory, 0755)
	suite.temporaryDirectory = temporaryDirectory
	suite.Require().NoError(err)
}

func TestFileServiceSuite(t *testing.T) {
	suite.Run(t, new(FileServiceSuite))
}

func (suite *FileServiceSuite) TearDownTest() {
	close(suite.service.Producer)
	os.RemoveAll(suite.temporaryDirectory)
}

func (suite *FileServiceSuite) TestSaveFileSuccess() {
	filename := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, filename)
	content := []byte("Hello, World!")

	request := &api.SaveFileRequest{
		Filename: filename,
		Content:  content,
	}

	res, err := suite.service.SaveFile(context.Background(), request)
	suite.Require().NotNil(res)
	suite.Require().NoError(err)

	suite.Assert().Equal("File 'testfile.txt' has been uploaded successfully", res.Response)

	fileContent, err := os.ReadFile(filePath)
	suite.Require().NotEmpty(fileContent)
	suite.Require().NoError(err)

	suite.Assert().Equal(content, fileContent)
}

func (suite *FileServiceSuite) TestSaveFileErrorContextEnded() {
	filename := "testfile.txt"
	content := []byte("Hello, World!")

	request := &api.SaveFileRequest{
		Filename: filename,
		Content:  content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := suite.service.SaveFile(ctx, request)
	suite.Require().Nil(res)
	suite.Require().ErrorIs(&files.ContextDoneError{}, err)
}

func (suite *FileServiceSuite) TestReadFileSuccess() {
	filename := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, filename)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	assert.NoError(suite.T(), err)

	response, err := suite.service.ReadFile(context.Background(), filename)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Hello, World!", response.Content)
}

func (suite *FileServiceSuite) TestReadFileErrorFileNotFound() {
	filename := "testfile.txt"

	response, err := suite.service.ReadFile(context.Background(), filename)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}
