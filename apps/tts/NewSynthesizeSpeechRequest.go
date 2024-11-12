package tts

import "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"

func NewSynthesizeSpeechRequest() *texttospeechpb.SynthesizeSpeechRequest {
	return &texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#synthesisinput
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		//https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#voice
		Voice: &texttospeechpb.VoiceSelectionParams{},
		// Select the type of audio file you want returned.
		// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#audioconfig
		AudioConfig: &texttospeechpb.AudioConfig{},
	}
}

type SynthesizeSpeechRequestBody struct {
	SynthesizeSpeechRequest texttospeechpb.SynthesizeSpeechRequest `json:"SynthesizeSpeechRequest"`
	Text                    string                                 `json:"text"`
	Ssml                    string                                 `json:"ssml"`
	InputBool               bool
}

func NewSynthesizeSpeechRequestBoby() *SynthesizeSpeechRequestBody {
	return &SynthesizeSpeechRequestBody{
		SynthesizeSpeechRequest: texttospeechpb.SynthesizeSpeechRequest{
			// Set the text input to be synthesized.
			// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#synthesisinput
			Input: &texttospeechpb.SynthesisInput{
				// InputSource: &texttospeechpb.SynthesisInput_Text{},
			},
			// Build the voice request, select the language code ("en-US") and the SSML
			// voice gender ("neutral").
			//https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#voice
			//https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages?hl=zh-cn
			Voice: &texttospeechpb.VoiceSelectionParams{},
			// Select the type of audio file you want returned.
			// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#audioconfig
			AudioConfig: &texttospeechpb.AudioConfig{},
		},
	}
}
