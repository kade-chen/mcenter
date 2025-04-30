package vertex

type GEMINI_2_LOCATION uint8

// https://cloud.google.com/vertex-ai/generative-ai/docs/learn/locations#middle-east
const (
	//United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	GEMINI2_LOCATION_US_East5 GEMINI_2_LOCATION = iota
	//Dallas, Texas 德克萨斯州-达拉斯
	GEMINI2_LOCATION_US_South1
	//Iowa  爱荷华州
	GEMINI2_LOCATION_US_Central1
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	GEMINI2_LOCATION_US_West4
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	GEMINI2_LOCATION_US_East1
	//Northern Virginia 北弗吉尼亚
	GEMINI2_LOCATION_US_East4
	//Oregon 俄勒冈州
	GEMINI2_LOCATION_US_West1

	//Canada
	//Montréal 	蒙特利尔
	// GEMINI2_LOCATION_NorthAmerica_Northeast1

	//South America
	//São Paulo, Brazil 巴西圣保罗
	// GEMINI2_LOCATION_SouthAmerica_East1

	//Europe
	//Netherlands 荷兰
	GEMINI2_LOCATION_Europe_West4
	//Paris, France 法国巴黎
	GEMINI2_LOCATION_Europe_West9
	//London, United Kingdom 英国伦敦
	// GEMINI2_LOCATION_Europe_West2
	//Frankfurt, Germany 德国法兰克福
	// GEMINI2_LOCATION_Europe_West3
	// Belgium 比利时
	GEMINI2_LOCATION_Europe_West1
	//Zürich, Switzerland 瑞士苏黎世
	// GEMINI2_LOCATION_Europe_West6
	//Madrid, Spain 西班牙马德里
	GEMINI2_LOCATION_Europe_Southwest1
	//Milan, Italy 意大利米兰
	GEMINI2_LOCATION_Europe_West8
	//Finland 芬兰
	GEMINI2_LOCATION_Europe_North1
	//Warsaw, Poland 波兰华沙
	GEMINI2_LOCATION_Europe_Central2

	//Asia Pacific
	//Tokyo, Japan 	日本-东京
	// GEMINI2_LOCATION_AsiaPacific_Northeast1
	//Sydney, Australia 澳大利亚-悉尼
	// GEMINI2_LOCATION_AsiaPacific_Australia_Southeast1
	//Singapore 新加坡
	// GEMINI2_LOCATION_AsiaPacific_Southeast1
	//Seoul, Korea 韩国-首尔
	// GEMINI2_LOCATION_AsiaPacific_Northeast3
	//Taiwan 台湾
	// GEMINI2_LOCATION_AsiaPacific_East1
	//Hong Kong, China 中国-香港
	// GEMINI2_LOCATION_AsiaPacific_East2
	//Mumbai, India 印度-孟买
	// GEMINI2_LOCATION_AsiaPacific_South1

	//Middle East
	//Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// GEMINI2_LOCATION_MiddleEast_Central2
	//Doha, Qatar 卡塔尔-多哈
	// GEMINI2_LOCATION_MiddleEast_Central1
	//Tel Aviv, Israel 以色列-特拉维夫
	// GEMINI2_LOCATION_MiddleEast_West1

	GEMINI2_LOCATION_Global
)

// 创建双向映射
var GET_GEMINI2_LOCATION = map[GEMINI_2_LOCATION]string{
	// United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	GEMINI2_LOCATION_US_East5: "us-east5",
	//Dallas, Texas 德克萨斯州-达拉斯
	GEMINI2_LOCATION_US_South1: "us-south1",
	//Iowa  爱荷华州
	GEMINI2_LOCATION_US_Central1: "us-central1",
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	GEMINI2_LOCATION_US_West4: "us-west4",
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	GEMINI2_LOCATION_US_East1: "us-east1",
	//Northern Virginia 北弗吉尼亚
	GEMINI2_LOCATION_US_East4: "us-east4",
	//Oregon 俄勒冈州
	GEMINI2_LOCATION_US_West1: "us-west1",

	// Canada
	//Montréal 	蒙特利尔
	// GEMINI2_LOCATION_NorthAmerica_Northeast1: "northamerica-northeast1",

	// South America
	//São Paulo, Brazil 巴西圣保罗
	// GEMINI2_LOCATION_SouthAmerica_East1: "southamerica-east1",

	// Europe
	//Netherlands 荷兰
	GEMINI2_LOCATION_Europe_West4: "europe-west4",
	//Paris, France 法国巴黎
	GEMINI2_LOCATION_Europe_West9: "europe-west9",
	//London, United Kingdom 英国伦敦
	// GEMINI2_LOCATION_Europe_West2: "europe-west2",
	//Frankfurt, Germany 德国法兰克福
	// GEMINI2_LOCATION_Europe_West3: "europe-west3",
	// Belgium 比利时
	GEMINI2_LOCATION_Europe_West1: "europe-west1",
	//Zürich, Switzerland 瑞士苏黎世
	// GEMINI2_LOCATION_Europe_West6: "europe-west6",
	//Madrid, Spain 西班牙马德里
	GEMINI2_LOCATION_Europe_Southwest1: "europe-southwest1",
	//Milan, Italy 意大利米兰
	GEMINI2_LOCATION_Europe_West8: "europe-west8",
	//Finland 芬兰
	GEMINI2_LOCATION_Europe_North1: "europe-north1",
	//Warsaw, Poland 波兰华沙
	GEMINI2_LOCATION_Europe_Central2: "europe-central2",

	// Asia Pacific
	//Tokyo, Japan 	日本-东京
	// GEMINI2_LOCATION_AsiaPacific_Northeast1: "asia-northeast1",
	//Sydney, Australia 澳大利亚-悉尼
	// GEMINI2_LOCATION_AsiaPacific_Australia_Southeast1: "australia-southeast1",
	//Singapore 新加坡
	// GEMINI2_LOCATION_AsiaPacific_Southeast1: "asia-southeast1",
	//Seoul, Korea 韩国-首尔
	// GEMINI2_LOCATION_AsiaPacific_Northeast3: "asia-northeast3",
	//Taiwan 台湾
	// GEMINI2_LOCATION_AsiaPacific_East1: "asia-east1",
	//Hong Kong, China 中国-香港
	// GEMINI2_LOCATION_AsiaPacific_East2: "asia-east2",
	//Mumbai, India 印度-孟买
	// GEMINI2_LOCATION_AsiaPacific_South1: "asia-south1",

	// Middle East
	//Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// GEMINI2_LOCATION_MiddleEast_Central2: "me-central2",
	//Doha, Qatar 卡塔尔-多哈
	// GEMINI2_LOCATION_MiddleEast_Central1: "me-central1",
	//Tel Aviv, Israel 以色列-特拉维夫
	// GEMINI2_LOCATION_MiddleEast_West1: "me-west1",

	GEMINI2_LOCATION_Global: "global",
}

