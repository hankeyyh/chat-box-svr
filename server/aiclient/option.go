package aiclient

type Option struct {
	APIKey  string
	BaseURL string
	ThinkingBudget int32
	Model string
	Temperature float32
}

type OptionFunc func(o *Option)

func WithAPIKey(apiKey string) OptionFunc {
	return func(o *Option) {
		o.APIKey = apiKey
	}
}

func WithBaseURL(baseURL string) OptionFunc {
	return func(o *Option) {
		o.BaseURL = baseURL
	}
}

func WithThinkingBudget(thinkingBudget int32) OptionFunc {
	return func(o *Option) {
		o.ThinkingBudget = thinkingBudget
	}
}

func WithModel(model string) OptionFunc {
	return func(o *Option) {
		o.Model = model
	}
}

func WithTemperature(temperature float32) OptionFunc {
	return func(o *Option) {
		o.Temperature = temperature
	}
}