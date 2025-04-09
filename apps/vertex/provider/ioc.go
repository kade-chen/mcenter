package provider

import (
	"github.com/kade-chen/library/ioc"
)

func init() {
	ioc.Config().Registry(&Vertex{})
}

const (
	AppName = "vertex_provider"
)

type Vertex struct {
	ioc.ObjectImpl
	PROJECT_GEMINI_2_0_Flash        string `toml:"project_gemini_2_0_flash" json:"project_gemini_2_0_flash" yaml:"project_gemini_2_0_flash"  env:"project_gemini_2_0_flash"`
	ServiceAccount_GEMINI_2_0_Flash string `toml:"service_account_gemini_2_0_flash" json:"service_account_gemini_2_0_flash" yaml:"service_account_gemini_2_0_flash"  env:"service_account_gemini_2_0_flash"`
}

func (Vertex) Name() string {
	return AppName
}

func (a *Vertex) Init() error {
	return nil
}
