package impl_test

import (
	"context"
	"testing"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/tts"
	_ "github.com/kade-chen/mcenter/apps/tts/v1/impl"
)

var (
	ctx  = context.Background()
	impl tts.Service
)

func TestSimpleTTS(t *testing.T) {
	// req := tts.NewSynthesizeSpeechRequest1("您好")
	req := tts.NewSynthesizeSpeechRequest()
	req.Input.InputSource = &texttospeechpb.SynthesisInput_Text{Text: "您好,音频效果"}
	req.Voice = &texttospeechpb.VoiceSelectionParams{
		//https://cloud.google.com/text-to-speech/docs/voices
		LanguageCode: "cmn-CN", //language code
		// https://cloud.google.com/text-to-speech/docs/voices
		Name:       "cmn-CN-Standard-A", //voice name
		SsmlGender: texttospeechpb.SsmlVoiceGender_MALE,
	}
	req.AudioConfig = &texttospeechpb.AudioConfig{
		// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.AudioEncoding
		AudioEncoding:   texttospeechpb.AudioEncoding_MULAW, //audio encoding
		SpeakingRate:    1.0,                                //speaking rate
		Pitch:           1.0,                                //pitch
		VolumeGainDb:    0.0,                               //volume gain
		SampleRateHertz: 16000,                              //sample rate  //alid values are: 8000-48000. 16000 is optimal. For best results
		// https://cloud.google.com/text-to-speech/docs/audio-profiles
		EffectsProfileId: []string{"telephony-class-application", "handset-class-device"}, //audio effects profile id
	}

	a, err := impl.TextToSpeech(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

func TestOptimizeTTS(t *testing.T) {
	// req := tts.NewSynthesizeSpeechRequest1("您好")
	req := tts.NewSynthesizeSpeechRequest()
	req.Input.InputSource = &texttospeechpb.SynthesisInput_Text{Text: "您好,音频效果"}
	req.Voice = &texttospeechpb.VoiceSelectionParams{
		//https://cloud.google.com/text-to-speech/docs/voices
		LanguageCode: "cmn-TW", //language code
		// https://cloud.google.com/text-to-speech/docs/voices
		Name:       "cmn-TW-Wavenet-A", //voice name
		SsmlGender: texttospeechpb.SsmlVoiceGender_FEMALE,
	}
	req.AudioConfig = &texttospeechpb.AudioConfig{
		// https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.AudioEncoding
		AudioEncoding:   texttospeechpb.AudioEncoding_MULAW, //audio encoding
		SpeakingRate:    1.14,                                //speaking rate
		Pitch:           1.3,                                //pitch
		VolumeGainDb:    +1.5,                               //volume gain
		SampleRateHertz: 48000,                              //sample rate
		// https://cloud.google.com/text-to-speech/docs/audio-profiles
		EffectsProfileId: []string{"telephony-class-application", "handset-class-device"}, //audio effects profile id
	}

	a, err := impl.TextToSpeech(ctx, req)
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
	impl = ioc.Controller().Get(tts.AppNameV1).(tts.Service)
}
