package vertex

type LOCATION int32

// https://cloud.google.com/vertex-ai/generative-ai/docs/learn/locations#middle-east
const (
	//United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	LOCATION_US_East5 LOCATION = iota
	//Dallas, Texas 德克萨斯州-达拉斯
	LOCATION_US_South1
	//Iowa  爱荷华州
	LOCATION_US_Central1
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	LOCATION_US_West4
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	LOCATION_US_East1
	//Northern Virginia 北弗吉尼亚
	LOCATION_US_East4
	//Oregon 俄勒冈州
	LOCATION_US_West1

	//Canada
	//Montréal 	蒙特利尔
	LOCATION_NorthAmerica_Northeast1

	//South America
	//São Paulo, Brazil 巴西圣保罗
	LOCATION_SouthAmerica_East1

	//Europe
	//Netherlands 荷兰
	LOCATION_Europe_West4
	//Paris, France 法国巴黎
	LOCATION_Europe_West9
	//London, United Kingdom 英国伦敦
	LOCATION_Europe_West2
	//Frankfurt, Germany 德国法兰克福
	LOCATION_Europe_West3
	// Belgium 比利时
	LOCATION_Europe_West1
	//Zürich, Switzerland 瑞士苏黎世
	LOCATION_Europe_West6
	//Madrid, Spain 西班牙马德里
	LOCATION_Europe_Southwest1
	//Milan, Italy 意大利米兰
	LOCATION_Europe_West8
	//Finland 芬兰
	LOCATION_Europe_North1
	//Warsaw, Poland 波兰华沙
	LOCATION_Europe_Central2

	//Asia Pacific
	//Tokyo, Japan 	日本-东京
	LOCATION_AsiaPacific_Northeast1
	//Sydney, Australia 澳大利亚-悉尼
	LOCATION_AsiaPacific_Australia_Southeast1
	//Singapore 新加坡
	LOCATION_AsiaPacific_Southeast1
	//Seoul, Korea 韩国-首尔
	LOCATION_AsiaPacific_Northeast3
	//Taiwan 台湾
	LOCATION_AsiaPacific_East1
	//Hong Kong, China 中国-香港
	LOCATION_AsiaPacific_East2
	//Mumbai, India 印度-孟买
	LOCATION_AsiaPacific_South1

	//Middle East
	//Dammam, Saudi Arabia 沙特阿拉伯-达曼
	LOCATION_MiddleEast_Central2
	//Doha, Qatar 卡塔尔-多哈
	LOCATION_MiddleEast_Central1
	//Tel Aviv, Israel 以色列-特拉维夫
	LOCATION_MiddleEast_West1
)

// 创建双向映射
var VertexLocationToString = map[LOCATION]string{
	// United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	LOCATION_US_East5: "us-east5",
	//Dallas, Texas 德克萨斯州-达拉斯
	LOCATION_US_South1: "us-south1",
	//Iowa  爱荷华州
	LOCATION_US_Central1: "us-central1",
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	LOCATION_US_West4: "us-west4",
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	LOCATION_US_East1: "us-east1",
	//Northern Virginia 北弗吉尼亚
	LOCATION_US_East4: "us-east4",
	//Oregon 俄勒冈州
	LOCATION_US_West1: "us-west1",

	// Canada
	//Montréal 	蒙特利尔
	LOCATION_NorthAmerica_Northeast1: "northamerica-northeast1",

	// South America
	//São Paulo, Brazil 巴西圣保罗
	LOCATION_SouthAmerica_East1: "southamerica-east1",

	// Europe
	//Netherlands 荷兰
	LOCATION_Europe_West4: "europe-west4",
	//Paris, France 法国巴黎
	LOCATION_Europe_West9: "europe-west9",
	//London, United Kingdom 英国伦敦
	LOCATION_Europe_West2: "europe-west2",
	//Frankfurt, Germany 德国法兰克福
	LOCATION_Europe_West3: "europe-west3",
	// Belgium 比利时
	LOCATION_Europe_West1: "europe-west1",
	//Zürich, Switzerland 瑞士苏黎世
	LOCATION_Europe_West6: "europe-west6",
	//Madrid, Spain 西班牙马德里
	LOCATION_Europe_Southwest1: "europe-southwest1",
	//Milan, Italy 意大利米兰
	LOCATION_Europe_West8: "europe-west8",
	//Finland 芬兰
	LOCATION_Europe_North1: "europe-north1",
	//Warsaw, Poland 波兰华沙
	LOCATION_Europe_Central2: "europe-central2",

	// Asia Pacific
	//Tokyo, Japan 	日本-东京
	LOCATION_AsiaPacific_Northeast1: "asia-northeast1",
	//Sydney, Australia 澳大利亚-悉尼
	LOCATION_AsiaPacific_Australia_Southeast1: "australia-southeast1",
	//Singapore 新加坡
	LOCATION_AsiaPacific_Southeast1: "asia-southeast1",
	//Seoul, Korea 韩国-首尔
	LOCATION_AsiaPacific_Northeast3: "asia-northeast3",
	//Taiwan 台湾
	LOCATION_AsiaPacific_East1: "asia-east1",
	//Hong Kong, China 中国-香港
	LOCATION_AsiaPacific_East2: "asia-east2",
	//Mumbai, India 印度-孟买
	LOCATION_AsiaPacific_South1: "asia-south1",

	// Middle East
	//Dammam, Saudi Arabia 沙特阿拉伯-达曼
	LOCATION_MiddleEast_Central2: "me-central2",
	//Doha, Qatar 卡塔尔-多哈
	LOCATION_MiddleEast_Central1: "me-central1",
	//Tel Aviv, Israel 以色列-特拉维夫
	LOCATION_MiddleEast_West1: "me-west1",
}

