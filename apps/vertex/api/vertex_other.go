package api

import (
	"path"
	"strings"

	"cloud.google.com/go/vertexai/genai"
)

var mimeType = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".mp3":  "audio/mpeg",
	".wav":  "audio/wav",
	".mp4":  "video/mp4",
	".pdf":  "application/pdf",
	".txt":  "text/plain",
	".json": "text/plain",
	".xml":  "text/xml",
	".aac":  "audio/vnd.dlna.adts",
	// ".svg":  "image/svg+xml",
	".csv": "text/csv",
	// ".pcm": "application/octet-stream",
	".ogg": "audio/ogg",
	// ".zip": "application/x-zip-compressed",
}

// get filedata from url
func (g *geminiHandler) getFileDataUrl(url string) (parts []genai.Part) {
	// A := make([]genai.Part, 50)
	// 使用 strings.Split 将长字符串按逗号分割成多个部分
	urlSlice := strings.Split(url, ",")

	// 输出分割后的所有 URL
	for _, url := range urlSlice {
		//0.get file extension form url and convert to lower case
		extension := strings.ToLower(path.Ext(url))

		if mimieType, exists := mimeType[extension]; exists {
			g.log.Info().Msgf("File extension: %s, MIME type: %s", extension, mimieType)
			parts = append(parts, genai.FileData{MIMEType: mimieType, FileURI: url})
			// return genai.FileData{MIMEType: mimieType, FileURI: url}
		} else {
			// 默认返回 application/octet-stream MIME 类型
			g.log.Error().Msgf("By default, the application/octet-stream mime type is returned")
			parts = append(parts, genai.FileData{
				MIMEType: "application/octet-stream",
				FileURI:  url,
			})
		}
	}
	return parts
}

// t1 := mimetype.Detect(video).String()
// blob := genai.Blob{
// 	MIMEType: t1,
// 	Data:     decoded,
// }
