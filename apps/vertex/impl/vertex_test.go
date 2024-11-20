package impl_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/vertex"
	_ "github.com/kade-chen/mcenter/apps/vertex/impl"
	"google.golang.org/api/iterator"
	// "google.golang.org/appengine/log"
)

var (
	ctx  = context.Background()
	impl vertex.Service
)

func TestNoStreamingGenerateContent(t *testing.T) {
	// fmt.Println(ioc.Controller().List(), ioc.Config().List(), *new(int32))
	req := vertex.NewGenaiSetting()
	// req.SystemInstruction.Parts = []genai.Part{
	// 	genai.Text("必须中文回答"),
	// 	// genai.FileData{},
	// 	// genai.Blob{},
	// }
	req.SystemInstruction.Parts = "必须中文回答"
	req.ModelName = "gemini-1.5-flash-002"
	prompt := genai.Text("what is images")
	img := genai.FileData{
		MIMEType: "image/jpeg",
		FileURI:  "https://storage.googleapis.com/kadeccc/cat.jpg",
	}
	resp, err := impl.NoStreamingGenerateContent(ctx, req, prompt,img)
	if err != nil {
		fmt.Println("--------", err)
	}
	rb, err := json.MarshalIndent(resp, "", "  ")
	// rb, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("--------", err)
	}
	t.Log(string(rb))
	// impl.GetLocationFromString()
}

func TestStreamingGenerateContent(t *testing.T) {
	req := vertex.NewGenaiSetting()
	// req.SystemInstruction.Parts = []genai.Part{
	// 	genai.Text("音频里说了什么"),
	// 	// genai.FileData{},
	// 	// genai.Blob{},
	// }
	req.ModelName = "gemini-1.5-pro-002"
	prompt := genai.Text("k8s是什么?250个字解释")
	// e := genai.FileData{
	// 	MIMEType: "audio/wav",
	// 	FileURI:  "gs://kadeccc/3.wav",
	// }
	// // genai.FileData{
	// // 	MIMEType: "image/jpeg",
	// // 	FileURI:  "gs://kadeccc/cat.jpg",
	// // },
	// genai.Text("音频里说了什么?"),
	// genai.Text("音频里说了什么?"),
	GenerateContentResponseIterator := impl.StreamingGenerateContent(ctx, req, prompt)
	for {
		resp, err := GenerateContentResponseIterator.Next()
		if err == iterator.Done {
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			exception.NewInternalServerError("error: %v", err)
		}
		if err != nil {
			exception.NewInternalServerError("error: %v", err)
		}

		// fmt.Fprint(os.Stdout, "generated response: ")
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				// fmt.Fprintf(os.Stdout, "%v", p)
				// fmt.Fprint(os.Stdout, p)
				fmt.Print(p)
			}
		}
		// fmt.Fprint(os.Stdout, "\n")
	}
	// t.Log(a)
}

func TestStreamingGenerateContent2(t *testing.T) {
	req := vertex.NewGenaiSetting()
	req.ModelName = "gemini-1.5-pro-002"
	prompt := genai.Text("k8s是什么？")
	prompt1 := genai.Text("200字回复")
	var cc []genai.Part = []genai.Part{
		prompt,
		prompt1,
	}
	GenerateContentResponseIterator := impl.StreamingGenerateContent(ctx, &vertex.GenaiSetting{
		ModelName: "gemini-1.5-pro-002",
	}, cc...)
	for {
		resp, err := GenerateContentResponseIterator.Next()
		if err == iterator.Done {
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			exception.NewInternalServerError("error: %v", err)
		}
		if err != nil {
			exception.NewInternalServerError("error: %v", err)
		}

		// fmt.Fprint(os.Stdout, "generated response: ")
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				fmt.Fprintf(os.Stdout, "%v", p)
				// fmt.Fprint(os.Stdout, p)
				// fmt.Print(p)
			}
		}
		fmt.Fprint(os.Stdout, "\n")
	}
	// t.Log(a)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "../../../etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName).(vertex.Service)
}
