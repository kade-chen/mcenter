package api

import (
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/kade-chen/mcenter/apps/stt"
)

// func (h *speechToTextV2Handler) streamingRecognize(r *restful.Request, w *restful.Response) {
// 	// 2.Upgrade to WebSocket
// 	h.log.Info().Msgf("WebSocket upgrade...")
// 	conn, err := upgrader.Upgrade(w.ResponseWriter, r.Request, nil)
// 	if err != nil {
// 		h.log.Error().Msgf("WebSocket upgrade failed: %v", err)
// 		return
// 	}
// 	h.log.Info().Msgf("WebSocket upgrade success")
// 	defer conn.Close()

// 	Speech_StreamingRecognizeClient, _ := h.stt.StreamingRecognize(r.Request.Context(), stt.ENDPOINTS_SPEECH_GOOGLEAPIS_COM)
// 	go func() {
// 		for {
// 			_, message, err := conn.ReadMessage()
// 			if err != nil {
// 				// response.Failed(w, exception.NewInternalServerError("读取消息失败:%s", err))
// 				h.log.Error().Msgf("Read message error: %v", err)
// 				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Read message error: %v", err)))
// 				break
// 			}
// 			trimmedString := strings.Trim(string(message), "¡¿")

// 			reqs := []*speechpb.StreamingRecognizeRequest{{
// 				Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
// 				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
// 					Audio: []byte(trimmedString)},
// 			}}

// 			for _, req := range reqs {
// 				if err := Speech_StreamingRecognizeClient.Send(req); err != nil {
// 					// TODO: Handle error.
// 				}
// 			}
// 			Speech_StreamingRecognizeClient.CloseSend()
// 		}
// 		// log.Println("读取消息成功:", string(message))

// 	}()

// 	// 5. Receive results from Speech API and send to WebSocket
// 	for {
// 		resp, err := Speech_StreamingRecognizeClient.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Printf("Receive message error: %v", err)
// 			// TODO: handle error.
// 		}
// 		// TODO: Use resp.
// 		fmt.Println(resp)
// 	}
// }

func (h *speechToTextV2Handler) streamingRecognize(r *restful.Request, w *restful.Response) {
	// 2.Upgrade to WebSocket
	h.log.Info().Msgf("WebSocket upgrade...")
	conn, err := upgrader.Upgrade(w.ResponseWriter, r.Request, nil)
	if err != nil {
		h.log.Error().Msgf("WebSocket upgrade failed: %v", err)
		return
	}
	h.log.Info().Msgf("WebSocket upgrade success")
	defer conn.Close()

	// 创建 gRPC StreamingRecognize 客户端
	Speech_StreamingRecognizeClient, err := h.stt.StreamingRecognize(r.Request.Context(), stt.ENDPOINTS_SPEECH_GOOGLEAPIS_COM)
	if err != nil {
		h.log.Error().Msgf("Failed to create gRPC stream: %v", err)
		return
	}

	// 启动 goroutine 读取 WebSocket 消息并发送到 gRPC
	go func() {
		defer Speech_StreamingRecognizeClient.CloseSend() // 确保在 goroutine 结束时关闭发送流
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				h.log.Error().Msgf("Read message error: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Read message error: %v", err)))
				break
			}

			trimmedString := strings.Trim(string(message), "¡¿")

			reqs := []*speechpb.StreamingRecognizeRequest{{
				Recognizer: "projects/kade-poc/locations/global/recognizers/long-recognizer1",
				StreamingRequest: &speechpb.StreamingRecognizeRequest_Audio{
					Audio: []byte(trimmedString),
				},
			},
			// TODO: Create requests.
			}
			for _, req := range reqs {
				if err := Speech_StreamingRecognizeClient.Send(req); err != nil {
					// TODO: Handle error.
					h.log.Error().Msgf("Failed to send gRPC message: %v", err)
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
			h.log.Error().Msgf("Failed err: %v", err)
			// fmt.Println(err)
			// TODO: handle error.
		}
		// TODO: Use resp.
		fmt.Println(resp)
	}
}
