package zerror

type Zerror interface {
	error
	GetCode() int
	GetMessage() string
}

type zerror struct {
	Code    int
	Message string
	err     error
}

func (e *zerror) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *zerror) GetCode() int {
	if e == nil {
		return 0
	}
	return e.Code
}

func (e *zerror) GetMessage() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *zerror) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func NewZError(code int, message string, err error) *zerror {
	return &zerror{
		Code:    code,
		Message: message,
		err:     err,
	}
}