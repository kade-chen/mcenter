package gemini2flash

import (
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
)

var _ provider.Issuer = (*gemini2Flash)(nil)

func init() {
	provider.Registe(&gemini2Flash{})
	fmt.Println("------")
}

type gemini2Flash struct {
}

func (g *gemini2Flash) GrantModel() vertex.GRANT_MODEL {
	// a := vertex.GRANT_MODEL_name[vertex.GRANT_MODEL_GEMINI_2_0_Flash]
	return vertex.GRANT_MODEL_GEMINI_2_0_Flash
}

func (g *gemini2Flash) Init() error {
	return nil
}

func (g *gemini2Flash) ModelIssue(context.Context, *vertex.GenaiSetting, ...genai.Part) *genai.GenerateContentResponseIterator {

	return nil
}
