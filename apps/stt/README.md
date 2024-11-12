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
