package api

import (
	"fmt"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/tts"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/emicklei/go-restful/v3"
)

func (h *textToSpeechV1Handler) textToSpeech(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("CREATE USER , context_user: %s", context_user.Spec.Username)

	// 1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "Text To Speech V1 Admin", "TextToSpeechV1.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	h.log.Info().Msgf("for user %s, check permission...", context_user.Spec.Username)
	_, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	h.log.Info().Msgf("check permission success, context_user: %s", context_user.Spec.Username)

	//2.read the request body parametars
	// 从前端读取请求
	h.log.Info().Msg("chenk request parametars...")
	req := tts.NewSynthesizeSpeechRequestBoby()
	if err := r.ReadEntity(&req); err != nil {
		h.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}
	h.log.Info().Msg("chenk request parametars success")
	//3.check the request boby parametars
	if req.Text != "" && req.Ssml != "" {
		h.log.Error().Msgf("Error: Only one can be entered for text and ssml")
		response.Failed(w, exception.NewBadRequest("Error: Only one can be entered for text and ssml"))
		return
	} else {
		switch !req.InputBool {
		case req.Text != "":
			req.SynthesizeSpeechRequest.Input.InputSource = &texttospeechpb.SynthesisInput_Text{
				Text: req.Text}
			h.log.Info().Msgf("SynthesizeSpeechRequest.Input.InputSource.Text: %v", req.Text)
		case req.Ssml != "":
			req.SynthesizeSpeechRequest.Input.InputSource = &texttospeechpb.SynthesisInput_Ssml{
				Ssml: req.Ssml}
			h.log.Info().Msgf("SynthesizeSpeechRequest.Input.InputSource.Ssml: %v", req.Ssml)
		default:
			h.log.Error().Msgf("Error: The value cannot be empty for text and ssml")
			response.Failed(w, exception.NewBadRequest("Error: The value cannot be empty for text and ssml"))
			return
		}
	}

	// 4. call the service
	s, err := h.ttsV1.TextToSpeech(r.Request.Context(), &req.SynthesizeSpeechRequest)
	if err != nil {
		h.log.Error().Msgf("err: %s", &req.SynthesizeSpeechRequest)
		response.Failed(w, exception.NewInternalServerError("create user error: %v", err))
		return
	}
	response.Success(w, fmt.Sprintf("create user success %v", s))
}
