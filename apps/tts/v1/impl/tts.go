package impl

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kade-chen/library/exception"
)

func (t *textToSpeechV1) TextToSpeech(ctx context.Context, req *texttospeechpb.SynthesizeSpeechRequest) (string, error) {
	t.log.Info().Msgf("text to speech conversion begins...")
	// https://cloud.google.com/text-to-speech/docs/voices
	resp, err := t.client.SynthesizeSpeech(ctx, req)
	if err != nil {
		t.log.Error().Msgf("text to speech conversion failed, error:%v", err)
		return "", exception.NewInternalServerError("text to speech conversion failed, error:%v", err)
	}
	t.log.Info().Msgf("text to speech conversion success")

	// The resp's AudioContent is binary.
	filename := fmt.Sprintf("audio/%s.wav", time.Now().Format("2006-01-02 15:04:05"))
	// filename := "audio/output1.wav"
	t.log.Info().Msgf("check whether the audio directory exists, if it does not exsit, create it")
	err = os.MkdirAll("audio", 0755)
	if err != nil {
		t.log.Error().Msgf("create audio directory error:%v", err)
		return "", exception.NewInternalServerError("create audio directory error:%v", err)
	}
	t.log.Info().Msgf("audio directory exists or created successfully")

	t.log.Info().Msgf("write audio content to file: %v...", filename)
	err = os.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		t.log.Error().Msgf("create user error:%v", err)
		return "", exception.NewInternalServerError("create user error:%v", err)
	}
	t.log.Info().Msgf("write audio content to file: %v success", filename)
	// fmt.Printf("Audio content written to file: %v\n", filename)
	// return time.Now().Format("2006-01-02 15:04") + ".wav"
	return "text to speech v1 client conversion was successful", nil
}
