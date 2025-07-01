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
	// Default_Service_Account_Bool bool   `toml:"default_service_account_bool" json:"default_service_account_bool" yaml:"default_service_account_bool"`
	Default_Service_Account_Name string `toml:"default_service_account_name" json:"default_service_account_name" yaml:"default_service_account_name"`
	Default_Project_ID_Bool      bool   `toml:"default_project_id_bool" json:"default_project_id_bool" yaml:"default_project_id_bool"`
	Default_Project_ID           string `toml:"default_project_id" json:"default_project_id" yaml:"default_project_id"`

	Gemini_25_Flash_Preview_Project_ID_Bool      bool   `toml:"gemini_25_flash_preview_project_id_bool" json:"gemini_25_flash_preview_project_id_bool" yaml:"gemini_25_flash_preview_project_id_bool`
	Gemini_25_Flash_Preview_Project_ID           string `toml:"gemini_25_flash_preview_project_id" json:"gemini_25_flash_preview_project_id" yaml:"gemini_25_flash_preview_project_id`
	// Gemini_25_Flash_Preview_Service_Account_Bool bool   `toml:"gemini_25_flash_preview_service_account_bool" json:"gemini_25_flash_preview_service_account_bool" yaml:"gemini_25_flash_preview_service_account_bool`
	Gemini_25_Flash_Preview_Service_Account_Name string `toml:"gemini_25_flash_preview_service_account_name" json:"gemini_25_flash_preview_service_account_name" yaml:"gemini_25_flash_preview_service_account_name`

	PROJECT_GEMINI_2_0_Flash                string `toml:"project_gemini_2_0_flash" json:"project_gemini_2_0_flash" yaml:"project_gemini_2_0_flash"  env:"project_gemini_2_0_flash"`
	ServiceAccount_GEMINI_2_0_Flash         string `toml:"service_account_gemini_2_0_flash" json:"service_account_gemini_2_0_flash" yaml:"service_account_gemini_2_0_flash"  env:"service_account_gemini_2_0_flash"`
	PROJECT_GEMINI_2_5_Flash                string `toml:"project_gemini_2_5_flash" json:"project_gemini_2_5_flash" yaml:"project_gemini_2_5_flash"  env:"project_gemini_2_5_flash`
	ServiceAccount_GEMINI_2_5_Flash_Preview string `toml:"service_account_gemini_2_5_flash_preview" json:"service_account_gemini_2_5_flash_preview" yaml:"service_account_gemini_2_5_flash_preview"  env:"service_account_gemini_2_5_flash_preview"`
}

func (Vertex) Name() string {
	return AppName
}

func (a *Vertex) Init() error {
	return nil
}
