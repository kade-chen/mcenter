package main

import (
	"context"
	"fmt"
	"io"

	"log"
	"net"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"gitee.com/go-kade/library/v2/exception"
	"github.com/kade-chen/mcenter/apps/stt/example/sttv2/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const chunkSize = 15 * 1024 // Each chunk size: 15KB

type Ss struct {
	pb.UnimplementedRPCServer
}

// WebSocket handler
func (Ss) SayHello1(ctx context.Context, AudioData []byte) (string, error) {
	// func (Ss)  SayHello(audioData []byte) string {

	// 加载 Google Cloud 认证
	credentialsFilePath := "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json"
	perRPCCreds, err := oauth.NewServiceAccountFromFile(credentialsFilePath, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	// 创建 gRPC 客户端连接
	conn, err := grpc.Dial(
		"speech.googleapis.com:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(perRPCCreds),
	)
	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	defer conn.Close()

	// 初始化 Speech 客户端
	client := speechpb.NewSpeechClient(conn)
	stream, err := client.StreamingRecognize(context.Background())
	if err != nil {
		log.Fatalf("Failed to create StreamingRecognize client: %v", err)
		return "", err
	}

	go func() {
		defer fmt.Println("cloe")
		defer stream.CloseSend()

		// 按块切分音频数据并发送
		for i := 0; i < len(AudioData); i += chunkSize {
			end := i + chunkSize
			if end > len(AudioData) {
				end = len(AudioData)
			}

			// 当前块
			chunk := AudioData[i:end]

			// 创建 gRPC 请求
			req := &speechpb.StreamingRecognizeRequest{
				Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
					Audio: chunk,
				},
			}

			// 发送请求
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send gRPC message: %v", err)
			}
		}

	}()

	// 接收返回结果
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive gRPC message: %v", err)
		}

		// 解析返回结果
		for _, result := range resp.GetResults() {
			for _, alternative := range result.GetAlternatives() {
				return fmt.Sprintf("Transcript: %s\n", alternative.GetTranscript()), nil
			}
		}
	}
	return "", nil
}

func (s *Ss) SayHello(stream pb.RPC_SayHelloServer) error {
	for {
		msg, err := stream.Recv()
		// fmt.Println(len(msg.AudioData))
		// if err == io.EOF {
		// 	log.Println("Client finished sending messages.")
		// 	return nil
		// }
		if err != nil {
			if stream.Context().Err() == context.Canceled {
				log.Println("Client canceled the request. Likely due to disconnection or timeout.")
				return stream.Context().Err()
			}
			if err == io.EOF {
				log.Println("Client finished sending messages.")
				return nil
			}
		}

		log.Printf("Received audio data of length: %v bytes", len(msg.AudioData))

		// 调用 SayHello1 处理音频
		transcript, err := s.SayHello1(context.Background(), msg.AudioData)
		if err != nil {
			return exception.NewNotFound("Error processing audio: %v", err.Error())
		}
		if err := stream.Send(&pb.Message{Message: transcript}); err != nil {
			return err
		}
		log.Printf("Transcript generated: %s", transcript)
	}
}

func main() {
	Listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRPCServer(s, &Ss{})
	// Start WebSocket server
	if err := s.Serve(Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
