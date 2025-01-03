package api

import (
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const chunkSize = 15 * 1024 // 每个分块的大小：15KB
func (h *speechToTextV2Handler) streamingRecognize(r *restful.Request, w *restful.Response) {
	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json")
	// 2.Upgrade to WebSocket
	h.log.Info().Msgf("WebSocket upgrade...")
	conn, err := upgrader.Upgrade(w.ResponseWriter, r.Request, nil)
	if err != nil {
		h.log.Error().Msgf("WebSocket upgrade failed: %v", err)
		return
	}
	h.log.Info().Msgf("WebSocket upgrade success")
	defer conn.Close()
	// 1. 加载服务账号 JSON 文件
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
	// defer conn.Close()
	// 3. 初始化 Speech 客户端
	SpeechClient := speechpb.NewSpeechClient(conn1)

	Speech_StreamingRecognizeClient, err := SpeechClient.StreamingRecognize(r.Request.Context())
	// defer Speech_StreamingRecognizeClient.CloseSend() // 确保在 goroutine 结束时关闭发送流
	// 创建 gRPC StreamingRecognize 客户端
	if err != nil {
		h.log.Error().Msgf("Failed to create gRPC stream: %v", err)
		return
	}
	// 创建缓冲区用于存储每次读取的块
	// 创建缓冲区用于存储每次读取的块
	// buffer := make([]byte, chunkSize) // 设置每块数据大小，chunkSize 为 15KB
	// 启动 goroutine 读取 WebSocket 消息并发送到 gRPC

	go func() {

		defer Speech_StreamingRecognizeClient.CloseSend() // 确保在 goroutine 结束时关闭发送流
		_, message, err := conn.ReadMessage()             // 读取 WebSocket 消息
		fmt.Println(len(message))
		if err != nil {
			// response.Failed(w, exception.NewInternalServerError("读取消息失败:%s", err))
			h.log.Error().Msgf("Read message error: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Read message error: %v", err)))
			// break
		}
		if err == io.EOF {
			fmt.Println("Read message EOF")
			// break // 文件读取完毕
		}

		for i := 0; i < len(message); i += chunkSize {
			end := i + chunkSize
			if end > len(message) {
				end = len(message)
			}
			// 当前块
			chunk := message[i:end]

			// 发送数据块
			req := &speechpb.StreamingRecognizeRequest{
				Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
					Audio: chunk,
				},
			}
			err := Speech_StreamingRecognizeClient.Send(req)

			if err == io.EOF {
				break // 文件读取完毕
			}

			if err != nil {
				// TODO: Handle error.
				h.log.Error().Msgf("Failed to send gRPC message: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Failed to send gRPC message: %v", err)))
				break
			}
		}

	}()

	// 5. Receive results from Speech API and send back to WebSocket
	// 接收返回结果
	for {
		fmt.Println("start recv-----")
		resp, err := Speech_StreamingRecognizeClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive gRPC message: %v", err)
		}

		// 解析返回结果
		for _, result := range resp.GetResults() {
			for _, alternative := range result.GetAlternatives() {
				fmt.Printf("Transcript: %s\n", alternative.GetTranscript())
			}
		}
	}
}
