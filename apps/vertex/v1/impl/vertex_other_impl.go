package impl

import (
	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/mcenter/apps/vertex"
)

// Returns a new GenerationConfig
func NewGeminiGenerationConfig(geminiSetting *vertex.GenaiSetting) genai.GenerationConfig {
	Config := genai.GenerationConfig{
		Temperature: &geminiSetting.GenerationConfig.Temperature,
		//这里采用核采样，不采用抽样采样
		TopP: &geminiSetting.GenerationConfig.TopP,
		// TopK: &geminiSetting.GenerationConfig.TopK,
		CandidateCount:  &geminiSetting.GenerationConfig.CandidateCount,
		MaxOutputTokens: &geminiSetting.GenerationConfig.MaxOutputTokens,
		// StopSequences:    geminiSetting.GenerationConfig.StopSequences,
		PresencePenalty:  &geminiSetting.GenerationConfig.PresencePenalty,
		FrequencyPenalty: &geminiSetting.GenerationConfig.FrequencyPenalty,
		ResponseMIMEType: geminiSetting.GenerationConfig.ResponseMIMEType,
		// ResponseSchema:   geminiSetting.GenerationConfig.ResponseSchema,
	}
	if geminiSetting.GenerationConfig.ResponseSchema != nil {
		Config.ResponseSchema = geminiSetting.GenerationConfig.ResponseSchema
	}
	if geminiSetting.GenerationConfig.StopSequences != nil {
		Config.StopSequences = geminiSetting.GenerationConfig.StopSequences
	}
	return Config
}

// Returns a new SafetySetting
func NewGeminiSafetySettings(geminiSetting *vertex.GenaiSetting) []*genai.SafetySetting {
	SafetySettings := []*genai.SafetySetting{
		{
			Category:  geminiSetting.SafetySetting[0].Category,
			Threshold: geminiSetting.SafetySetting[0].Threshold,
			Method:    geminiSetting.SafetySetting[0].Method,
		},
	}
	return SafetySettings
}

// Returns a new SystemInstruction
func NewGeminiSystemInstruction(geminiSetting *vertex.GenaiSetting) *genai.Content {
	// if geminiSetting.SystemInstruction.Parts
	SystemInstruction := &genai.Content{
		Role: geminiSetting.SystemInstruction.Role,
		//The text setting count toward the token limit
		Parts: []genai.Part{
			genai.Text(geminiSetting.SystemInstruction.Parts),
		},
	}

	return SystemInstruction
}
