package api

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/kade-chen/mcenter/apps/vertex"
)

// streaming generate content for api
func (g *vertexHandler) nostreamHandler(r *restful.Request, w *restful.Response) {
	g.log.Info().Msg("----------------------------The system starts to check user permission----------------------------")
	// 0.0 get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	g.log.Info().Msgf("CREATE USER , context_user: %s", context_user.Spec.Username)

	// 0.1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "Vertex Admin", "Vertex.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	g.log.Info().Msgf("for user %s, check permission...", context_user.Spec.Username)
	_, err := g.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		g.log.Error().Msg("----------------------------The system fails to check the user permissions----------------------------")
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	g.log.Info().Msgf("check permission success, context_user: %s", context_user.Spec.Username)
	g.log.Info().Msg("----------------------------The system passes the user permission check----------------------------")

	gemini2Config := vertex.NewDefaultsGemini2Config()
	if err := r.ReadEntity(&gemini2Config); err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}

	a, err := g.gemini.NoStreamingGenerateContent(r.Request.Context(), gemini2Config)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, a)
}

func (g *vertexHandler) streamHandler(r *restful.Request, w *restful.Response) {

	g.log.Info().Msg("----------------------------The system starts to check user permission----------------------------")
	// 0.0 get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	g.log.Info().Msgf("CREATE USER , context_user: %s", context_user.Spec.Username)

	// 0.1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "Vertex Admin", "Vertex.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	g.log.Info().Msgf("for user %s, check permission...", context_user.Spec.Username)
	_, err := g.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		g.log.Error().Msg("----------------------------The system fails to check the user permissions----------------------------")
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	g.log.Info().Msgf("check permission success, context_user: %s", context_user.Spec.Username)
	g.log.Info().Msg("----------------------------The system passes the user permission check----------------------------")

	gemini2Config := vertex.NewDefaultsGemini2Config()
	if err := r.ReadEntity(&gemini2Config); err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}

	// gemini2Config = vertex.NewGemini2Config()
	// gemini2Config.ModelName = "gemini-2.5-pro-preview-05-06"
	// gemini2Config.Contents = []*genai.Content{
	// 	{
	// 		Parts: []*genai.Part{
	// 			{
	// 				Text: "k8s是什么",
	// 			},
	// 		},
	// 		Role: "user",
	// 	},
	// }

	ass, err := g.gemini.StreamingGenerateContent(r.Request.Context(), gemini2Config)

	if err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, err)
		return
	}

	for a, err := range ass {
		if err != nil {
			g.log.Error().Msgf("check request parametars failed; error: %v", err)
			response.Failed(w, exception.NewInternalServerError("ERROR: %v", err.Error()))
			return
		}
		fmt.Fprint(w, a.Candidates[0].Content.Parts[0].Text)
		// fmt.Fprintf(os.Stdout, "%v", p) == fmt.Print(p)
		// 立即发送
		if f, ok := w.ResponseWriter.(http.Flusher); ok {
			f.Flush()
		}
	}
	// response.Success(w, result)
}
