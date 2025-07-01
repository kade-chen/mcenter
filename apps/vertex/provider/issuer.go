package provider

import (
	"context"
	"fmt"
	"iter"

	"github.com/kade-chen/mcenter/apps/vertex"
	"google.golang.org/genai"
)

var (
	m = make(map[vertex.GRANT_MODEL]Issuer)
)

type Issuer interface {
	Init() error
	GrantModel() vertex.GRANT_MODEL
	ModelIssuer
}

type ModelIssuer interface {
	NoStreamingGenerateContent(context.Context, *vertex.Gemini_Config) (*genai.GenerateContentResponse, error)
	StreamingGenerateContent(context.Context, *vertex.Gemini_Config) iter.Seq2[*genai.GenerateContentResponse, error]
}

func Registe(i Issuer) {
	m[i.GrantModel()] = i
}

func ProviderRegistry() (cc []vertex.GRANT_MODEL, err error) {
	for k, _ := range m {
		cc = append(cc, k)
	}

	if len(cc) == 0 {
		return nil, fmt.Errorf("no issuer")
	}

	return cc, nil
}

func CheckProviderRegistry(a vertex.GRANT_MODEL, b []vertex.GRANT_MODEL) bool {
	for _, v := range b {
		if v == a {
			return true
		}
	}
	return false
}

func Init() error {
	for k, v := range m {
		if err := v.Init(); err != nil {
			return fmt.Errorf("init %s issuer error", k)
		}
	}
	return error(nil)
}

func GetModelIssuer(gt vertex.GRANT_MODEL) ModelIssuer {
	if v, ok := m[gt]; ok {
		// v.Init()
		return v
	}
	return nil
}
