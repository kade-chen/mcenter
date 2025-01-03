package impl

import (
	"context"
	"fmt"
	"os"

	"github.com/kade-chen/library/tools/format"
	"github.com/kade-chen/mcenter/apps/vertex"
	"google.golang.org/genai"
)

func (g *gemini) NoStreamingGenerateContent(ctx context.Context) {

	config := vertex.NewTestGenerateContentConfig()

	a, err := g.clientsv2.Models.GenerateContent(ctx, "gemini-2.0-flash-exp", []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "生成一张苹果照片",
				},
			},
		},
	}, config)

	if err != nil {
		fmt.Println("Error: InlineData is nil")
		return
		// panic(err)
	}
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
}
