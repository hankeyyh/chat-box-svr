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
	End     bool   `json:"end"`
}

type AppReleaseRequest struct {
	AppId uint64 `json:"app_id"`
}

type AppUnReleaseRequest struct {
	AppId uint64 `json:"app_id"`
}

type AppUpsertRequest struct {
	AppId           uint64     `json:"app_id"`
	ModelId         uint64     `json:"model_id"`
	Name            string  `json:"name"`
	Temperature     float32 `json:"temperature"`//控制模型回复的创造性和随机性。数值越高，模型的回复越有创造性
	TopP            float32 `json:"top_p"`//控制模型回复时候选内容的多少。数值太低模型的回复会比较固定，数值太高模型的回复会比较放飞自我
	MaxOutputTokens int     `json:"max_output_tokens"`//控制模型一次回复的最大回复长度
	Context         int     `json:"context"` 			// 上下文数量
	CreatedBy       uint64  `json:"created_by"`
	Introduction    string  `json:"introduction"`
	Prologue        string  `json:"prologue"` 			// 开场白
	Prompt          string  `json:"prompt"`
	IsPublic        int8    `json:"is_public"`
}

type AppUpsertResponseData struct {
	AppId uint64 `json:"app_id"`
}

type AppChatRequest struct {
	AppId uint64    `json:"app_id"`
	Content string `json:"content"`
}