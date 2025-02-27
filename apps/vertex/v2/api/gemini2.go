package api

import (
	"fmt"
	"net/http"
	"github.com/kade-chen/mcenter/apps/user"

	"github.com/emicklei/go-restful/v3"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/vertex"
)

func (g *gemini2Handler) stream(r *restful.Request, w *restful.Response) {

	gemini2Config := vertex.NewDefaultsGemini2Config()
	if err := r.ReadEntity(&gemini2Config); err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}

	iter1 := g.gemini2.StreamingGenerateContent(r.Request.Context(), gemini2Config)
	for a, err := range iter1 {
		if err != nil {
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
}

func (g *gemini2Handler) noStream(r *restful.Request, w *restful.Response) {

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

	resp, err := g.gemini2.NoStreamingGenerateContent(r.Request.Context(), gemini2Config)
	if err != nil {
		// TODO: Handle error.
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}
	// TODO: Use resp.
	response.Success(w, resp.Candidates[0].Content.Parts[0].Text)
}

// for a, err := range iter1 {
// 	if err != nil {
// 		response.Failed(w, exception.NewInternalServerError("ERROR: %v", err.Error()))
// 		return
// 	}
// 	fmt.Fprint(w, a.Candidates[0].Content.Parts[0].Text)
// 	// fmt.Fprintf(os.Stdout, "%v", p) == fmt.Print(p)
// 	// 立即发送
// 	if f, ok := w.ResponseWriter.(http.Flusher); ok {
// 		f.Flush()
// 	}
// }
