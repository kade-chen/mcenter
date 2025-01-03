package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const chunkSize = 15 * 1024 // 每个分块的大小：15KB

func main() {
	credentialsFilePath := "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json"
	perRPCCreds, err := oauth.NewServiceAccountFromFile(credentialsFilePath, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		fmt.Printf("Failed to create credentials: %v\n", err)
		return
	}

	// 2. 创建 gRPC 客户端连接
	conn1, err := grpc.Dial(
		"speech.googleapis.com:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(perRPCCreds),
	)
	if err != nil {
		fmt.Printf("Failed to create gRPC connection: %v\n", err)
		return
	}
	// 3. 初始化 Speech 客户端
	SpeechClient := speechpb.NewSpeechClient(conn1)

	Speech_StreamingRecognizeClient, err := SpeechClient.StreamingRecognize(context.Background())
	// 创建 gRPC StreamingRecognize 客户端
	if err != nil {
		fmt.Printf("Failed to create StreamingRecognize client: %v\n")
		return
	}
	// 读取音频文件并分段发送
	file1, err := os.Open("/Users/kade.chen/Downloads/3.wav") // 替换为你的音频文件路径
	if err != nil {
		log.Fatalf("Failed to open audio file: %v", err)
		return
	}

	defer file1.Close()
	// 逐块读取音频文件并发送到 gRPC
	buffer := make([]byte, chunkSize) // 用于存储每块音频数据

	// 启动 goroutine 读取 WebSocket 消息并发送到 gRPC
	go func() {
		defer Speech_StreamingRecognizeClient.CloseSend() // 确保在 goroutine 结束时关闭发送流
		for {
			n, err := file1.Read(buffer)
			// fmt.Println(n)
			if err == io.EOF {
				break // 文件读取完毕
			}
			// trimmedString := strings.Trim(string(message), "¡¿")
			reqs := []*speechpb.StreamingRecognizeRequest{{
				Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
					Audio: []byte(buffer[:n]),
				},
			},
			// TODO: Create requests.
			}
			for _, req := range reqs {
				if err := Speech_StreamingRecognizeClient.Send(req); err != nil {
					// TODO: Handle error.
					fmt.Printf("Failed to send gRPC message: %v")
				}
			}
		}
	}()

	// 5. Receive results from Speech API and send back to WebSocket
	for {
		resp, err := Speech_StreamingRecognizeClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to receive gRPC message: %v")
			break
			// fmt.Println(err)
			// TODO: handle error.
		}
		// TODO: Use resp.
		fmt.Println(resp)
	}

}
