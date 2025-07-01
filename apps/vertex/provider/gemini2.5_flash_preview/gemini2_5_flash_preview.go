package gemini25flashpreview

import (
	"context"
	"iter"

	"github.com/kade-chen/library/exception"

	"google.golang.org/genai"

	"github.com/kade-chen/mcenter/apps/vertex"
)

func (g *Gemini_25_flash_preview) NoStreamingGenerateContent(ctx context.Context, gemini2Config *vertex.Gemini_Config) (*genai.GenerateContentResponse, error) {
	// fmt.Println("----Location", gemini2Config.Location)
	a, exists := vertex.Gemini_25_FLASH_PREVIEW_String_To_LOCATION[gemini2Config.Location]
	if !exists {
		return nil, exception.NewBadRequest("Model Name: %s && Location: %s , And cannot be empty or no in the registry", gemini2Config.ModelName, gemini2Config.Location)
	}
	// // fmt.Println(format.ToJSON(re))
	// fmt.Println(format.ToJSON(err))
	return g.gemini_clients[a].Models.GenerateContent(ctx, gemini2Config.ModelName, gemini2Config.Contents, gemini2Config.GenerateContentConfig)
}

func (g *Gemini_25_flash_preview) StreamingGenerateContent(ctx context.Context, gemini2Config *vertex.Gemini_Config) iter.Seq2[*genai.GenerateContentResponse, error] {
	a, exists := vertex.Gemini_25_FLASH_PREVIEW_String_To_LOCATION[gemini2Config.Location]
	if !exists {
		return (func(yield func(*genai.GenerateContentResponse, error) bool) {
			yield(nil, exception.NewBadRequest("Model Name: %s && Location: %s , And cannot be empty or no in the registry", gemini2Config.ModelName, gemini2Config.Location))
		})
	}
	return g.gemini_clients[a].Models.GenerateContentStream(ctx, gemini2Config.ModelName, gemini2Config.Contents, gemini2Config.GenerateContentConfig)
}

// return (func(yield func(*genai.GenerateContentResponse, error) bool) {
// 	yield(nil, fmt.Errorf("invalid location: %s", gemini2Config.Location))
// })

// return func(yield func(*genai.GenerateContentResponse, error) bool) {
// 		for i := 0; i < 1; i++ {
// 			resp := &genai.GenerateContentResponse{
// 				Candidates: []*genai.Candidate{
// 					{
// 						Content: &genai.Content{
// 							Parts: []*genai.Part{
// 								{Text: "sss"},
// 							},
// 							Role: "user",
// 						},
// 					},
// 				},
// 			}
// 			result := yield(resp, nil)
// 			fmt.Println("yield returned:", result)

// 			if !result {
// 				fmt.Println("调用者不想继续了，停止流式生成")
// 				return
// 			}
// 		}
// 	}
