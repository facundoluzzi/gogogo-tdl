package files_test

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

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

	res, err := suite.service.SaveFile(request)
	suite.Require().NotNil(res)
	suite.Require().NoError(err)

	suite.Assert().Equal("Archivo 'testfile.txt' ha sido subido exitosamente", res.Response)

	fileContent, err := os.ReadFile(filePath)
	suite.Require().NotEmpty(fileContent)
	suite.Require().NoError(err)

	suite.Assert().Equal(content, fileContent)
}

func (suite *FileServiceSuite) TestReadFileSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	assert.NoError(suite.T(), err)

	response, err := suite.service.ReadFile(&api.ReadFileRequest{Filename: fileName})
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Hello, World!", response.Content)
}

func (suite *FileServiceSuite) TestReadFileErrorFileNotFound() {
	fileName := "testfile.txt"

	response, err := suite.service.ReadFile(&api.ReadFileRequest{Filename: fileName})
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

	response, err := suite.service.FindText(request)
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

	response, err := suite.service.FindText(request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestReadAllFilesSuccess() {
	testFiles := []files.FileContent{
		{"file1.txt", "Hello, file 1!"},
		{"file2.txt", "Hello, file 2!"},
		{"file3.txt", "Hello, file 3!"},
	}

	for _, file := range testFiles {
		filePath := filepath.Join(suite.temporaryDirectory, file.Name)
		err := os.WriteFile(filePath, []byte(file.Content), 0644)
		suite.Require().NoError(err)
	}

	response, err := suite.service.ReadAllFiles()
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Len(response.Content, len(testFiles))

	sort.Slice(response.Content, func(i, j int) bool {
		return response.Content[i].Name < response.Content[j].Name
	})

	for i, fileContent := range response.Content {
		suite.Assert().Equal(testFiles[i].Name, fileContent.Name)
		suite.Assert().Equal(testFiles[i].Content, fileContent.Content)
	}
}

func (suite *FileServiceSuite) TestReadAllFilesSuccessNoFilesFoundReturnEmptyResponse() {
	response, err := suite.service.ReadAllFiles()
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Len(response.Content, 0)
}

func (suite *FileServiceSuite) TestDeleteTextSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.DeleteTextRequest{
		Filename:      fileName,
		StartPosition: 0,
		Length:        7,
	}

	response, err := suite.service.DeleteText(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
	suite.Assert().Equal("texto eliminado exitosamente", response.Message)
	fileContent, err := os.ReadFile(filePath)
	suite.Require().NoError(err)
	suite.Assert().Equal("World!", string(fileContent))
}

func (suite *FileServiceSuite) TestDeleteTextErrorFileNotFound() {
	request := &api.DeleteTextRequest{
		Filename:      "testfile.txt",
		StartPosition: 0,
		Length:        7,
	}

	response, err := suite.service.DeleteText(request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestDeleteTextErrorLengthOutOfRange() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.DeleteTextRequest{
		Filename:      fileName,
		StartPosition: 7,
		Length:        15,
	}

	response, err := suite.service.DeleteText(request)

	suite.Require().ErrorIs(&files.OutOfRangeError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestFindAndReplaceSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.FindAndReplaceRequest{
		Filename:    fileName,
		FindText:    "World",
		ReplaceText: "Universe",
	}

	response, err := suite.service.FindAndReplace(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
}

func (suite *FileServiceSuite) TestFindAndReplaceErrorFileNotFound() {
	request := &api.FindAndReplaceRequest{
		Filename:    "testfile.txt",
		FindText:    "World",
		ReplaceText: "Universe",
	}

	response, err := suite.service.FindAndReplace(request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestFindAndReplaceReturnsZeroCount() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.FindAndReplaceRequest{
		Filename:    fileName,
		FindText:    "Universe",
		ReplaceText: "World",
	}

	response, err := suite.service.FindAndReplace(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
	suite.Assert().Equal(int64(0), response.Count)
	suite.Assert().Len(response.Positions, 0)
}

func (suite *FileServiceSuite) TestFindAndReplaceSuccessMultipleReplacements() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World! Hello, World! Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.FindAndReplaceRequest{
		Filename:    fileName,
		FindText:    "World",
		ReplaceText: "Universe",
	}

	response, err := suite.service.FindAndReplace(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
	suite.Assert().Equal(int64(3), response.Count)
	suite.Assert().Len(response.Positions, 3)
}

func (suite *FileServiceSuite) TestNewFileSuccess() {
	fileName := "testfile.txt"
	content := "Hello, World!"

	request := &api.NewFileRequest{
		Filename: fileName,
		Content:  content,
	}

	response, err := suite.service.NewFile(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Archivo 'testfile.txt' ha sido creado exitosamente", response.Response)
}

func (suite *FileServiceSuite) TestNewFileErrorFileAlreadyExists() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.NewFileRequest{
		Filename: fileName,
		Content:  "Hello, World!",
	}

	response, err := suite.service.NewFile(request)
	suite.Require().ErrorIs(&files.NewFileAlreadyExistsError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestNewFileWithoutContentSucces() {

	fileName := "testfile.txt"
	request := &api.NewFileRequest{
		Filename: fileName,
		Content:  "",
	}

	response, err := suite.service.NewFile(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Archivo 'testfile.txt' ha sido creado exitosamente", response.Response)
}

func (suite *FileServiceSuite) TestAppendTextSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.AppendTextRequest{
		Filename: fileName,
		Content:  "Hello, Universe!",
	}

	response, err := suite.service.AppendText(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Texto 'Hello, Universe!' ha sido agregado exitosamente a 'testfile.txt'", response.Message)
}

func (suite *FileServiceSuite) TestAppendTextErrorFileNotFound() {
	request := &api.AppendTextRequest{
		Filename: "testfile.txt",
		Content:  "Hello, Universe!",
	}

	response, err := suite.service.AppendText(request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}

func (suite *FileServiceSuite) TestAppendTextWithoutContentSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.AppendTextRequest{
		Filename: fileName,
		Content:  "",
	}

	response, err := suite.service.AppendText(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Texto '' ha sido agregado exitosamente a 'testfile.txt'", response.Message)
}

func (suite *FileServiceSuite) TestDeleteFileSuccess() {
	fileName := "testfile.txt"
	filePath := filepath.Join(suite.temporaryDirectory, fileName)
	content := []byte("Hello, World!")

	err := os.WriteFile(filePath, content, 0644)
	suite.Require().NoError(err)

	request := &api.DeleteFileRequest{
		Filename: fileName,
	}

	response, err := suite.service.DeleteFile(request)
	suite.Require().NoError(err)
	suite.Require().NotNil(response)

	suite.Assert().Equal("Archivo 'testfile.txt' ha sido eliminado exitosamente", response.Message)
}

func (suite *FileServiceSuite) TestDeleteFileErrorFileNotFound() {
	request := &api.DeleteFileRequest{
		Filename: "testfile.txt",
	}

	response, err := suite.service.DeleteFile(request)
	suite.Require().ErrorIs(&files.FileNotFoundError{}, err)
	suite.Require().Nil(response)
}
