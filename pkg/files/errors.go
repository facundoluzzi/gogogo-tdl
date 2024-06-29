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

type NewFileAlreadyExistsError struct {
	*CustomError
}

type OutOfRangeError struct {
	*CustomError
}

func NewOutOfRangeError(message string) *OutOfRangeError {
	return &OutOfRangeError{
		CustomError: &CustomError{
			message: message,
		},
	}
}

func (f *OutOfRangeError) Is(target error) bool {
	_, ok := target.(*OutOfRangeError)
	return ok
}

func NewNewFileAlreadyExistsError(message string) *NewFileAlreadyExistsError {
	return &NewFileAlreadyExistsError{
		CustomError: &CustomError{
			message: message,
		},
	}
}

func (f *NewFileAlreadyExistsError) Is(target error) bool {
	_, ok := target.(*NewFileAlreadyExistsError)
	return ok
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
