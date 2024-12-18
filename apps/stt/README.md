# 文档明细

**NOTE:** Speech To Text V1

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
