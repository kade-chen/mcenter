package impl

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/library/tools/format"
	"github.com/kade-chen/mcenter/apps/vertex"
)

func (g *gemini) NoStreamingGenerateContent(ctx context.Context, geminiSetting *vertex.GenaiSetting, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	//0.get regional client
	g.log.Info().Msgf("get regional client: %s", vertex.GetLocationString(geminiSetting.Regional))
	client := g.clients[geminiSetting.Regional]

	//1.for the regional client, get the generative model
	g.log.Info().Msgf("get generative model: %s", geminiSetting.ModelName)
	gemini := client.GenerativeModel(geminiSetting.ModelName)

	//optional, set safety settings
	if geminiSetting.SafetySetting != nil {
		gemini.SafetySettings = NewGeminiSafetySettings(geminiSetting)
		g.log.Info().Msgf("set safety settings: %#v", format.MustToYaml(geminiSetting.SafetySetting))
	}
	//optional, set generation config
	if geminiSetting.GenerationConfig != nil {
		gemini.GenerationConfig = NewGeminiGenerationConfig(geminiSetting)
		g.log.Info().Msgf("set generation config: %#v", geminiSetting.GenerationConfig)
	}
	//optional, set system instruction
	if geminiSetting.SystemInstruction != nil {
		gemini.SystemInstruction = NewGeminiSystemInstruction(geminiSetting)
		g.log.Info().Msgf("set system instruction: %#v", geminiSetting.SystemInstruction)
	}

	// resp, err := gemini.GenerateContent(ctx, parts...)
	// if err != nil {
	// 	return nil, fmt.Errorf("error generating content: %w", err)
	// }

	// B := make(map[int]interface{}, 10)
	// // B := make([]any, 10)
	// for q, v := range resp.Candidates {
	// 	B[q] = v.Content.Parts[0]
	// }

	// count := make(map[any]int, 10)

	// fmt.Println(B)
	// fmt.Println(resp.Candidates[0].Content.Parts)
	// rb, err := json.MarshalIndent(resp, "", "  ")
	// // rb, err := json.Marshal(resp)
	// if err != nil {
	// 	return nil, fmt.Errorf("json.MarshalIndent: %w", err)
	// }
	// fmt.Println(string(rb))
	return gemini.GenerateContent(ctx, parts...)
}

func (g *gemini) StreamingGenerateContent(ctx context.Context, geminiSetting *vertex.GenaiSetting, parts ...genai.Part) *genai.GenerateContentResponseIterator {
	//0.get regional client
	g.log.Info().Msgf("get regional client: %s", vertex.GetLocationString(geminiSetting.Regional))
	client := g.clients[geminiSetting.Regional]
	//1.for the regional client, get the generative model
	g.log.Info().Msgf("get generative model: %s", geminiSetting.ModelName)
	gemini := client.GenerativeModel(geminiSetting.ModelName)
	//optional, set safety settings
	if geminiSetting.SafetySetting != nil {
		gemini.SafetySettings = NewGeminiSafetySettings(geminiSetting)
		g.log.Info().Msgf("set safety settings: %#v", format.MustToYaml(geminiSetting.SafetySetting))
	}
	//optional, set generation config
	if geminiSetting.GenerationConfig != nil {
		gemini.GenerationConfig = NewGeminiGenerationConfig(geminiSetting)
		g.log.Info().Msgf("set generation config: %#v", geminiSetting.GenerationConfig)
	}
	//optional, set system instruction
	if geminiSetting.SystemInstruction != nil {
		gemini.SystemInstruction = NewGeminiSystemInstruction(geminiSetting)
		g.log.Info().Msgf("set system instruction: %#v", geminiSetting.SystemInstruction)
	}
	//2. generate streaming content
	return gemini.GenerateContentStream(ctx, parts...)
	/*
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
				return "", exception.NewInternalServerError("error: %v", err)
			}
			if err != nil {
				return "", exception.NewInternalServerError("error: %v", err)
			}

			// fmt.Fprint(os.Stdout, "generated response: ")
			for _, c := range resp.Candidates {
				for _, p := range c.Content.Parts {
					fmt.Fprintf(os.Stdout, "%v", p)
					// fmt.Fprint(os.Stdout, p)
					// fmt.Print(p)
				}
			}
			fmt.Fprint(os.Stdout, "\n")
		}
		return "OK, No More items in iterator", nil
	*/
}
