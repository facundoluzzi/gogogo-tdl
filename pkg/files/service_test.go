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
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	request := &api.SaveFileRequest{
		Filename: fileName,
		Content:  content,
	}

	res, err := suite.service.SaveFile(context.Background(), request)
	suite.Require().NotNil(res)
	suite.Require().NoError(err)

	suite.Assert().Equal("Archivo 'testfile.txt' ha sido subido exitosamente", res.Response)

	fileContent, err := os.ReadFile(filePath)
	suite.Require().NotEmpty(fileContent)
	suite.Require().NoError(err)

	suite.Assert().Equal(content, fileContent)
}

func (suite *FileServiceSuite) TestSaveFileErrorContextEnded() {
	fileName := "testfile.txt"
	content := []byte("Hello, World!")

	request := &api.SaveFileRequest{
		Filename: fileName,
		Content:  content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := suite.service.SaveFile(ctx, request)
	suite.Require().Nil(res)
	suite.Require().ErrorIs(&files.ContextDoneError{}, err)
}

func (suite *FileServiceSuite) TestReadFileSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	assert.NoError(suite.T(), err)

	response, err := suite.service.ReadFile(context.Background(), fileName)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Hello, World!", response.Content)
}

func (suite *FileServiceSuite) TestReadFileErrorFileNotFound() {
	fileName := "testfile.txt"

	response, err := suite.service.ReadFile(context.Background(), fileName)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestFindTextSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, 1 Hello!\nHello, 2!\nHello, 3!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.FindTextRequest{
		Filename:   fileName,
		SearchText: "Hello",
	}

	response, err := suite.service.FindText(context.Background(), request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal(int64(4), response.Count)
	suite.Assert().Len(response.Lines, 3)
	suite.Assert().Contains(response.Lines, "Hello, 1 Hello!")
	suite.Assert().Contains(response.Lines, "Hello, 2!")
	suite.Assert().Contains(response.Lines, "Hello, 3!")
}

func (suite *FileServiceSuite) TestFindTextErrorFileNotFound() {
	request := &api.FindTextRequest{
		Filename:   "testfile.txt",
		SearchText: "Hello",
	}

	response, err := suite.service.FindText(context.Background(), request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}