var GET_GEMINI2_LOCATION_STRING = map[string]GEMINI_2_LOCATION{
	// United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	"us-east5": GEMINI2_LOCATION_US_East5,
	//Dallas, Texas 德克萨斯州-达拉斯
	"us-south1": GEMINI2_LOCATION_US_South1,
	//Iowa  爱荷华州
	"us-central1": GEMINI2_LOCATION_US_Central1,
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	"us-west4": GEMINI2_LOCATION_US_West4,
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	"us-east1": GEMINI2_LOCATION_US_East1,
	//Northern Virginia 北弗吉尼亚
	"us-east4": GEMINI2_LOCATION_US_East4,
	//Oregon 俄勒冈州
	"us-west1": GEMINI2_LOCATION_US_West1,

	// Canada
	//Montréal 	蒙特利尔
	// "northamerica-northeast1": GEMINI2_LOCATION_NorthAmerica_Northeast1,

	// South America
	//São Paulo, Brazil 巴西圣保罗
	// "southamerica-east1": GEMINI2_LOCATION_SouthAmerica_East1,

	// Europe
	//Netherlands 荷兰
	"europe-west4": GEMINI2_LOCATION_Europe_West4,
	//Paris, France 法国巴黎
	"europe-west9": GEMINI2_LOCATION_Europe_West9,
	//London, United Kingdom 英国伦敦
	// "europe-west2": GEMINI2_LOCATION_Europe_West2,
	//Frankfurt, Germany 德国法兰克福
	// "europe-west3": GEMINI2_LOCATION_Europe_West3,
	// Belgium 比利时
	"europe-west1": GEMINI2_LOCATION_Europe_West1,
	//Zürich, Switzerland 瑞士苏黎世
	// "europe-west6": GEMINI2_LOCATION_Europe_West6,
	//Madrid, Spain 西班牙马德里
	"europe-southwest1": GEMINI2_LOCATION_Europe_Southwest1,
	//Milan, Italy 意大利米兰
	"europe-west8": GEMINI2_LOCATION_Europe_West8,
	//Finland 芬兰
	"europe-north1": GEMINI2_LOCATION_Europe_North1,
	//Warsaw, Poland 波兰华沙
	"europe-central2": GEMINI2_LOCATION_Europe_Central2,

	// // Asia Pacific
	// //Tokyo, Japan 	日本-东京
	// "asia-northeast1": GEMINI2_LOCATION_AsiaPacific_Northeast1,
	// //Sydney, Australia 澳大利亚-悉尼
	// "australia-southeast1": GEMINI2_LOCATION_AsiaPacific_Australia_Southeast1,
	// //Singapore 新加坡
	// "asia-southeast1": GEMINI2_LOCATION_AsiaPacific_Southeast1,
	// //Seoul, Korea 韩国-首尔
	// "asia-northeast3": GEMINI2_LOCATION_AsiaPacific_Northeast3,
	// //Taiwan 台湾
	// "asia-east1": GEMINI2_LOCATION_AsiaPacific_East1,
	// //Hong Kong, China 中国-香港
	// "asia-east2": GEMINI2_LOCATION_AsiaPacific_East2,
	// //Mumbai, India 印度-孟买
	// "asia-south1": GEMINI2_LOCATION_AsiaPacific_South1,

	// // Middle East
	// //Dammam, Saudi Arabia 沙特阿拉伯-达曼
	// "me-central2": GEMINI2_LOCATION_MiddleEast_Central2,
	// //Doha, Qatar 卡塔尔-多哈
	// "me-central1": GEMINI2_LOCATION_MiddleEast_Central1,
	// //Tel Aviv, Israel 以色列-特拉维夫
	// "me-west1": GEMINI2_LOCATION_MiddleEast_West1,

	"global": GEMINI2_LOCATION_Global,
}

// 获取 GEMINI2_LOCATION 对应的字符串
func Get_GEMINI2_LOCATION_STRING(GEMINI2_LOCATION GEMINI_2_LOCATION) string {
	return GET_GEMINI2_LOCATION[GEMINI2_LOCATION]
}

// 获取字符串对应的 GEMINI2_LOCATION
func Get_GEMINI2_LOCATION(GEMINI2_LOCATIONStr string) (GEMINI_2_LOCATION, bool) {
	GEMINI2_LOCATION, exists := GET_GEMINI2_LOCATION_STRING[GEMINI2_LOCATIONStr]
	return GEMINI2_LOCATION, exists
}
