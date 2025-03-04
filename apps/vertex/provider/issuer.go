package provider

import (
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/mcenter/apps/vertex"
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
	ModelIssue(context.Context, *vertex.GenaiSetting, ...genai.Part) *genai.GenerateContentResponseIterator
}

func Registe(i Issuer) {
	m[i.GrantModel()] = i
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
