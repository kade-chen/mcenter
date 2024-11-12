package vertex

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
)

const (
	AppName = "vertex"
)

type Service interface {
	NoStreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) (*genai.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) *genai.GenerateContentResponseIterator
}
