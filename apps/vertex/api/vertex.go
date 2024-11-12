package api

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/emicklei/go-restful/v3"
	"google.golang.org/api/iterator"
)

func (g *geminiHandler) streamHandler(r *restful.Request, w *restful.Response) {
	g.log.Info().Msg("----------------------------streaming generate content begings----------------------------")
	defer func() {
		if err := recover(); err != nil {
			g.log.Error().Msgf("Internal Server Error; error: %v", err)
			response.Failed(w, exception.NewInternalServerError("Internal Server Error; error: %v", err))
			return
		}
	}()

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

		if genaiSetting.Url != "" {
			url := g.getFileDataUrl(genaiSetting.Url)
			return []genai.Part{url, prompt}
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
			g.log.Info().Msg(`{"status":"completed","message":"Stream completed successfully"}`)
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
				// 立即发送
				if f, ok := w.ResponseWriter.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
		// fmt.Fprint(os.Stdout, "\n")
	}

}
