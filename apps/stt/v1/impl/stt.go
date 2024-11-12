package impl

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"github.com/kade-chen/library/exception"
	"google.golang.org/api/option"
)

// SpeechToText converts speech to text. <60s
// Synchronous speech recognition returns the recognized text for short audio (less than 60 seconds).
// To process a speech recognition request for audio longer than 60 seconds, use SpeechToTextLong.
func (s *speechToText) SpeechToTextShort(ctx context.Context, req *speechpb.RecognizeRequest) (string, error) {
	resp, err := s.client.Recognize(ctx, req)

	if err != nil {
		return "", fmt.Errorf("failed to recognize speech: %v", err)
	}
	// 提取转录结果
	// fmt.Println("----------2", resp.Results)
	var transcription strings.Builder
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			// fmt.Println("----------2", alt.Words)
			transcription.WriteString(alt.Transcript)
		}
	}

	return transcription.String(), nil
}

// Asynchronous speech recognition starts a long running audio processing operation.Use asynchronous speech recognition to transcribe audio that is longer than 60 seconds. For shorter audio, synchronous speech recognition is faster and simpler. The upper limit for asynchronous speech recognition is 480 minutes.
// Audio content can be sent directly to Speech-to-Text from a local file for asynchronous processing. However, the audio time limit for local files is 60 seconds. Attempting to transcribe local audio files that are longer than 60 seconds will result in an error. To use asynchronous speech recognition to transcribe audio longer than 60 seconds, you must have your data saved in a Google Cloud Storage bucket.
// You can retrieve the results of the operation using the google.longrunning.Operations method. Results remain available for retrieval for 5 days (120 hours). You also have the option of uploading your results directly to a Google Cloud Storage bucket.
func (s *speechToText) SpeechToTextLong(ctx context.Context, req *speechpb.LongRunningRecognizeRequest) (string, error) {
	longRunningRecognizeOperation, err := s.client.LongRunningRecognize(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to start long running recognize: %v", err)
	}
	resp, err := longRunningRecognizeOperation.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to start long running recognize: %v", err)
	}
	// Print the results.
	var transcription strings.Builder
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			transcription.WriteString(alt.Transcript)
		}
	}
	return transcription.String(), nil
}

// Speech To Text Streaming Recognize
func (s *speechToText) SpeechToTextStreamingRecognize(ctx context.Context, req *speechpb.StreamingRecognizeRequest) (speechpb.Speech_StreamingRecognizeClient, error) {

	//StreamingRecognize performs bidirectional streaming speech recognition: receive results while sending audio. This method is only available via the gRPC API (not REST).
	// This method is not supported for the REST transport.
	Speech_StreamingRecognizeClient, err := s.client.StreamingRecognize(ctx)
	if err != nil {
		s.log.Error().Msgf("Streaming Recoginze init client error: %v", err)
		return Speech_StreamingRecognizeClient, exception.NewInternalServerError("Streaming Recoginze init client error: %v", err)
	}

	s.log.Info().Msgf("Streaming Recoginze send configuration initializing...")
	err = Speech_StreamingRecognizeClient.Send(req)
	if err != nil {
		s.log.Error().Msgf("Streaming Recoginze send configuration error: %v", err)
		return Speech_StreamingRecognizeClient, exception.NewInternalServerError("Streaming Recoginze send configuration error: %v", err)
	}

	s.log.Info().Msgf("Streaming Recognize send configuration success")
	return Speech_StreamingRecognizeClient, err
	/*
	   // 发送音频数据到 Google Cloud

	   	go func() {
	   		for {
	   			_, message, err := conn.ReadMessage()
	   			if err != nil {
	   				log.Println("读取消息失败:", err)
	   				break
	   			}
	   			trimmedString := strings.Trim(string(message), "¡¿")

	   			if err := Speech_StreamingRecognizeClient.Send(&speechpb.StreamingRecognizeRequest{
	   				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
	   					AudioContent: []byte(trimmedString),
	   				},
	   			}); err != nil {
	   				log.Println("发送消息失败:", err)

	   			}
	   			// log.Println("发送消息message:", []byte(trimmedString))
	   		}
	   		// log.Println("读取消息成功:", string(message))

	   }()

	   接收识别结果并返回给客户端

	   	for {
	   		// log.Println("Waiting for response...")
	   		resp, err := Speech_StreamingRecognizeClient.Recv()
	   		// if err == io.EOF {
	   		//    break[]
	   		// }
	   		if err != nil {
	   			log.Fatalf("Cannot stream results: %v", err)
	   		}
	   		if err := resp.Error; err != nil {
	   			// Workaround while the API doesn't give a more informative error.
	   			if err.Code == 3 || err.Code == 11 {
	   				log.Printf("WARNING: Speech recognition request exceeded limit of 60 seconds.%v %v", err.Code, err.Message)
	   			}
	   			log.Fatalf("Could not recognize: %v", err)
	   		}
	   		// log.Println("Received response.")
	   		for _, result := range resp.Results {
	   			log.Printf("Result: %+v\n", result)
	   			// return fmt.Sprintf("Result: %+v\n", result), nil
	   		}
	   	}
	*/
}

