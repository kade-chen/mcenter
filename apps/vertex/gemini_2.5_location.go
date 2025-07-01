package vertex

type Gemini_25_FLASH_PREVIEW_LOCATION uint8

// https://cloud.google.com/vertex-ai/generative-ai/docs/learn/locations#google_model_endpoint_locations
const (
	//United States
	//Iowa  爱荷华州
	Gemini_25_FLASH_PREVIEW_LOCATION_US_Central1 Gemini_25_FLASH_PREVIEW_LOCATION = iota

	// //Columbus, Ohio 俄亥俄州-哥伦布市
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East5 Gemini_25_FLASH_PREVIEW_LOCATION = iota
	// //Dallas, Texas 德克萨斯州-达拉斯
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_South1
	// //Iowa  爱荷华州
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_Central1
	// //Las Vegas, Nevada 内华达州-拉斯维加斯
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_West4
	// //Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East1
	// //Northern Virginia 北弗吉尼亚
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East4
	// //Oregon 俄勒冈州
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_West1

	// //Canada
	// //Montréal 	蒙特利尔
	// Gemini_25_FLASH_PREVIEW_LOCATION_NorthAmerica_Northeast1

	// //South America
	// //São Paulo, Brazil 巴西圣保罗
	// Gemini_25_FLASH_PREVIEW_LOCATION_SouthAmerica_East1

	// //Europe
	// //Netherlands 荷兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West4
	// //Paris, France 法国巴黎
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West9
	// //London, United Kingdom 英国伦敦
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West2
	// //Frankfurt, Germany 德国法兰克福
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West3
	// // Belgium 比利时
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West1
	// //Zürich, Switzerland 瑞士苏黎世
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West6
	// //Madrid, Spain 西班牙马德里
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Southwest1
	// //Milan, Italy 意大利米兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West8
	// //Finland 芬兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_North1
	// //Warsaw, Poland 波兰华沙
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Central2

	// //Asia Pacific
	// //Tokyo, Japan 	日本-东京
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast1
	// //Sydney, Australia 澳大利亚-悉尼
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Australia_Southeast1
	// //Singapore 新加坡
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Southeast1
	// //Seoul, Korea 韩国-首尔
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast3
	// //Taiwan 台湾
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East1
	// //Hong Kong, China 中国-香港
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East2
	// //Mumbai, India 印度-孟买
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_South1

	// //Middle East
	// //Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central2
	// //Doha, Qatar 卡塔尔-多哈
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central1
	// //Tel Aviv, Israel 以色列-特拉维夫
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_West1

	Gemini_25_FLASH_PREVIEW_LOCATION_Global
)

// 创建双向映射
var Gemini_25_FLASH_PREVIEW_LOCATION_To_String = map[Gemini_25_FLASH_PREVIEW_LOCATION]string{
	// United States
	// //Columbus, Ohio 俄亥俄州-哥伦布市
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East5: "us-east5",
	// //Dallas, Texas 德克萨斯州-达拉斯
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_South1: "us-south1",
	// //Iowa  爱荷华州
	Gemini_25_FLASH_PREVIEW_LOCATION_US_Central1: "us-central1",
	// //Las Vegas, Nevada 内华达州-拉斯维加斯
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_West4: "us-west4",
	// //Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East1: "us-east1",
	// //Northern Virginia 北弗吉尼亚
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_East4: "us-east4",
	// //Oregon 俄勒冈州
	// Gemini_25_FLASH_PREVIEW_LOCATION_US_West1: "us-west1",

	// // Canada
	// //Montréal 	蒙特利尔
	// Gemini_25_FLASH_PREVIEW_LOCATION_NorthAmerica_Northeast1: "northamerica-northeast1",

	// // South America
	// //São Paulo, Brazil 巴西圣保罗
	// Gemini_25_FLASH_PREVIEW_LOCATION_SouthAmerica_East1: "southamerica-east1",

	// // Europe
	// //Netherlands 荷兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West4: "europe-west4",
	// //Paris, France 法国巴黎
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West9: "europe-west9",
	// //London, United Kingdom 英国伦敦
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West2: "europe-west2",
	// //Frankfurt, Germany 德国法兰克福
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West3: "europe-west3",
	// // Belgium 比利时
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West1: "europe-west1",
	// //Zürich, Switzerland 瑞士苏黎世
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West6: "europe-west6",
	// //Madrid, Spain 西班牙马德里
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Southwest1: "europe-southwest1",
	// //Milan, Italy 意大利米兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West8: "europe-west8",
	// //Finland 芬兰
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_North1: "europe-north1",
	// //Warsaw, Poland 波兰华沙
	// Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Central2: "europe-central2",

	// // Asia Pacific
	// //Tokyo, Japan 	日本-东京
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast1: "asia-northeast1",
	// //Sydney, Australia 澳大利亚-悉尼
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Australia_Southeast1: "australia-southeast1",
	// //Singapore 新加坡
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Southeast1: "asia-southeast1",
	// //Seoul, Korea 韩国-首尔
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast3: "asia-northeast3",
	// //Taiwan 台湾
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East1: "asia-east1",
	// //Hong Kong, China 中国-香港
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East2: "asia-east2",
	// //Mumbai, India 印度-孟买
	// Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_South1: "asia-south1",

	// // Middle East
	// //Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central2: "me-central2",
	// //Doha, Qatar 卡塔尔-多哈
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central1: "me-central1",
	// //Tel Aviv, Israel 以色列-特拉维夫
	// Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_West1: "me-west1",

	Gemini_25_FLASH_PREVIEW_LOCATION_Global: "global",
}