var VertexStringToLocation = map[string]LOCATION{
	// United States
	//Columbus, Ohio 俄亥俄州-哥伦布市
	"us-east5": LOCATION_US_East5,
	//Dallas, Texas 德克萨斯州-达拉斯
	"us-south1": LOCATION_US_South1,
	//Iowa  爱荷华州
	"us-central1": LOCATION_US_Central1,
	//Las Vegas, Nevada 内华达州-拉斯维加斯
	"us-west4": LOCATION_US_West4,
	//Moncks Corner, South Carolina 南卡罗来纳州-蒙克斯科纳
	"us-east1": LOCATION_US_East1,
	//Northern Virginia 北弗吉尼亚
	"us-east4": LOCATION_US_East4,
	//Oregon 俄勒冈州
	"us-west1": LOCATION_US_West1,

	// Canada
	//Montréal 	蒙特利尔
	"northamerica-northeast1": LOCATION_NorthAmerica_Northeast1,

	// South America
	//São Paulo, Brazil 巴西圣保罗
	"southamerica-east1": LOCATION_SouthAmerica_East1,

	// Europe
	//Netherlands 荷兰
	"europe-west4": LOCATION_Europe_West4,
	//Paris, France 法国巴黎
	"europe-west9": LOCATION_Europe_West9,
	//London, United Kingdom 英国伦敦
	"europe-west2": LOCATION_Europe_West2,
	//Frankfurt, Germany 德国法兰克福
	"europe-west3": LOCATION_Europe_West3,
	// Belgium 比利时
	"europe-west1": LOCATION_Europe_West1,
	//Zürich, Switzerland 瑞士苏黎世
	"europe-west6": LOCATION_Europe_West6,
	//Madrid, Spain 西班牙马德里
	"europe-southwest1": LOCATION_Europe_Southwest1,
	//Milan, Italy 意大利米兰
	"europe-west8": LOCATION_Europe_West8,
	//Finland 芬兰
	"europe-north1": LOCATION_Europe_North1,
	//Warsaw, Poland 波兰华沙
	"europe-central2": LOCATION_Europe_Central2,

	// Asia Pacific
	//Tokyo, Japan 	日本-东京
	"asia-northeast1": LOCATION_AsiaPacific_Northeast1,
	//Sydney, Australia 澳大利亚-悉尼
	"australia-southeast1": LOCATION_AsiaPacific_Australia_Southeast1,
	//Singapore 新加坡
	"asia-southeast1": LOCATION_AsiaPacific_Southeast1,
	//Seoul, Korea 韩国-首尔
	"asia-northeast3": LOCATION_AsiaPacific_Northeast3,
	//Taiwan 台湾
	"asia-east1": LOCATION_AsiaPacific_East1,
	//Hong Kong, China 中国-香港
	"asia-east2": LOCATION_AsiaPacific_East2,
	//Mumbai, India 印度-孟买
	"asia-south1": LOCATION_AsiaPacific_South1,

	// Middle East
	//Dammam, Saudi Arabia 沙特阿拉伯-达曼
	"me-central2": LOCATION_MiddleEast_Central2,
	//Doha, Qatar 卡塔尔-多哈
	"me-central1": LOCATION_MiddleEast_Central1,
	//Tel Aviv, Israel 以色列-特拉维夫
	"me-west1": LOCATION_MiddleEast_West1,
}

// 获取 LOCATION 对应的字符串
func GetLocationString(location LOCATION) string {
	return VertexLocationToString[location]
}

// 获取字符串对应的 LOCATION
func GetLocationFromString(locationStr string) (LOCATION, bool) {
	location, exists := VertexStringToLocation[locationStr]
	return location, exists
}
