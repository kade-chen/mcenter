package vertex

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
)

const (
	AppName  = "vertex"
	AppName2 = "gemini2"
)

type Service interface {
	NoStreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) (*genai.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) *genai.GenerateContentResponseIterator
}

type ServiceGemini2 interface {
	NoStreamingGenerateContent(context.Context)
	// GenerateContent(context.Context, string, []*gemini2.Content, *gemini2.GenerateContentConfig) (*gemini2.GenerateContentResponse, error)
}
