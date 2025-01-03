package api

import (
	"fmt"
	"net/http"

	"github.com/kade-chen/mcenter/apps/user"

	"cloud.google.com/go/vertexai/genai"
	"github.com/emicklei/go-restful/v3"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/vertex"
	"google.golang.org/api/iterator"
)

// streaming generate content for api
func (g *geminiHandler) streamHandler(r *restful.Request, w *restful.Response) {

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

	g.log.Info().Msg("----------------------------streaming generate content begings----------------------------")

	//1.get gemini setting
	g.log.Info().Msgf("get gemini setting")
	genaiSetting := vertex.NewGenaiSetting()
	// req.ModelName = "gemini-1.5-flash-002"
	if err := r.ReadEntity(&genaiSetting); err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}
	//2.get gemini setting was successful
	// g.log.Info().Msgf("get gemini setting was successful, setting: %v", genaiSetting)
	part := func() []genai.Part {
		prompt := genai.Text(genaiSetting.Text)
		// A := make([]genai.Part, 50)
		if genaiSetting.Url != "" {
			// url := g.getFileDataUrl(genaiSetting.Url)
			parts := g.getFileDataUrl(genaiSetting.Url)
			return append([]genai.Part{prompt}, parts...)
		}
		return []genai.Part{prompt}
	}()
	// g.log.Info().Msgf("get gemini setting was successful, part: %+v", part)
	g.log.Info().Interface("body", part).Msg("get gemini request body,")

	//3.get gemini streaming generate content iterator
	iterators := g.gemini.StreamingGenerateContent(r.Request.Context(), genaiSetting, part...)
	for {
		resp, err := iterators.Next()
		if err == iterator.Done {
			// 结束消息，标识流结束
			g.log.Info().Msg("----------------------------The streaming generate content is complete----------------------------")
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			g.log.Error().Msgf("candidates or parts is empty, error: %v", err)
			response.Failed(w, exception.NewInternalServerError("candidates or parts is empty, error: %v", err))
			return
		}
		if err != nil {
			g.log.Error().Msgf("streaming generate content, error: %v", err)
			response.Failed(w, exception.NewInternalServerError("streaming generate content, error: %v", err))
			return
		}

		// fmt.Fprint(os.Stdout, "generated response: ")
		//4. return the generated content to the client
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				fmt.Fprint(w, p)
				// fmt.Fprintf(os.Stdout, "%v", p) == fmt.Print(p)
				// 立即发送
				if f, ok := w.ResponseWriter.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
		// fmt.Fprint(os.Stdout)
	}

}

// no streaming generate content for api
func (g *geminiHandler) noStreamHandler(r *restful.Request, w *restful.Response) {

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

	g.log.Info().Msg("----------------------------Nostreaming generate content begings----------------------------")

	//1.get gemini setting
	g.log.Info().Msgf("get gemini setting")
	genaiSetting := vertex.NewGenaiSetting()
	// req.ModelName = "gemini-1.5-flash-002"
	if err := r.ReadEntity(&genaiSetting); err != nil {
		g.log.Error().Msgf("check request parametars failed; error: %v", err)
		response.Failed(w, exception.NewInternalServerError("check request parametars failed; error: %v", err))
		return
	}
	//2.get gemini setting was successful
	// g.log.Info().Msgf("get gemini setting was successful, setting: %v", genaiSetting)
	part := func() []genai.Part {
		prompt := genai.Text(genaiSetting.Text)
		// A := make([]genai.Part, 50)
		if genaiSetting.Url != "" {
			// url := g.getFileDataUrl(genaiSetting.Url)
			parts := g.getFileDataUrl(genaiSetting.Url)
			return append([]genai.Part{prompt}, parts...)
		}
		return []genai.Part{prompt}
	}()
	// g.log.Info().Msgf("get gemini setting was successful, part: %+v", part)
	g.log.Info().Interface("body", part).Msg("get gemini request body,")

	//3.get gemini streaming generate content iterator
	iterators, err := g.gemini.NoStreamingGenerateContent(r.Request.Context(), genaiSetting, part...)
	if err != nil {
		g.log.Error().Msgf("streaming generate content, error: %v", err)
		response.Failed(w, exception.NewInternalServerError("streaming generate content, error: %v", err))
		return
	}
	g.log.Info().Msg("----------------------------The Nostreaming generate content is complete----------------------------")
	// rb := format.MustToYaml(iterators)
	response.Success(w, iterators)
}
