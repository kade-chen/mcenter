package stt

import (
	"context"

	"cloud.google.com/go/speech/apiv1/speechpb"
	speech "cloud.google.com/go/speech/apiv2/speechpb"
)

const (
	AppNameV1     = "SpeechToTextV1"
	AppNameV1beta = "SpeechToTextV1beta"
	AppNameV2     = "SpeechToTextV2"
)

type Service interface {
	// SpeechToText converts speech to text. <60s and <10MB.
	//Synchronous speech recognition returns the recognized text for short audio (less than 60 seconds).
	//To process a speech recognition request for audio longer than 60 seconds,>60s and >10MB, use SpeechToTextLong.
	SpeechToTextShort(context.Context, *speechpb.RecognizeRequest) (string, error)
	//Asynchronous speech recognition starts a long running audio processing operation.Use asynchronous speech recognition to transcribe audio that is longer than 60 seconds. For shorter audio, synchronous speech recognition is faster and simpler. The upper limit for asynchronous speech recognition is 480 minutes.
	//Audio content can be sent directly to Speech-to-Text from a local file for asynchronous processing. However, the audio time limit for local files is 60 seconds. Attempting to transcribe local audio files that are longer than 60 seconds will result in an error. To use asynchronous speech recognition to transcribe audio longer than 60 seconds, you must have your data saved in a Google Cloud Storage bucket.
	//You can retrieve the results of the operation using the google.longrunning.Operations method. Results remain available for retrieval for 5 days (120 hours). You also have the option of uploading your results directly to a Google Cloud Storage bucket.
	SpeechToTextLong(context.Context, *speechpb.LongRunningRecognizeRequest) (string, error)
	//This section demonstrates how to transcribe streaming audio, like the input from a microphone, to text.
	//Streaming speech recognition allows you to stream audio to Speech-to-Text and receive a stream speech recognition results in real time as the audio is processed.
	//There is a 10 MB limit on all streaming requests sent to the API. This limit applies to to both the initial StreamingRecognize request and the size of each individual message in the stream. Exceeding this limit will throw an error.
	//limit 305s per request
	SpeechToTextStreamingRecognize(context.Context, *speechpb.StreamingRecognizeRequest) (speechpb.Speech_StreamingRecognizeClient, error)

	LocalSpeechToTextStreamingRecognize(ctx context.Context, req *speechpb.StreamingRecognizeRequest) (string, error)
}

type ServiceV2 interface {
	ListLocations(ctx context.Context)
	GetModel(context.Context, ENDPOINT)
	ListRecognizers(context.Context, ENDPOINT) error
	CreateRecognizer(context.Context, ENDPOINT)
	StreamingRecognize(context.Context, ENDPOINT) (speech.Speech_StreamingRecognizeClient, error)
	//get service account
	GetServiceAccount() string
}
