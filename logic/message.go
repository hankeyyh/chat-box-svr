package logic

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

type ChatSpan struct {
	Content string `json:"content"`
	End bool `json:"end"`
}
