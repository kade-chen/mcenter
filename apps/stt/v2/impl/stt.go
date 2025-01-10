package impl

import (
	"context"
	"fmt"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/stt"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/cloud/location"
)

// List Locations lists information about the supported locations for this service.
func (s *speechToTextV2) ListLocations(ctx context.Context) {
	locationIterator := s.clinets[0].ListLocations(ctx, &location.ListLocationsRequest{
		Name:     fmt.Sprintf("projects/%s", s.ProjectID),
		Filter:   "asia-southeast1",
		PageSize: 13,
	})
	// Locations, nextPageToken, err := locationIterator.InternalFetch(10, "")
	// if err != nil {
	// 	// 处理错误（可以选择跳出循环）
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(Locations, nextPageToken)
	for {
		Locationmessage, err := locationIterator.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			// 处理错误（可以选择跳出循环）
			fmt.Println("Error:", err)
			break
		}

		// 使用 `recognizer`
		fmt.Println(Locationmessage)
	}
}

// List Recognizers lists information about the supported recognizers for this service.
func (s *speechToTextV2) ListRecognizers(ctx context.Context, endpoint stt.ENDPOINT) error {
	iterators := s.clinets[endpoint].ListRecognizers(ctx, &speechpb.ListRecognizersRequest{
		Parent: s.parentString(s.ProjectID, endpoint),
		// Parent: fmt.Sprintf("projects/%s/locations/%s", s.ProjectID, stt.ENDPOINTS_PARENT[endpoint]),
	})
	// Recognizers, nextPageToken, err := iterators.InternalFetch(10, "")
	// if err != nil {
	// 	// 处理错误（可以选择跳出循环）
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(Recognizers, nextPageToken)
	for {
		recognizer, err := iterators.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// 处理错误（可以选择跳出循环）
			return exception.NewInternalServerError("Error: %s", err)
		}

		// 使用 `recognizer`
		fmt.Println("Recognizer:", recognizer)
	}
	return nil
}

