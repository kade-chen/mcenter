package impl_test

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/speech/apiv1/speechpb"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/stt"
	_ "github.com/kade-chen/mcenter/apps/stt/v1/impl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	ctx  = context.Background()
	impl stt.Service
)

func TestOptimizeSTTShort(t *testing.T) {
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			//https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#audioencoding
			Encoding:        speechpb.RecognitionConfig_ENCODING_UNSPECIFIED, //This field is optional for FLAC and WAV audio files and required for all other audio formats
			SampleRateHertz: 16000,                                           //alid values are: 8000-48000. 16000 is optimal. For best results
			//All encodings support only 1 channel (mono) audio, unless the audio_channel_count and enable_separate_recognition_per_channel fields are set.
			AudioChannelCount:                   1,    //audio channel count enable_separate_recognition_per_channel
			EnableSeparateRecognitionPerChannel: true, //enable_separate_recognition_per_channel
			//https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages
			LanguageCode: "cmn-Hans-CN", //en-AU
			//https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages
			AlternativeLanguageCodes: []string{"en-HK", "en-US", "en-GB"}, //backup language
			//Maximum number of recognition hypotheses to be returned. Specifically
			//the maximum number of SpeechRecognitionAlternative messages within each SpeechRecognitionResult.
			//The server may return fewer than max_alternatives. Valid values are 0-30. A value of 0 or 1 will return a maximum of one. If omitted, will return a maximum of one.
			MaxAlternatives: 2,    //many results,0-30,defult 1,if 0 or 1 will return a maximum of one.
			ProfanityFilter: true, //屏蔽脏话
			//Speech adaptation configuration improves the accuracy of speech recognition
			//这个功能暂时未研究，先注释掉
			Adaptation: &speechpb.SpeechAdaptation{
				PhraseSets: []*speechpb.PhraseSet{
					{
						Phrases: []*speechpb.PhraseSet_Phrase{
							{
								Value: "你好，音频效过",
								Boost: 0,
							},
						},
						Boost: 0,
						Name:  "cc",
					},
				},
				// PhraseSetReferences: []string{"Name:cc"},
			},
			// Optional.  Use transcription normalization to automatically replace parts of
			// Custom replacement words
			TranscriptNormalization: &speechpb.TranscriptNormalization{
				Entries: []*speechpb.TranscriptNormalization_Entry{
					{
						Search:        "音频效果",
						Replace:       "小狗",
						CaseSensitive: false,
					},
				},
			},
			//包含单词和短语“提示”的字符串列表，以便语音识别更有可能识别它们。
			SpeechContexts: []*speechpb.SpeechContext{{
				Phrases: []string{
					"$MONTH", //设置自然语言的提示，在readme.md
				},
				Boost: 20,
			}},
			EnableWordTimeOffsets:      true,                               //是否顶部结果包括单词列表以及这些单词的开始和结束时间偏移（时间戳）。
			EnableWordConfidence:       true,                               //是否返回单词级置信度信息
			EnableAutomaticPunctuation: true,                               //标点
			EnableSpokenPunctuation:    &wrapperspb.BoolValue{Value: true}, //请求中的相应符号替换口语标点符号。例如，“你好吗问号”变为“你好吗？”  defult true
			EnableSpokenEmojis:         &wrapperspb.BoolValue{Value: true}, //是否替换语音表情符号 defult true
			DiarizationConfig: &speechpb.SpeakerDiarizationConfig{
				EnableSpeakerDiarization: true, //启用说话人分类
				MinSpeakerCount:          2,    //说话最少人数量 defult 2
				MaxSpeakerCount:          6,    //说话最多人数量 defult 6
			}, //启用说话人分类
			//https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#google.cloud.speech.v1.RecognitionMetadata
			Metadata: &speechpb.RecognitionMetadata{}, //This item is deprecated!
			// https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionaudio
			// Model: , default automatically selected
			UseEnhanced: true, //设置为 true 以使用增强模型进行语音识别
		},
		// https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionaudio
		Audio: &speechpb.RecognitionAudio{
			// AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: "gs://kadeccc/3.wav"},
		},
	}

	a, err := impl.SpeechToTextShort(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

func TestOptimizeSTTLong(t *testing.T) {
	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			//https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#audioencoding
			Encoding:        speechpb.RecognitionConfig_ENCODING_UNSPECIFIED, //This field is optional for FLAC and WAV audio files and required for all other audio formats
			SampleRateHertz: 16000,                                           //alid values are: 8000-48000. 16000 is optimal. For best results
			//All encodings support only 1 channel (mono) audio, unless the audio_channel_count and enable_separate_recognition_per_channel fields are set.
			AudioChannelCount:                   1,    //audio channel count enable_separate_recognition_per_channel
			EnableSeparateRecognitionPerChannel: true, //enable_separate_recognition_per_channel
			//https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages
			LanguageCode: "cmn-Hans-CN", //af-ZA
			//https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages
			AlternativeLanguageCodes: []string{"en-HK", "en-US", "en-GB"}, //backup language
			//Maximum number of recognition hypotheses to be returned. Specifically
			//the maximum number of SpeechRecognitionAlternative messages within each SpeechRecognitionResult.
			//The server may return fewer than max_alternatives. Valid values are 0-30. A value of 0 or 1 will return a maximum of one. If omitted, will return a maximum of one.
			MaxAlternatives: 30,   //many results,0-30,defult 1,if 0 or 1 will return a maximum of one.
			ProfanityFilter: true, //屏蔽脏话
			//Speech adaptation configuration improves the accuracy of speech recognition
			//这个功能暂时未研究，先注释掉
			Adaptation: &speechpb.SpeechAdaptation{
				PhraseSets: []*speechpb.PhraseSet{
					{
						Phrases: []*speechpb.PhraseSet_Phrase{
							{
								Value: "你好，音频效过",
								Boost: 0,
							},
						},
						Boost: 0,
						Name:  "cc",
					},
				},
				// PhraseSetReferences: []string{"Name:cc"},
			},
			// Optional.  Use transcription normalization to automatically replace parts of
			// Custom replacement words
			TranscriptNormalization: &speechpb.TranscriptNormalization{
				Entries: []*speechpb.TranscriptNormalization_Entry{
					{
						Search:        "您好，音频效果",
						Replace:       "小狗",
						CaseSensitive: false,
					},
				},
			},
			//包含单词和短语“提示”的字符串列表，以便语音识别更有可能识别它们。
			SpeechContexts: []*speechpb.SpeechContext{{
				Phrases: []string{
					"$MONTH", //设置自然语言的提示，在readme.md
				},
				Boost: 20,
			}},
			EnableWordTimeOffsets:      true,                               //是否顶部结果包括单词列表以及这些单词的开始和结束时间偏移（时间戳）。
			EnableWordConfidence:       true,                               //是否返回单词级置信度信息
			EnableAutomaticPunctuation: true,                               //标点
			EnableSpokenPunctuation:    &wrapperspb.BoolValue{Value: true}, //请求中的相应符号替换口语标点符号。例如，“你好吗问号”变为“你好吗？”  defult true
			EnableSpokenEmojis:         &wrapperspb.BoolValue{Value: true}, //是否替换语音表情符号 defult true
			DiarizationConfig: &speechpb.SpeakerDiarizationConfig{
				EnableSpeakerDiarization: true, //启用说话人分类
				MinSpeakerCount:          2,    //说话最少人数量 defult 2
				MaxSpeakerCount:          6,    //说话最多人数量 defult 6
			}, //启用说话人分类
			//https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#google.cloud.speech.v1.RecognitionMetadata
			Metadata: &speechpb.RecognitionMetadata{}, //This item is deprecated!
			// https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionaudio
			// Model: , default automatically selected
			UseEnhanced: true, //设置为 true 以使用增强模型进行语音识别
		},
		// https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionaudio
		Audio: &speechpb.RecognitionAudio{
			// AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: "gs://kadeccc/3.wav"},
		},
	}
	a, err := impl.SpeechToTextLong(ctx, req)
	fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

// The method is not support
func TestOptimizeSTTStreamingRecognize(t *testing.T) {
	req := &speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "cmn-Hans-CN",
				},
				InterimResults: true,
			},
		},
	}
	a, err := impl.SpeechToTextStreamingRecognize(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "../../../../etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(stt.AppNameV1).(stt.Service)
}
