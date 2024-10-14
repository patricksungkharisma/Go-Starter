package error

type TemplateError struct {
	StatusCode int
	Error      error
}

func GenerateError(statusCode int, err error) *TemplateError {
	return &TemplateError{
		StatusCode: statusCode,
		Error:      err,
	}
}