var Gemini_25_FLASH_PREVIEW_String_To_LOCATION = map[string]Gemini_25_FLASH_PREVIEW_LOCATION{
	// // United States
	// //Columbus, Ohio 俄亥俄州-哥伦布市
	// "us-east5": Gemini_25_FLASH_PREVIEW_LOCATION_US_East5,
	// //Dallas, Texas 德克萨斯州-达拉斯
	// "us-south1": Gemini_25_FLASH_PREVIEW_LOCATION_US_South1,
	//Iowa  爱荷华州
	"us-central1": Gemini_25_FLASH_PREVIEW_LOCATION_US_Central1,
	// //Las Vegas, Nevada 内华达州-拉斯维加斯
	// "us-west4": Gemini_25_FLASH_PREVIEW_LOCATION_US_West4,
	// //Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	// "us-east1": Gemini_25_FLASH_PREVIEW_LOCATION_US_East1,
	// //Northern Virginia 北弗吉尼亚
	// "us-east4": Gemini_25_FLASH_PREVIEW_LOCATION_US_East4,
	// //Oregon 俄勒冈州
	// "us-west1": Gemini_25_FLASH_PREVIEW_LOCATION_US_West1,

	// // Canada
	// //Montréal 	蒙特利尔
	// "northamerica-northeast1": Gemini_25_FLASH_PREVIEW_LOCATION_NorthAmerica_Northeast1,

	// // South America
	// //São Paulo, Brazil 巴西圣保罗
	// "southamerica-east1": Gemini_25_FLASH_PREVIEW_LOCATION_SouthAmerica_East1,

	// // Europe
	// //Netherlands 荷兰
	// "europe-west4": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West4,
	// //Paris, France 法国巴黎
	// "europe-west9": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West9,
	// //London, United Kingdom 英国伦敦
	// "europe-west2": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West2,
	// //Frankfurt, Germany 德国法兰克福
	// "europe-west3": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West3,
	// // Belgium 比利时
	// "europe-west1": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West1,
	// //Zürich, Switzerland 瑞士苏黎世
	// "europe-west6": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West6,
	// //Madrid, Spain 西班牙马德里
	// "europe-southwest1": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Southwest1,
	// //Milan, Italy 意大利米兰
	// "europe-west8": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_West8,
	// //Finland 芬兰
	// "europe-north1": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_North1,
	// //Warsaw, Poland 波兰华沙
	// "europe-central2": Gemini_25_FLASH_PREVIEW_LOCATION_Europe_Central2,

	// // Asia Pacific
	// //Tokyo, Japan 	日本-东京
	// "asia-northeast1": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast1,
	// //Sydney, Australia 澳大利亚-悉尼
	// "australia-southeast1": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Australia_Southeast1,
	// //Singapore 新加坡
	// "asia-southeast1": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Southeast1,
	// //Seoul, Korea 韩国-首尔
	// "asia-northeast3": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_Northeast3,
	// //Taiwan 台湾
	// "asia-east1": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East1,
	// //Hong Kong, China 中国-香港
	// "asia-east2": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_East2,
	// //Mumbai, India 印度-孟买
	// "asia-south1": Gemini_25_FLASH_PREVIEW_LOCATION_AsiaPacific_South1,

	// // Middle East
	// //Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// "me-central2": Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central2,
	// //Doha, Qatar 卡塔尔-多哈
	// "me-central1": Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_Central1,
	// //Tel Aviv, Israel 以色列-特拉维夫
	// "me-west1": Gemini_25_FLASH_PREVIEW_LOCATION_MiddleEast_West1,

	"global": Gemini_25_FLASH_PREVIEW_LOCATION_Global,
}

// 获取 LOCATION 对应的字符串
func Get_Gemini_25_FLASH_PREVIEW_LOCATION_To_String(location Gemini_25_FLASH_PREVIEW_LOCATION) string {
	return Gemini_25_FLASH_PREVIEW_LOCATION_To_String[location]
}

// 获取字符串对应的 LOCATION
func Get_Gemini_25_FLASH_PREVIEW_String_To_LOCATION(locationStr string) (Gemini_25_FLASH_PREVIEW_LOCATION, bool) {
	location, exists := Gemini_25_FLASH_PREVIEW_String_To_LOCATION[locationStr]
	return location, exists
}
