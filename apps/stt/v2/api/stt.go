package api

import (
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
)

const chunkSize = 15 * 1024 // 每个分块的大小：15KB

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 在这里可以根据请求的 Origin 进行判断，返回 true 表示允许连接
		// 返回 false 表示拒绝连接
		return true // 允许所有请求
	},
}

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

	// 3. 初始化 Speech 客户端
	SpeechClient := speechpb.NewSpeechClient(h.grpc)

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
		defer h.log.Info().Msg("Closing gRPC stream")

		_, message, err := conn.ReadMessage() // 读取 WebSocket 消息
		h.log.Info().Msgf("Received message length: %d", len(message))

		if err != nil {
			// response.Failed(w, exception.NewInternalServerError("读取消息失败:%s", err))
			h.log.Error().Msgf("Read message error: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Read message error: %v", err)))
			return
		}
		if err == io.EOF {
			h.log.Info().Msg("Send message EOF")
			return // 文件读取完毕
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
				h.log.Info().Msg("Send message EOF")
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
		h.log.Info().Msg("Received message...")
		resp, err := Speech_StreamingRecognizeClient.Recv()
		if err == io.EOF {
			h.log.Info().Msg("Received message EOF")
			break
		}
		if err != nil {
			h.log.Error().Msgf("Failed to receive gRPC message: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Failed to receive gRPC message: %v", err)))
			break
		}

		// 解析返回结果
		for _, result := range resp.Results {

			go func(SpeechRecognitionAlternative []*speechpb.SpeechRecognitionAlternative) {
				for _, alternative := range SpeechRecognitionAlternative {
					h.log.Info().Msgf("识别结果列表：%v %f", alternative.Transcript, alternative.Confidence)
				}
			}(result.Alternatives)

			if result.IsFinal {
				h.log.Info().Msgf("最后一次结果：%v %f", result.Alternatives[0].Transcript, result.Alternatives[0].Confidence)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("最后一次结果：%v", result.Alternatives[0].Transcript)))
				break
			} else {
				if result.Stability > 0.8 {
					h.log.Info().Msgf("大于0.8 stable value: %v, 中间结果大于0.8 stable results: %v", result.Stability, result.Alternatives[0].Transcript)
					conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("中间结果大于0.8的稳定结果：%v", result.Alternatives[0].Transcript)))
					break
				}
			}
		}
	}
}
