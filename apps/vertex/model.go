package vertex

import "fmt"

// 授权类型
type GRANT_MODEL uint8

const (
	// bson:"gemini-2.0-flash" json:"gemini-2.0-flash"
	GRANT_MODEL_GEMINI_2_0_Flash_001 GRANT_MODEL = iota
	// bson:"gemini-2.0-pro" json:"gemini-2.0-pro"
	GRANT_MODEL_Gemini_2_0_Pro
	// bson:"gemini-2.0-flash-lite" json:"gemini-2.0-flash-lite"
	GRANT_MODEL_Gemini_2_0_Flash_Lite
	// bson:"gemini-2.0-flash-thinking" json:"gemini-2.0-flash-thinking"
	GRANT_MODEL_Gemini_2_0_Flash_Thinking
	// bson:"gemini-2.5-flash-preview-05-20" json:"gemini-2.5-flash-preview-05-20"
	GRANT_MODEL_Gemini_2_5_Flash_Preview
)

// Enum value maps for GRANT_MODEL.
var (
	GRANT_MODEL_name = map[GRANT_MODEL]string{
		0: "gemini-2.0-flash",
		1: "gemini-2.0-pro-exp-02-05",
		2: "gemini-2.0-flash-lite",
		3: "gemini-2.0-flash-thinking-exp-01-21",
		4: "gemini-2.5-flash-preview-05-20",
	}
	GRANT_MODEL_value = map[string]GRANT_MODEL{
		"gemini-2.0-flash":                    0,
		"gemini-2.0-pro-exp-02-05":            1,
		"gemini-2.0-flash-lite":               2,
		"gemini-2.0-flash-thinking-exp-01-21": 3,
		"gemini-2.5-flash-preview-05-20":      4,
	}
)

// 实现 String() 方法
func (g GRANT_MODEL) String() string {
	if name, ok := GRANT_MODEL_name[GRANT_MODEL(g)]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_GRANT_MODEL(%d)", g)
}