func (s *speechToTextV2) CreateRecognizer(ctx context.Context, endpoint stt.ENDPOINT) {
	//https://cloud.google.com/speech-to-text/v2/docs/reference/rpc/google.cloud.speech.v2#createrecognizerrequest
	b, err := s.clinets[endpoint].CreateRecognizer(ctx, &speechpb.CreateRecognizerRequest{
		//https://cloud.google.com/speech-to-text/v2/docs/reference/rpc/google.cloud.speech.v2#recognizer
		Recognizer: &speechpb.Recognizer{
			//User-settable, human-readable name for the recognizer.
			DisplayName: "kadecc",
			DefaultRecognitionConfig: &speechpb.RecognitionConfig{
				//Decoding parameters for audio being sent for recognition.
				DecodingConfig: &speechpb.RecognitionConfig_AutoDecodingConfig{
					AutoDecodingConfig: &speechpb.AutoDetectDecodingConfig{},
				},
				//https://cloud.google.com/speech-to-text/v2/docs/transcription-model#transcription_models
				Model:         "long",
				LanguageCodes: []string{"cmn-Hans-CN", "en_au"},
				//启用语音识别功能
				//https://cloud.google.com/speech-to-text/v2/docs/reference/rpc/google.cloud.speech.v2#recognitionfeatures
				Features: &speechpb.RecognitionFeatures{
					//If set to true, the server will attempt to filter out profanities, replacing all but the initial character in each filtered word with asterisks, for instance, "f***". If set to false or omitted, profanities won't be filtered out.
					ProfanityFilter:       true,
					EnableWordTimeOffsets: false,
					//如果true，则顶部结果包括单词列表和这些单词的置信度。如果false，则不返回单词级置信度信息。默认值为false。
					EnableWordConfidence: true,
					//如果true，则在识别结果假设中添加标点符号。此功能仅适用于部分语言。默认false值为不向结果假设添加标点符号
					EnableAutomaticPunctuation: true,
					//通话的口语标点符号行为。如果true，则用请求中的相应符号替换口语标点符号。例如，“你好吗问号”变为“你好吗？”。
					//部分language不支持
					//请参阅https://cloud.google.com/speech-to-text/docs/spoken-punctuation获取支持。如果false，则不会替换口语标点符号
					EnableSpokenPunctuation: false,
					// 通话的语音表情符号行为。如果true，则为请求添加语音表情符号格式。这将在最终记录中用相应的 Unicode 符号替换语音表情符号。如果false，则不会替换语音表情符号。
					//部分language不支持
					//https://cloud.google.com/speech-to-text/v2/docs/spoken-emoji
					EnableSpokenEmojis: false,
					//识别多声道音频的模式。
					//MULTI_CHANNEL_MODE_UNSPECIFIED-多声道模式的默认值。如果音频包含多个声道，则只转录第一个声道；其他声道将被忽略
					//SEPARATE_RECOGNITION_PER_CHANNEL-model如果选择，则提供的音频中的每个频道都会单独转录。如果选择的 是 ，则无法选择此选项。latest_short
					MultiChannelMode: speechpb.RecognitionFeatures_SEPARATE_RECOGNITION_PER_CHANNEL,
					// DiarizationConfig: &speechpb.SpeakerDiarizationConfig{
					// 必填。对话中发言者的最小数量。此范围可让系统自动确定正确的发言者数量，从而为您提供更大的灵活性。
					// 要修复音频中检测到的扬声器数量，请设置min_speaker_count= max_speaker_count。
					// MinSpeakerCount: 1,
					// 必填。对话中发言者的最大数量。有效值为：1-6。必须 >= min_speaker_count。此范围可让您更加灵活，因为系统可以自动确定正确的发言者数量。
					// MaxSpeakerCount: 6,
					// },
					//Maximum number of recognition hypotheses to be returned. The server may return fewer than max_alternatives. Valid values are 0-30. A value of 0 or 1 will return a maximum of one. If omitted, will return a maximum of one.
					MaxAlternatives: 4,
				},
				// Adaptation: &speechpb.SpeechAdaptation{
				// 	PhraseSets: []*speechpb.SpeechAdaptation_AdaptationPhraseSet{
				// 		{Value: &speechpb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet{
				// 			PhraseSet: "你号",
				// 		}},
				// 	},
				// },
				TranscriptNormalization: &speechpb.TranscriptNormalization{
					Entries: []*speechpb.TranscriptNormalization_Entry{
						{
							Search:        "您好",
							Replace:       "你号",
							CaseSensitive: false,
						},
					},
				},
				//可选。可选配置用于自动将给定的音频翻译为受支持模型所需的语言。
				// TranslationConfig: &speechpb.TranslationConfig{},

			},
			//Allows users to store small amounts of arbitrary data, both the key-value must be < 64 characters.At most 100 annotations
			Annotations: map[string]string{
				"region":   "us-central1",
				"project":  "my-gcp-project",
				"endpoint": "databricks-nginx-target-proxy",
			},
		},
		//if set, validate the request and preview the recognizer, but do not actually create it.
		ValidateOnly: false,
		//The ID to use for the recognizer
		//This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/.
		RecognizerId: "long-recognizer1",
		//Required. The project and location where this recognizer will be created.
		//The expected format is projects/{project}/locations/{location}.
		Parent: s.parentString(s.ProjectID, endpoint),
	})
	if err != nil {
		fmt.Println(err)
	}
	a, _ := b.Metadata()
	fmt.Println(a)
}

