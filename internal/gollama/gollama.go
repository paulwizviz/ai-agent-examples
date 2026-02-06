package gollama

import (
	"encoding/base64"
)

const (
	DefaultLocalBaseURL = "http://localhost:11434"
)

type RequestGenerate struct {
	Model     string   `json:"model"`
	Prompt    string   `json:"prompt"`
	Suffix    string   `json:"suffix,omitempty"`
	Images    []string `json:"images,omitempty"`
	Format    string   `json:"format,omitempty"`
	Options   any      `json:"options,omitempty"`
	System    string   `json:"system,omitempty"`
	Raw       bool     `json:"raw,omitempty"`
	KeepAlive string   `json:"keep_alive,omitempty"`
	Stream    bool     `json:"stream"`
}

type ResponseGenerate struct {
	Model              string `json:"model"`
	Created            string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

type Client interface {
	GenerateAPI(request RequestGenerate) ([]ResponseGenerate, error)
}

func NewDefaultClient(timeout int, baseURL string) Client {
	return &defaultClient{
		baseURL: baseURL,
		timeout: timeout,
	}
}

func EncodeBase64(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}
