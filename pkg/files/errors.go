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

type ContextDoneError struct {
	*CustomError
}

func NewContextDoneError(message string) *ContextDoneError {
	return &ContextDoneError{
		CustomError: &CustomError{
			message: message,
		},
	}
}

func (c *ContextDoneError) Is(target error) bool {
	_, ok := target.(*ContextDoneError)
	return ok
}
