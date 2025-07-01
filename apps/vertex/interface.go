package vertex

import (
	"context"
	"iter"

	gemini2 "google.golang.org/genai"
)

const (
	AppName = "vertex_platform"
)

type Service interface {
	NoStreamingGenerateContent(context.Context, *Gemini_Config) (*gemini2.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *Gemini_Config) (iter.Seq2[*gemini2.GenerateContentResponse, error], error)
	// GenerateContent(context.Context, string, []*gemini2.Content, *gemini2.GenerateContentConfig) (*gemini2.GenerateContentResponse, error)
}
