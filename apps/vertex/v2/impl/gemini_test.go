package impl_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
	"github.com/kade-chen/mcenter/apps/vertex"
	_ "github.com/kade-chen/mcenter/apps/vertex/v2/impl"
	"google.golang.org/genai"
	// "google.golang.org/appengine/log"
)

var (
	ctx  = context.Background()
	impl vertex.ServiceGemini2
)

func TestNoGenerateContent(t *testing.T) {
	config := vertex.NewGemini2Config()
	config.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
		},
	}
	a, err := impl.NoStreamingGenerateContent(ctx, config)

	if err != nil {
		// fmt.Printf("Error: InlineData is nil %v",err)
		t.Fatal(err)
		// panic(err)
	}
	fmt.Println("InlineData:", format.ToJSON(a))
}

func TestNoStreamingGenerateContent(t *testing.T) {
	config := vertex.NewGemini2Config()
	config.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "生成一张苹果照片",
				},
			},
		},
	}
	a, err := impl.NoStreamingGenerateContent(ctx, config)

	if err != nil {
		fmt.Println("Error: InlineData is nil")
		return
		// panic(err)
	}
	fmt.Println("InlineData:", format.ToJSON(a))
	if a.Candidates[0].Content != nil {
		fmt.Println("Content:", len(a.Candidates[0].Content.Parts))
		fmt.Printf("%#v\n,", format.ToJSON(a.Candidates[0].Content.Parts))
		// fmt.Println("Content0:", format.ToJSON(a.Candidates[0].Content.Parts[0]))
		// fmt.Println("Content1:", format.ToJSON(a.Candidates[0].Content.Parts[1]))
		// fmt.Println("Content2:", format.ToJSON(a.Candidates[0].Content.Parts[2]))
		// 		Content: 3
		// Content0: {
		//   "text": "好的，为你生成一张哆啦A梦的图片。\n"
		// }
		// Content1: {
		//   "inlineData": {
		//     "data": "
		//     "mimeType": "image/png"
		//   }
		// }
		// Content2: {
		//   "text": "\n希望你喜欢！如果你有其他要求，请随时提出。\n"
		// }
		if len(a.Candidates[0].Content.Parts) > 1 {
			err = os.WriteFile("image1.png", a.Candidates[0].Content.Parts[1].InlineData.Data, 0644)
			if err != nil {
				fmt.Println("Error: create image.png failed")
				return
			}
		} else {
			fmt.Println("生成一半被block")
		}
	} else {
		fmt.Println("Error: InlineData is block")
		return
	}
	// fmt.Printf("%#v\n", format.ToJSON(a))
	// fmt.Printf("%#v\n", a.Candidates[0].Content)
	// a := format.ToJSON(resp)
	// fmt.Println(a, err)
}

func TestStreamingGenerateContent(t *testing.T) {
	config := vertex.NewGemini2Config()
	config.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
		},
	}
	iter1 := impl.StreamingGenerateContent(ctx, config)
	fmt.Println("InlineData:", format.ToJSON(iter1))
	for a, err := range iter1 {
		if err != nil {
			exception.NewInternalServerError("ERROR: %v", err.Error())
		}
		fmt.Print(a.Candidates[0].Content.Parts[0].Text)
	}
	// a := format.ToJSON(resp)
	// fmt.Println(a, err)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName2).(vertex.ServiceGemini2)
}
