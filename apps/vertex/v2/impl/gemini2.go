package impl

import (
	"context"
	"iter"

	"github.com/kade-chen/mcenter/apps/vertex"
	"google.golang.org/genai"
)

func (g *gemini) NoStreamingGenerateContent(ctx context.Context, config *vertex.Gemini2Config) (*genai.GenerateContentResponse, error) {
	return g.clientsv2.Models.GenerateContent(ctx, config.ModelName, config.Contents, config.GenerateContentConfig)
}

func (g *gemini) StreamingGenerateContent(ctx context.Context, config *vertex.Gemini2Config) iter.Seq2[*genai.GenerateContentResponse, error] {
	// config := vertex.NewTestGenerateContentConfig()
	return g.clientsv2.Models.GenerateContentStream(ctx, config.ModelName, config.Contents, config.GenerateContentConfig)
}
