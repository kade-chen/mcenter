# 文档明细

**NOTE** : Speech To Text V1

1.[Speech To Text V1 Go Client Interfaces](https://cloud.google.com/go/docs/reference/cloud.google.com/go/speech/latest/apiv1)

2.[Speech-to-Text supported languages](https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages)

3.[Package google.cloud.speech.v1, requests Parametars explain](https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1)

4.[Supported class tokens explain](https://cloud.google.com/speech-to-text/docs/class-tokens)

5.[Supported spoken punctuation explain](https://cloud.google.com/speech-to-text/docs/spoken-punctuation)

6.[Supported spoken emoji explain](https://cloud.google.com/speech-to-text/docs/spoken-emoji)

7.[RecognitionConfig explain](https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionconfig)

8.[RecognitionAudio explain](https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#recognitionaudio)

9.[Supported voices and languages](https://cloud.google.com/text-to-speech/docs/voices)

10.[Detect different speakers in an audio recording](https://cloud.google.com/speech-to-text/docs/multiple-voices#speech_transcribe_diarization_beta-go) 服务支持人声分离识别，可检测录音中不同的讲话人，并在 转录的音频数据中为不同的讲话人添加标签

11.[Transcribe audio with multiple channels](https://cloud.google.com/speech-to-text/docs/multi-channel) 最多可支持 8 个音轨

12.[Send a recognition request with model adaptation](https://cloud.google.com/speech-to-text/docs/adaptation)

13.[Measure and improve speech accuracy](https://cloud.google.com/speech-to-text/docs/speech-accuracy)

14.[Streaming Speech-to-Text API Recognition Requests](https://cloud.google.com/speech-to-text/docs/speech-to-text-requests#streaming-recognition)

15.[Quotas and limits](https://cloud.google.com/speech-to-text/v2/quotas)

16.[处理时长](https://cloud.google.com/speech-to-text/docs/speech-to-text-requests#synchronous-recognition)

17.[Release notes](https://cloud.google.com/speech-to-text/docs/release-notes?hl=zh-cn)

18.[Endpoints](https://cloud.google.com/speech-to-text/docs/endpoints#speech-sync-recognize-protocol)

19.[Voice activity events and timeouts](https://cloud.google.com/speech-to-text/v2/docs/voice-activity-events)

20.[sla](https://cloud.google.com/speech/sla?hl=en)

21.[ASR 语音模型响应指标压力测试报告](https://docs.google.com/document/d/12MujZnydR9QYBhZXRT7OZZnnwomhnywQGwv4qRyYVzI/edit?tab=t.flgjvcdcrtm9#heading=h.wv1941qd7cy6)

**NOTE:** Speech To Text V2

1.[Speech To Text V2 Go Client Interfaces](https://cloud.google.com/go/docs/reference/cloud.google.com/go/speech/latest/apiv2)

2.[Speech-to-Text V2 supported languages](https://cloud.google.com/speech-to-text/v2/docs/speech-to-text-supported-languages)

3.[Package google.cloud.speech.v2, requests Parametars explain](https://cloud.google.com/speech-to-text/v2/docs/reference/rpc/google.cloud.speech.v2)

4.[Supported class tokens explain](https://cloud.google.com/speech-to-text/v2/docs/class-tokens)

5.[Supported spoken punctuation explain](https://cloud.google.com/speech-to-text/v2/docs/spoken-punctuation)

6.[Supported spoken emoji explain](https://cloud.google.com/speech-to-text/v2/docs/spoken-emoji)

7.[Supported transcription models](https://cloud.google.com/speech-to-text/v2/docs/transcription-model#transcription_models)

7.[EndPoints](https://cloud.google.com/speech-to-text/v2/docs/reference/rest#service-endpoint)

# GRPC AUTHENTICATION EXAMPLE


| 选项                         | 功能                                              | 侧重点                                         |
|:----------------------------|:--------------------------------------------------|:---------------------------------------------:|
| **WithTransportCredentials** | 配置底层 TLS 加密和服务器身份验证                 | 用于建立加密的安全通道（SSL/TLS）。            |
| **WithPerRPCCredentials**    | 配置每个 RPC 调用的身份认证信息                   | 用于提供 OAuth 令牌，确保每次请求都带有授权信息。 |

## 1. service account key.json

```go
package main

import (
    "context"
    "fmt"

    "cloud.google.com/go/speech/apiv2/speechpb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/credentials/oauth"
)

func main() {
    // 1. 加载服务账号 JSON 文件
    credentialsFilePath := "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json"
    perRPCCreds, err := oauth.NewServiceAccountFromFile(credentialsFilePath, "https://www.googleapis.com/auth/cloud-platform")
    if err != nil {
        fmt.Printf("Failed to create credentials: %v\n", err)
        return
    }

    // 2. 创建 gRPC 客户端连接
    conn, err := grpc.Dial(
        "speech.googleapis.com:443",
        grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
        grpc.WithPerRPCCredentials(perRPCCreds),
    )
    if err != nil {
        fmt.Printf("Failed to create gRPC connection: %v\n", err)
        return
    }
    defer conn.Close()

    // 3. 初始化 Speech 客户端
    client := speechpb.NewSpeechClient(conn)

    // 4. 调用 ListRecognizers API
    resp, err := client.ListRecognizers(context.Background(), &speechpb.ListRecognizersRequest{
        Parent: "projects/kade-poc/locations/global",
    })
    if err != nil {
        fmt.Printf("Failed to list recognizers: %v\n", err)
        return
    }

    // 5. 输出 API 响应
    fmt.Printf("Recognizers: %v\n", resp)
}

```

## 2. gcloud auth application-default login

```go
	google.NewDefaultCredentials() 是 Google 官方提供的用于简化认证的默认凭证管理器。它会按照以下顺序查找和加载凭证：

		1.	环境变量 GOOGLE_APPLICATION_CREDENTIALS如果设置了该环境变量，它会尝试加载服务账号 JSON 文件中的凭证。

		2.	本地用户凭证（Application Default Credentials, ADC）如果您使用 gcloud auth application-default login 进行了身份验证，google.NewDefaultCredentials() 会使用本地用户的默认凭证。这种方式在开发过程中很常用。

		在 Mac/Linux 上，通常位于 ~/.config/gcloud/application_default_credentials.json。

		3.	运行环境中的默认凭证如果您的代码运行在 Google Cloud 上（如 GKE、Compute Engine 等），会自动使用运行环境中	分配给实例的服务账号。

		4.	通过其他配置加载凭证如 Kubernetes Secret、默认配置文件等。
```
```go
package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/google"
)

func main() {
	// 手动指定服务账号 JSON 文件路径
	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json")

	// 创建 gRPC 客户端
	conn, err := grpc.Dial(
		"speech.googleapis.com:443",
		grpc.WithTransportCredentials(google.NewDefaultCredentials().TransportCredentials()),
		grpc.WithPerRPCCredentials(google.NewDefaultCredentials().PerRPCCredentials()),
	)
	if err != nil {
		fmt.Printf("Failed to create gRPC client: %v\n", err)
		return
	}
	defer conn.Close()

	// 初始化 Speech 客户端
	SpeechClient := speechpb.NewSpeechClient(conn)

	// 调用 ListRecognizers
	resp, err := SpeechClient.ListRecognizers(context.Background(), &speechpb.ListRecognizersRequest{
		Parent: "projects/kade-poc/locations/global", // 替换为有效区域
	})
	if err != nil {
		fmt.Printf("Failed to list recognizers: %v\n", err)
		return
	}

	// 输出结果
	fmt.Printf("List of recognizers: %v\n", resp)
}
```

```go
设置为表示自然语言

	1.	$NUMERIC_SEQUENCE：用于识别一系列数字，比如电话号码、信用卡号、序列号等。这类提示可以提高识别复杂数字串的准确性。
	2.	$DATE：表示日期的类别，可以帮助正确识别不同格式的日期，例如“2024年10月10日”或“October 10, 2024”。
	3.	$TIME：识别时间表达式，例如“下午三点”或“3:00 PM”，有助于避免误识别时间格式。
	4.	$CURRENCY：表示货币的类别，可以提高系统对金钱相关表达的准确性，例如“100美元”或“€50”。
	5.	$ORDINAL：用于表示序数词的类别，例如“第一”、“第二”、“第三”，可以帮助正确转录数字的顺序形式。
	6.	$ADDRESS：与地址相关的词汇类别，用于提高识别地址信息的精度，例如“1234 Main St”或“5th Avenue”。
	7.	$EMAIL：表示电子邮件地址的类别，有助于语音系统正确识别和转录电子邮件地址中的字符和符号。
	8.	$WEBSITE_URL：识别和转录网址，有助于提高对 URL 格式文本的识别准确性。
	9.	$PROPER_NAME：用于识别人名、地名、公司名等专有名词，减少在这些词汇上的误差。
	10.	$PRODUCT_NAME：表示产品名称的类别，可以提高语音识别对特定产品名称的识别精度，特别适用于电商场景。
```

```go
Cloud Speech-to-Text 服务 及 Google Cloud 提供相关数据安全与隐私保护说明和证明。
参考：
https://cloud.google.com/speech-to-text/docs/data-usage-faq
https://cloud.google.com/terms
https://cloud.google.com/privacy/common-privacy-principles
https://cloud.google.com/sensitive-data-protection/docs/support/data-security?hl=zh-cn

提供 ISO27001资质证书
https://cloud.google.com/security/compliance/iso-27001?hl=zh-cn

其他相关合规认证参考：
https://cloud.google.com/security/compliance/offerings?hl=zh-cn
https://cloud.google.com/privacy/gdpr?hl=zh-cn
https://cloud.google.com/security/compliance/ccpa?hl=zh-cn

```

```go
数据输入安全：

Cloud Speech-to-Text 服务 及 Google Cloud 提供相关数据安全与隐私保护说明和证明。
参考：
https://cloud.google.com/speech-to-text/docs/data-usage-faq
https://cloud.google.com/terms
https://cloud.google.com/privacy/common-privacy-principles
https://cloud.google.com/sensitive-data-protection/docs/support/data-security?hl=zh-cn

数据输出安全：

Cloud Speech-to-Text 服务可与 Cloud Data Loss Prevention 服务一起构建敏感数据保护体系，对敏感数据/信息进行分类和去标识化。Cloud Data Loss Prevention 服务支持全球以及不同国家的敏感数据保护使用信息类型，并支持用户通过自定义检测器来自定义检测行为，以便敏感数据保护对与用户指定的模式匹配的敏感数据进行检查或去标识化。

Cloud Data Loss Prevention 服务
参考：
https://cloud.google.com/sensitive-data-protection/docs/sensitive-data-protection-overview?hl=zh-cn
https://cloud.google.com/sensitive-data-protection/docs/supported-file-types?hl=zh-cn
https://cloud.google.com/sensitive-data-protection/docs/infotypes-reference?hl=zh-cn

完整解决方案可参考：
https://github.com/GoogleCloudPlatform/dataflow-speech-redaction
```

```go
Google Cloud 提供相关安全认证证明，服务符合行业标准。
ISO
https://cloud.google.com/security/compliance/iso-9001?hl=zh-cn
https://cloud.google.com/security/compliance/iso-27017?hl=zh-cn
https://cloud.google.com/security/compliance/iso-27018?hl=zh-cn

SOC 2/3
https://cloud.google.com/security/compliance/soc-2?hl=zh-cn
https://cloud.google.com/security/compliance/soc-3?hl=zh-cn

PCI DSS
https://cloud.google.com/security/compliance/pci-dss?hl=zh-cn

更多合规认证参考：
https://cloud.google.com/security/compliance/offerings?hl=zh-cn
```
