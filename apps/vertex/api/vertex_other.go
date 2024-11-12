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

func (g *geminiHandler) getFileDataUrl(url string) genai.FileData {
	//0.get file extension form url and convert to lower case
	extension := strings.ToLower(path.Ext(url))

	if mimieType, exists := mimeType[extension]; exists {
		g.log.Info().Msgf("File extension: %s, MIME type: %s", extension, mimieType)
		return genai.FileData{MIMEType: mimieType, FileURI: url}
	}

	// 默认返回 application/octet-stream MIME 类型
	g.log.Info().Msgf("By default, the application/octet-stream mime type is returned")
	return genai.FileData{
		MIMEType: "application/octet-stream",
		FileURI:  url,
	}
}