func (s *speechToTextV2) GetModel(ctx context.Context, endpoint stt.ENDPOINT) {
	a, err := s.clinets[endpoint].Recognize(ctx, &speechpb.RecognizeRequest{
		// Recognizer: "projects/kade-poc/locations/europe-west4/recognizers/kade88",
		Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
		Config: &speechpb.RecognitionConfig{
			//Decoding parameters for audio being sent for recognition.
			DecodingConfig: &speechpb.RecognitionConfig_AutoDecodingConfig{
				AutoDecodingConfig: &speechpb.AutoDetectDecodingConfig{},
			},
			//https://cloud.google.com/speech-to-text/v2/docs/transcription-model#transcription_models
			Model:         "long",
			LanguageCodes: []string{"cmn-Hans-CN", "en_au"},
			//启用语音识别功能
			//https://cloud.google.com/speech-to-text/v2/docs/reference/rpc/google.cloud.speech.v2#recognitionfeatures
			Features: &speechpb.RecognitionFeatures{
				//If set to true, the server will attempt to filter out profanities, replacing all but the initial character in each filtered word with asterisks, for instance, "f***". If set to false or omitted, profanities won't be filtered out.
				ProfanityFilter:       false,
				EnableWordTimeOffsets: false,
				//如果true，则顶部结果包括单词列表和这些单词的置信度。如果false，则不返回单词级置信度信息。默认值为false。
				EnableWordConfidence: false,
				//如果true，则在识别结果假设中添加标点符号。此功能仅适用于部分语言。默认false值为不向结果假设添加标点符号
				EnableAutomaticPunctuation: false,
				//通话的口语标点符号行为。如果true，则用请求中的相应符号替换口语标点符号。例如，“你好吗问号”变为“你好吗？”。
				//部分language不支持
				//请参阅https://cloud.google.com/speech-to-text/docs/spoken-punctuation获取支持。如果false，则不会替换口语标点符号
				EnableSpokenPunctuation: false,
				// 通话的语音表情符号行为。如果true，则为请求添加语音表情符号格式。这将在最终记录中用相应的 Unicode 符号替换语音表情符号。如果false，则不会替换语音表情符号。
				//部分language不支持
				//https://cloud.google.com/speech-to-text/v2/docs/spoken-emoji
				EnableSpokenEmojis: false,
				//识别多声道音频的模式。
				//MULTI_CHANNEL_MODE_UNSPECIFIED-多声道模式的默认值。如果音频包含多个声道，则只转录第一个声道；其他声道将被忽略
				//SEPARATE_RECOGNITION_PER_CHANNEL-model如果选择，则提供的音频中的每个频道都会单独转录。如果选择的 是 ，则无法选择此选项。latest_short
				MultiChannelMode: speechpb.RecognitionFeatures_SEPARATE_RECOGNITION_PER_CHANNEL,
				// DiarizationConfig: &speechpb.SpeakerDiarizationConfig{
				// 必填。对话中发言者的最小数量。此范围可让系统自动确定正确的发言者数量，从而为您提供更大的灵活性。
				// 要修复音频中检测到的扬声器数量，请设置min_speaker_count= max_speaker_count。
				// MinSpeakerCount: 1,
				// 必填。对话中发言者的最大数量。有效值为：1-6。必须 >= min_speaker_count。此范围可让您更加灵活，因为系统可以自动确定正确的发言者数量。
				// MaxSpeakerCount: 6,
				// },
				//Maximum number of recognition hypotheses to be returned. The server may return fewer than max_alternatives. Valid values are 0-30. A value of 0 or 1 will return a maximum of one. If omitted, will return a maximum of one.
				MaxAlternatives: 4,
			},
			// Adaptation: &speechpb.SpeechAdaptation{
			// 	PhraseSets: []*speechpb.SpeechAdaptation_AdaptationPhraseSet{
			// 		{Value: &speechpb.SpeechAdaptation_AdaptationPhraseSet_PhraseSet{
			// 			PhraseSet: "你号",
			// 		}},
			// 	},
			// },
			TranscriptNormalization: &speechpb.TranscriptNormalization{
				Entries: []*speechpb.TranscriptNormalization_Entry{
					{
						Search:        "您好",
						Replace:       "你号11111,",
						CaseSensitive: false,
					},
				},
			},
			//可选。可选配置用于自动将给定的音频翻译为受支持模型所需的语言。
			// TranslationConfig: &speechpb.TranslationConfig{},

		},
		// Config: &speechpb.RecognitionConfig{
		// 	DecodingConfig: &speechpb.RecognitionConfig_AutoDecodingConfig{
		// 		AutoDecodingConfig: &speechpb.AutoDetectDecodingConfig{},
		// 	},
		// 	// LanguageCodes: []string{"cmn-Hans-CN"},
		// 	// Model:         "long",
		// },
		// ConfigMask:  &field_mask.FieldMask{},
		AudioSource: &speechpb.RecognizeRequest_Uri{Uri: "gs://kadeccc/3.wav"},
	})
	if err != nil {
		fmt.Println(err)
	}
	b := a.GetResults()
	// b = append(b, a.GetResults()...)
	for i, v := range b {
		fmt.Printf("%d,%+v\n", i, b)
		for _, v := range v.Alternatives {
			fmt.Println(v.Transcript)
		}
	}
}

func (s *speechToTextV2) StreamingRecognize(ctx context.Context, endpoint stt.ENDPOINT) (speechpb.Speech_StreamingRecognizeClient, error) {
	return s.clinets[endpoint].StreamingRecognize(ctx)
}

func (s *speechToTextV2) GetServiceAccount() string {
	return s.ServiceAccount
}
