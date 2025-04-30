package vertex

import "fmt"

// 授权类型
type GRANT_MODEL uint8

const (
	// bson:"gemini-2.0-flash" json:"gemini-2.0-flash"
	GRANT_MODEL_GEMINI_2_0_Flash_001 GRANT_MODEL = 0
	// bson:"gemini-2.0-pro" json:"gemini-2.0-pro"
	GRANT_MODEL_Gemini_2_0_Pro GRANT_MODEL = 1
	// bson:"gemini-2.0-flash-lite" json:"gemini-2.0-flash-lite"
	GRANT_MODEL_Gemini_2_0_Flash_Lite GRANT_MODEL = 2
	// bson:"gemini-2.0-flash-thinking" json:"gemini-2.0-flash-thinking"
	GRANT_MODEL_Gemini_2_0_Flash_Thinking GRANT_MODEL = 3
	// bson:"gemini-1.5-flash" json:"gemini-1.5-flash"
	GRANT_MODEL_Gemini_1_5_Flash GRANT_MODEL = 4
	// bson:"gemini-1.5-pro" json:"gemini-1.5-pro"
	GRANT_MODEL_Gemini_1_5_Pro GRANT_MODEL = 5
	// bson:"gemini-1.0-pro" json:"gemini-1.0-pro"
	GRANT_MODEL_Gemini_1_0_Pro GRANT_MODEL = 6
	// bson:"gemini-1.0-pro-vision" json:"gemini-1.0-pro-vision"
	GRANT_MODEL_Gemini_1_0_Pro_Vision GRANT_MODEL = 7
)

// Enum value maps for GRANT_MODEL.
var (
	GRANT_MODEL_name = map[GRANT_MODEL]string{
		0: "gemini-2.0-flash-001",
		1: "gemini-2.0-pro-exp-02-05",
		2: "gemini-2.0-flash-lite",
		3: "gemini-2.0-flash-thinking-exp-01-21",
		4: "gemini-1.5-flash",
		5: "gemini-1.5-pro",
		6: "gemini-1.0-pro",
		7: "gemini-1.0-pro-vision",
	}
	GRANT_MODEL_value = map[string]GRANT_MODEL{
		"gemini-2.0-flash-001":                0,
		"gemini-2.0-pro-exp-02-05":            1,
		"gemini-2.0-flash-lite":               2,
		"gemini-2.0-flash-thinking-exp-01-21": 3,
		"gemini-1.5-flash":                    4,
		"gemini-1.5-pro":                      5,
		"gemini-1.0-pro":                      6,
		"gemini-1.0-pro-vision":               7,
	}
)

// 实现 String() 方法
func (g GRANT_MODEL) String() string {
	if name, ok := GRANT_MODEL_name[GRANT_MODEL(g)]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_GRANT_MODEL(%d)", g)
}
