package files

type CustomError struct {
	message string
}

func (p *CustomError) Error() string {
	return p.message
}

type FileNotFoundError struct {
	*CustomError
}

func NewFileNotFoundError(message string) *FileNotFoundError {
	return &FileNotFoundError{
		CustomError: &CustomError{
			message: message,
		},
	}
}

func (f *FileNotFoundError) Is(target error) bool {
	_, ok := target.(*FileNotFoundError)
	return ok
}
