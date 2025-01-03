package api

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/kade-chen/mcenter/apps/user"

	"cloud.google.com/go/speech/apiv1/speechpb"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (h *speechToTextHandler) streamingRecognize(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("CREATE USER , context_user: %s", context_user.Spec.Username)

	// 1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "Speech To Text V1 Admin", "SpeechToTextV1.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	h.log.Info().Msgf("for user %s, check permission...", context_user.Spec.Username)
	_, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	h.log.Info().Msgf("check permission success, context_user: %s", context_user.Spec.Username)

	// 2.Upgrade to WebSocket
	h.log.Info().Msgf("WebSocket upgrade...")
	conn, err := upgrader.Upgrade(w.ResponseWriter, r.Request, nil)
	if err != nil {
		h.log.Error().Msgf("WebSocket upgrade failed: %v", err)
		return
	}
	h.log.Info().Msgf("WebSocket upgrade success")
	defer conn.Close()

	// 3.Send the initial configuration menssage
	Speech_StreamingRecognizeClient, err := h.stt.SpeechToTextStreamingRecognize(r.Request.Context(), &speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 48000,
					LanguageCode:    "cmn-Hans-CN",
					// MaxAlternatives:            1,
					EnableWordConfidence:       true,                               //是否返回单词级置信度信息
					EnableAutomaticPunctuation: true,                               //标点
					EnableSpokenPunctuation:    &wrapperspb.BoolValue{Value: true}, //请求中的相应符号替换口语标点符号。例如，“你好吗问号”变为“你好吗？”  defult true
					EnableSpokenEmojis:         &wrapperspb.BoolValue{Value: true}, //是否替换语音表情符号 defult true
					UseEnhanced:                true,
				},
				InterimResults:            true, // 如果需要中间结果，可以设置为 true
				EnableVoiceActivityEvents: true, //检测到语音活动语音事件时将返回响应。
				// VoiceActivityTimeout: &speechpb.StreamingRecognitionConfig_VoiceActivityTimeout{
				// 	SpeechStartTimeout: &durationpb.Duration{Seconds: 10},
				// 	SpeechEndTimeout:   &durationpb.Duration{Seconds: 10},
				// },
			},
		},
	})
	if err != nil {
		// if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err))); err != nil {
		// 	log.Printf("Error: %v", err)
		// 	return
		// }
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	//4.read message And send to client
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				// response.Failed(w, exception.NewInternalServerError("读取消息失败:%s", err))
				h.log.Error().Msgf("Read message error: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Read message error: %v", err)))
				break
			}
			trimmedString := strings.Trim(string(message), "¡¿")

			if err := Speech_StreamingRecognizeClient.Send(&speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: []byte(trimmedString),
				},
			}); err != nil {
				h.log.Error().Msgf("Send message error: %v", err)
				// response.Failed(w, exception.NewInternalServerError("发送消息失败:%s", err))
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Send message error: %v", err)))
				break
			}
		}
		// log.Println("读取消息成功:", string(message))

	}()

	// 5. Receive results from Speech API and send to WebSocket
	for {
		resp, err := Speech_StreamingRecognizeClient.Recv()
		if err == io.EOF {
			// if err := conn.WriteMessage(websocket.TextMessage, []byte{}); err != nil {
			// 	log.Printf("Stream closed by server error: %v", err)
			// 	return
			// }
			h.log.Error().Msgf("Stream closed by server error: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Stream closed by server error: %v", err)))
			break
		}
		if err != nil {
			// if err := conn.WriteMessage(websocket.TextMessage, []byte{}); err != nil {
			// 	log.Printf("Receiving speech results error: %v", err)
			// 	return
			// }
			h.log.Error().Msgf("Receiving speech results error: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Receiving speech results error: %v", err)))
			break
		}
		if err := resp.Error; err != nil {
			// Workaround while the API doesn't give a more informative error.
			if err.Code == 3 || err.Code == 11 {
				Speech_StreamingRecognizeClient.CloseSend()
				h.log.Error().Msgf("WARNING: Speech recognition request exceeded limit of 60 seconds.%v %v", err.Code, err.Message)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("WARNING2: Speech recognition request exceeded limit of 60 seconds.%v %v", err.Code, err.Message)))
				break
			}
		}

		// 6. Process recognition results
		for _, result := range resp.Results {
			if result.IsFinal {
				h.log.Info().Msgf("最后一次结果：%v", result.Alternatives[0].Transcript)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("最后一次结果：%v", result.Alternatives[0].Transcript)))

			} else {
				if result.Stability > 0.8 {
					h.log.Info().Msgf("大于0.8 stable value: %v, 中间结果大于0.8 stable results: %v", result.Stability, result.Alternatives[0].Transcript)
					conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("中间结果大于0.8的稳定结果：%v", result.Alternatives[0].Transcript)))
				}
			}
		}

	}
}

func (h *speechToTextHandler) localStreamingRecognize(r *restful.Request, w *restful.Response) {
	req := &speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "cmn-Hans-CN",
				},
				InterimResults: true,
			},
		},
	}
	a, err := h.stt.LocalSpeechToTextStreamingRecognize(context.Background(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, fmt.Sprintf("create user success %v", a))
}
