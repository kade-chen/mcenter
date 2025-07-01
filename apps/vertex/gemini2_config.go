package vertex

import (
	"github.com/kade-chen/library/tools/generics"
	"google.golang.org/genai"
	// genai2 "google.golang.org/genai"
)

type Gemini_Config struct {
	ModelName             string `json:"model_name"`
	Location              string `json:"location"`
	GenerateContentConfig *genai.GenerateContentConfig
	Contents              []*genai.Content
}

// this method is default Gemini2Config
func NewDefaultsGemini2Config() *Gemini_Config {
	return &Gemini_Config{}
}

// this method is to configure the default Gemini2Config configuration
func NewGemini2Config() *Gemini_Config {
	return &Gemini_Config{
		ModelName:             "gemini-2.0-flash-lite-001",
		GenerateContentConfig: NewTestGenerateContentConfig(),
	}
}

// this method is to configure the default gemini configuration
func NewGenerateContentConfig() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		Temperature:     generics.Generics[float32](2.0),
		TopP:            generics.Generics[float32](0.0),
		TopK:            generics.Generics[float32](1),
		MaxOutputTokens: 8192,
	}
}

// this method is to test configure the default gemini configuration
func NewTestGenerateContentConfig() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		// SystemInstruction: &genai.Content{
		// 	Role: "",
		// 	//https://cloud.google.com/vertex-ai/generative-ai/docs/reference/rpc/google.cloud.aiplatform.v1#google.cloud.aiplatform.v1.Part
		// 	Parts: []*genai.Part{
		// 		// genai.Text("1.细节都描写出来，包括颜色，眼神，心理，格式也必须分 2.分点 3.必须中文回答"),
		// 		// genai.FileData{},
		// 		// genai.Blob{},
		// 	},
		// },
		//温度temperature [0-2.0001], The default is 0.
		Temperature: generics.Generics[float32](0.0),
		// 核采样原理：
		// 核采样是一种基于核函数的采样方法。它不像简单的随机采样那样随机抽取样本，而是根据样本之间的相似性来进行采样。
		// 核心思想是：相似的数据点，只需要采样其中一个代表即可，而相异的数据点则需要更多样本进行代表。 这使得核采样能够有效地减少样本数量，同时保留数据的关键信息。
		// 具体来说，核采样会先计算样本之间的相似度，通常用核函数来衡量。常用的核函数包括高斯核、线性核等等。核函数会将原始数据映射到一个高维特征空间，在这个空间中，样本之间的相似度更容易计算。然后，算法会根据相似度选择具有代表性的样本进行采样。 那些在高维空间中“聚集”在一起的样本，只需要保留一个或少数几个代表即可；而那些在高维空间中“分散”的样本，则需要保留更多样本以保证信息不丢失
		// 指定较低的值可减少响应的随机性，指定较高的值可增加响应的随机性。
		//TopP [0-1], The default is 0.
		// Default for gemini-1.5-flash: 0.95
		// Default for gemini-1.5-pro: 0.95
		// Default for gemini-1.0-pro: 1.0
		// Default for gemini-1.0-pro-vision: 1.0
		TopP: generics.Generics[float32](1.0),
		TopK: generics.Generics[float32](1.0),
		//这里采用核采样，不采用抽样采样
		//TopK 算法的核心思想是避免对整个数据集进行完全排序，而是利用数据结构和算法的特性，高效地找到前 K 个最大或最小元素。
		// 指定较低的值可减少响应的随机性，指定较高的值可增加响应的随机性。
		// 对于每个 token 选择步骤，都会对具有最高概率的前 K 个 token 进行采样。然后根据 top-P 进一步过滤 token，并使用温度采样选择最终的 token。
		//TopK [1-41), The default is 0. //仅受支持gemini-1.0-pro-vision。
		// TopK: 1,
		//CandidateCount [1-9), The default is 1. candidate count
		// CandidateCount:  1,
		//MaxOutputTokens [1,8193) ，The default is 0.
		MaxOutputTokens: 8192,
		// StopSequences:      []string{},
		//logprobs 参数指的是模型生成文本的对数概率。它是一个列表，其中每个元素对应于生成文本中每个 token 的对数概率。
		//logprobs 参数让你访问 Gemini 模型生成的文本中每个 token 的对数概率，为你提供更深入的理解和控制。它在评估模型的不确定性、调试问题、分析文本和控制生成过程等方面都有重要的应用价值。
		ResponseLogprobs: false,
		//[0-5]
		// Logprobs: generics.Generics[int64](5.0),
		//Stop Sequences ，[1,17)，The default is [""].
		// StopSequences: []string{"阴道"},
		//可选参数，指的是存在惩罚。 在大型语言模型（LLM）的生成过程中，它控制模型对已生成文本中出现过的词语或概念的重复使用进行惩罚。 数值越高，模型越倾向于避免重复使用已出现的词语或概念，从而鼓励生成更具多样性和创造性的文本。 反之，数值越低或为负值，模型越倾向于重复使用已出现的词语或概念。
		// PresencePenalty [-2，2), The default is 0.
		// 这个参数惩罚模型生成在文本中已经出现过的词语。值越高，模型越倾向于避免使用已经出现过的词语，从而鼓励生成更广泛、更具多样性的文本。值越低，模型越有可能重复使用相同的词语。
		//[-2,2) 存在
		PresencePenalty: generics.Generics[float32](1.9),
		// FrequencyPenalty [-2，2), The default is 0.
		// 这个参数惩罚模型生成在文本中出现频率高的词语。值越高，模型越倾向于避免使用出现频率高的词语，即使这些词语之前已经出现过。这有助于减少重复，并鼓励模型使用更广泛的词汇。值越低，模型越有可能重复使用高频词语。
		//[-2,2) 概率
		FrequencyPenalty: generics.Generics[float32](0.0),
		//TYPE_INT32
		//seed 参数（通常称为随机种子）主要用于控制模型生成结果的随机性和可复现性
		//简单来说，seed 值就像一个控制随机性的开关。它让你在需要时可以锁定模型的行为，以便更好地控制和理解它的输出。
		//默认是这个随机数种子
		Seed: generics.Generics[int32](-2147483648),
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
		ThinkingConfig: &genai.ThinkingConfig{
			// "message": "Error 400, Message: Unable to submit request because thinking is not configurable in this model; please remove the thinking_config setting and try it again. Learn more: https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/gemini, Status: INVALID_ARGUMENT, Details: []"
			// IncludeThoughts: true,
			// ThinkingBudget:  generics.Generics[int32](0),
		},
	}
}
