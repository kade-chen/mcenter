package vertex

import (
	"cloud.google.com/go/vertexai/genai"
)

type GenaiSetting struct {
	//https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/inference#request
	SafetySetting []*genai.SafetySetting
	//https://cloud.google.com/vertex-ai/generative-ai/docs/reference/rpc/google.cloud.aiplatform.v1#google.cloud.aiplatform.v1.Content
	//Available for gemini-1.5-flash, gemini-1.5-pro, and gemini-1.0-pro-002.
	SystemInstruction *Content
	GenerationConfig  *GenerationConfig
	Regional          LOCATION `json:"regional"`
	ModelName         string   `json:"model_name"`
	Text              string   `json:"text"`
	Url               string   `json:"url"`
	BinaryData        []byte   `json:"binary_data"`
}

type GenerationConfig struct {
	// Optional. Controls the randomness of predictions.
	Temperature float32 `json:"temperature"`
	// Optional. If specified, nucleus sampling will be used.
	TopP float32 `json:"top_p"`
	// Optional. If specified, top-k sampling will be used.
	TopK int32 `json:"top_k"`
	// Optional. Number of candidates to generate.
	CandidateCount int32 `json:"candidate_count"`
	// Optional. The maximum number of output tokens to generate per message.
	MaxOutputTokens int32 `json:"max_output_tokens"`
	// Optional. Stop sequences.
	StopSequences []string `json:"stop_sequences"`
	// Optional. Positive penalties.
	PresencePenalty float32 `json:"presence_penalty"`
	// Optional. Frequency penalties.
	FrequencyPenalty float32 `json:"frequency_penalty"`
	// Optional. Output response mimetype of the generated candidate text.
	// Supported mimetype:
	// - `text/plain`: (default) Text output.
	// - `application/json`: JSON response in the candidates.
	// The model needs to be prompted to output the appropriate response type,
	// otherwise the behavior is undefined.
	// This is a preview feature.
	ResponseMIMEType string `json:"response_mime_type"`
	// Optional. The `Schema` object allows the definition of input and output
	// data types. These types can be objects, but also primitives and arrays.
	// Represents a select subset of an [OpenAPI 3.0 schema
	// object](https://spec.openapis.org/oas/v3.0.3#schema).
	// If set, a compatible response_mime_type must also be set.
	// Compatible mimetypes:
	// `application/json`: Schema for JSON response.
	ResponseSchema *genai.Schema
}

type Content struct {
	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	// Useful to set for multi-turn conversations, otherwise can be left blank
	// or unset.
	Role string
	// Required. Ordered `Parts` that constitute a single message. Parts may have
	// different IANA MIME types.
	Parts string
}

// Return a new NewGenaiSetting
func NewGenaiSetting() *GenaiSetting {
	return &GenaiSetting{
		// https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/inference#safetysetting
		SafetySetting: []*genai.SafetySetting{
			{
				Category:  genai.HARM_CATEGORY_CIVICINTEGRITY,
				Threshold: genai.HarmBlockOFF,
				// Method:    genai.HarmBlockMethodProbability,
			},
		},
		SystemInstruction: &Content{},
		// SystemInstruction: &genai.Content{
		// 	Role: "",
		// 	//https://cloud.google.com/vertex-ai/generative-ai/docs/reference/rpc/google.cloud.aiplatform.v1#google.cloud.aiplatform.v1.Part
		// 	Parts: []genai.Part{
		// 		// genai.Text("1.细节都描写出来，包括颜色，眼神，心理，格式也必须分 2.分点 3.必须中文回答"),
		// 		// genai.FileData{},
		// 		// genai.Blob{},
		// 	},
		// },
		GenerationConfig: &GenerationConfig{
			//温度temperature [0-2.0001], The default is 0.
			Temperature: 0,
			// 核采样原理：
			// 核采样是一种基于核函数的采样方法。它不像简单的随机采样那样随机抽取样本，而是根据样本之间的相似性来进行采样。
			// 核心思想是：相似的数据点，只需要采样其中一个代表即可，而相异的数据点则需要更多样本进行代表。 这使得核采样能够有效地减少样本数量，同时保留数据的关键信息。
			// 具体来说，核采样会先计算样本之间的相似度，通常用核函数来衡量。常用的核函数包括高斯核、线性核等等。核函数会将原始数据映射到一个高维特征空间，在这个空间中，样本之间的相似度更容易计算。然后，算法会根据相似度选择具有代表性的样本进行采样。 那些在高维空间中“聚集”在一起的样本，只需要保留一个或少数几个代表即可；而那些在高维空间中“分散”的样本，则需要保留更多样本以保证信息不丢失
			//TopP [0-1], The default is 0.
			TopP: 1,
			//这里采用核采样，不采用抽样采样
			//TopK 算法的核心思想是避免对整个数据集进行完全排序，而是利用数据结构和算法的特性，高效地找到前 K 个最大或最小元素。
			//TopK [1-41), The default is 0.
			// TopK: 1,
			//CandidateCount [1-9), The default is 1. candidate count
			CandidateCount: 1,
			//MaxOutputTokens [1,8193) ，The default is 0.
			MaxOutputTokens: 8192,
			//Stop Sequences ，[1,17)，The default is [""].
			StopSequences: []string{"ll"},
			//可选参数，指的是存在惩罚。 在大型语言模型（LLM）的生成过程中，它控制模型对已生成文本中出现过的词语或概念的重复使用进行惩罚。 数值越高，模型越倾向于避免重复使用已出现的词语或概念，从而鼓励生成更具多样性和创造性的文本。 反之，数值越低或为负值，模型越倾向于重复使用已出现的词语或概念。
			// PresencePenalty [-2，2), The default is 0.
			// 这个参数惩罚模型生成在文本中已经出现过的词语。值越高，模型越倾向于避免使用已经出现过的词语，从而鼓励生成更广泛、更具多样性的文本。值越低，模型越有可能重复使用相同的词语。
			PresencePenalty: -2,
			// FrequencyPenalty [-2，2), The default is 0.
			// 这个参数惩罚模型生成在文本中出现频率高的词语。值越高，模型越倾向于避免使用出现频率高的词语，即使这些词语之前已经出现过。这有助于减少重复，并鼓励模型使用更广泛的词汇。值越低，模型越有可能重复使用高频词语。
			FrequencyPenalty: 1.9,
			// ResponseMIMEType: "application/json",
			// ResponseSchema: &genai.Schema{
			// 	Type: genai.TypeArray,
			// 	Items: &genai.Schema{
			// 		Type: genai.TypeObject,
			// 		Properties: map[string]*genai.Schema{
			// 			"图片": {
			// 				Type: genai.TypeString,
			// 			},
			// 		},
			// 		Required: []string{"图片"},
			// 	},
			// },
		},
	}
}
