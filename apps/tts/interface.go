package tts

import (
	"context"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

const (
	AppNameV1     = "TextToSpeechV1"
	AppNameV1beta = "TextToSpeechV1beta"
)

type Service interface {
	// TextToSpeech converts text to speech.
	TextToSpeech(context.Context, *texttospeechpb.SynthesizeSpeechRequest) (string, error)
}
