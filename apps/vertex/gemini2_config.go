package vertex

import (
	"github.com/kade-chen/library/tools/generics"
	"google.golang.org/genai"
	// genai2 "google.golang.org/genai"
)

type Gemini2Config struct {
	ModelName             string `json:"model_name"`
	GenerateContentConfig *genai.GenerateContentConfig
	Contents              []*genai.Content
}

// this method is default Gemini2Config
func NewDefaultsGemini2Config() *Gemini2Config {
	return &Gemini2Config{}
}

// this method is to configure the default Gemini2Config configuration
func NewGemini2Config() *Gemini2Config {
	return &Gemini2Config{
		ModelName:             "gemini-2.0-flash-001",
		GenerateContentConfig: NewTestGenerateContentConfig(),
	}
}

// this method is to configure the default gemini configuration
func NewGenerateContentConfig() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		Temperature:     generics.Generics(2.0),
		TopP:            generics.Generics(0.0),
		TopK:            generics.Generics(0.0),
		MaxOutputTokens: generics.Generics[int64](8192),
	}
}

// this method is to test configure the default gemini configuration
func NewTestGenerateContentConfig() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		// SystemInstruction: &genai.Content{
		// 	Parts: []*genai.Part{},
		// },
		Temperature: generics.Generics(0.0),
		TopP:        generics.Generics(1.0),
		TopK:        generics.Generics(0.0),
		//must be set to 1
		// CandidateCount:  1,
		MaxOutputTokens: generics.Generics[int64](8192),
		// StopSequences:      []string{},
		//logprobs 参数指的是模型生成文本的对数概率。它是一个列表，其中每个元素对应于生成文本中每个 token 的对数概率。
		//logprobs 参数让你访问 Gemini 模型生成的文本中每个 token 的对数概率，为你提供更深入的理解和控制。它在评估模型的不确定性、调试问题、分析文本和控制生成过程等方面都有重要的应用价值。
		ResponseLogprobs: false,
		//[0-5]
		// Logprobs: generics.Generics[int64](5.0),
		//[-2,2) 存在
		PresencePenalty: generics.Generics(1.9),
		//[-2,2) 概率
		FrequencyPenalty: generics.Generics(0.0),
		//TYPE_INT32
		//seed 参数（通常称为随机种子）主要用于控制模型生成结果的随机性和可复现性
		//简单来说，seed 值就像一个控制随机性的开关。它让你在需要时可以锁定模型的行为，以便更好地控制和理解它的输出。
		//默认是这个随机数种子
		Seed: generics.Generics[int64](-2147483648),
		// ResponseMIMEType: "application/json",
		/*
			ResponseSchema: &genai.Schema{
				Type: "object",
				Properties: map[string]*genai.Schema{
					"name": {
						Type:        "string",
						Description: "菜名",
					},
					"time": {
						Type:        "integer",
						Description: "用时",
					},
					"do": {
						Type: "string",
						// Format:      "email",
						Description: "做法",
					},
				},
				Required: []string{"name", "time", "do"},
			},
		*/
		//Configuration for model router requests.
		// RoutingConfig: &genai.GenerationConfigRoutingConfig{
		// 	AutoMode: &genai.GenerationConfigRoutingConfigAutoRoutingMode{},
		// 	// ManualMode: &genai.GenerationConfigRoutingConfigManualRoutingMode{},
		// },
		SafetySettings: []*genai.SafetySetting{
			{
				Method:    genai.HarmBlockMethodProbability,
				Category:  genai.HarmCategoryHateSpeech,
				Threshold: genai.HarmBlockThresholdOff,
			},
			{
				Method:    genai.HarmBlockMethodProbability,
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockThresholdOff,
			},
			{
				Method:    genai.HarmBlockMethodProbability,
				Category:  genai.HarmCategoryHarassment,
				Threshold: genai.HarmBlockThresholdOff,
			},
			{
				Method:    genai.HarmBlockMethodProbability,
				Category:  genai.HarmCategorySexuallyExplicit,
				Threshold: genai.HarmBlockThresholdOff,
			},
			{
				Method:    genai.HarmBlockMethodProbability,
				Category:  genai.HarmCategoryCivicIntegrity,
				Threshold: genai.HarmBlockThresholdOff,
			},
		},

		// Tools:              []*genai.Tool{},
		// ToolConfig:         &genai.ToolConfig{},
		//Resource name of a context cache that can be used in subsequent requests.
		// CachedContent:      "",
		//可以返回的模态集。
		ResponseModalities: []string{"TEXT"}, // "AUDIO", "IMAGE",
		// ResponseModalities: []string{"TEXT", "IMAGE"},

		//如果指定，将使用指定的媒体分辨率。
		// MediaResolution: genai.MediaResolutionHigh,
		//The speech generation configuration.
		SpeechConfig: &genai.SpeechConfig{
			VoiceConfig: &genai.VoiceConfig{
				PrebuiltVoiceConfig: &genai.PrebuiltVoiceConfig{
					VoiceName: "Aoede",
				},
			},
		},
	}
}