// This section demonstrates how to transcribe streaming audio, like the input from a microphone, to text.
// Streaming speech recognition allows you to stream audio to Speech-to-Text and receive a stream speech recognition results in real time as the audio is processed.
// There is a 10 MB limit on all streaming requests sent to the API. This limit applies to to both the initial StreamingRecognize request and the size of each individual message in the stream. Exceeding this limit will throw an error.
func (s *speechToText) LocalSpeechToTextStreamingRecognize(ctx context.Context, req *speechpb.StreamingRecognizeRequest) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	Speech_StreamingRecognizeClient, err := s.client.StreamingRecognize(ctx)
	if err != nil {
		return "", exception.NewInternalServerError("error: %v", err)
	}
	defer Speech_StreamingRecognizeClient.CloseSend()

	err = Speech_StreamingRecognizeClient.Send(req)
	if err != nil {
		return "", exception.NewInternalServerError("error: %v", err)
	}
	fmt.Println("-----------1")
	// 使用 PortAudio 获取麦克风音频流
	// 创建一个带缓冲的通道
	streamingAudio := make(chan []byte, 1024)

	// 启动音频捕获的goroutine
	// go captureAudio(streamingAudio)
	// go captureAudio(streamingAudio)
	fmt.Println("-----------")
	// 发送音频数据到 Google Cloud
	go func() {
		for audioData := range streamingAudio {
			if err := Speech_StreamingRecognizeClient.Send(&speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: audioData,
				},
			}); err != nil {
				exception.NewInternalServerError("error: %v", err)
			}
		}
	}()
	fmt.Println("----------2-")
	// 接收识别结果并返回给客户端
	for {
		fmt.Println("---------400-")
		resp, err := Speech_StreamingRecognizeClient.Recv()
		fmt.Println("---------4010-")
		if err == io.EOF {
			fmt.Println("---------41-")
			break
		}
		if err != nil {
			fmt.Println("---------42-")
			exception.NewInternalServerError("error: %v", err)
		}
		for _, result := range resp.Results {
			fmt.Println("---------43-")
			if result.IsFinal {
				fmt.Println("---------44-")
				fmt.Fprintf(io.MultiWriter(), "Final Result: %s\n", result.Alternatives[0].Transcript)
			} else {
				fmt.Println("---------45-")
				fmt.Printf("Interim Result: %s\n", result.Alternatives[0].Transcript)
			}
		}
	}
	fmt.Println("---------4-")
	return "", nil
}

// This method is converted directly to main.go
// and then run sox-d-r 16000 -e signed-integer -b 16 -c 1 -t raw - | go run main.go to translate it locally in real time
func (s *speechToText) soxSpeechToTextStreamingRecognize() {
	ctx := context.Background()
	client, err := speech.NewClient(ctx, option.WithCredentialsFile("/Users/kade.chen/Downloads/kade-poc-0c88b6757f62.json"))
	if err != nil {
		log.Fatal(err)
	}
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "cmn-Hans-CN",
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	go func() {
		// Pipe stdin to the API.
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				if err := stream.Send(&speechpb.StreamingRecognizeRequest{
					StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
						AudioContent: buf[:n],
					},
				}); err != nil {
					log.Printf("Could not send audio: %v", err)
				}
				// fmt.Println("--------------12010210", buf[:n])
			}
			if err == io.EOF {
				// Nothing else to pipe, close the stream.
				if err := stream.CloseSend(); err != nil {
					log.Fatalf("Could not close stream: %v", err)
				}
				return
			}
			if err != nil {
				log.Printf("Could not read from stdin: %v", err)
				continue
			}
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cannot stream results: %v", err)
		}
		if err := resp.Error; err != nil {
			// Workaround while the API doesn't give a more informative error.
			if err.Code == 3 || err.Code == 11 {
				log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
			}
			log.Fatalf("Could not recognize: %v", err)
		}
		for _, result := range resp.Results {
			fmt.Printf("Result: %+v\n", result)
		}
	}
}
