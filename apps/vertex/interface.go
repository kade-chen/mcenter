package vertex

import (
	"context"
	"iter"

	"cloud.google.com/go/vertexai/genai"
	gemini2 "google.golang.org/genai"
)

const (
	AppName  = "vertex"
	AppName2 = "gemini2"
	AppName3 = "vertexplatform"
)

type Service interface {
	NoStreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) (*genai.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *GenaiSetting, ...genai.Part) *genai.GenerateContentResponseIterator
}

type ServiceGemini2 interface {
	NoStreamingGenerateContent(context.Context, *Gemini2Config) (*gemini2.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *Gemini2Config) iter.Seq2[*gemini2.GenerateContentResponse, error]
	// GenerateContent(context.Context, string, []*gemini2.Content, *gemini2.GenerateContentConfig) (*gemini2.GenerateContentResponse, error)
}

type ServiceGemini3 interface {
	NoStreamingGenerateContent(context.Context, *Gemini2Config) (*gemini2.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *Gemini2Config) (iter.Seq2[*gemini2.GenerateContentResponse, error], error)
	// GenerateContent(context.Context, string, []*gemini2.Content, *gemini2.GenerateContentConfig) (*gemini2.GenerateContentResponse, error)
}
