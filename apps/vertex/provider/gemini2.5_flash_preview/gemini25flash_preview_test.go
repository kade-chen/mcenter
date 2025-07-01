package gemini25flashpreview_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider" //ioc
	"google.golang.org/genai"
	// _ "github.com/kade-chen/mcenter/apps/vertex/provider/registry"
	// _ "github.com/kade-chen/mcenter/apps/vertex/provider/gemini2_flash"
)

var (
	impl provider.ModelIssuer
	ctx  = context.Background()
)

func TestMain(t *testing.T) {
	fmt.Println(ioc.Config().List())
	fmt.Println(provider.ProviderRegistry())
	ppp := vertex.NewGemini2Config()
	ppp.ModelName = "gemini-2.5-pro-preview-03-25"
	ppp.Location = "global"
	// ppp.GenerateContentConfig.Logprobs = generics.Generics[int32](1)
	// ppp.ModelName = "gemini-2.5-flash-preview-04-17"
	ppp.GenerateContentConfig.ThinkingConfig.IncludeThoughts = false
	// ppp.GenerateContentConfig.ThinkingConfig.ThinkingBudget = generics.Generics[int32](0)
	// "gemini-2.5-pro-preview-03-25"
	//"gemini-2.5-pro-preview-05-06"
	b := make(map[string]string, 100)
	b["kade-gemini-2.5-pro-preview-03-25"] = "gemini-2.5-pro-preview-03-25-kade-01"
	b["kade-gemini-2.5-pro-preview-03-25"] = "gemini-2.5-pro-preview-03-25-kade-02"
	// b["gemini-2.5-pro-preview-05-06"] = "gemini-2.5-pro-preview-05-06"
	ppp.GenerateContentConfig.Labels = b
	ppp.Contents = []*genai.Content{
		// {
		// 	Parts: []*genai.Part{
		// 		{
		// 			FileData: &genai.FileData{
		// 				FileURI:  "http://ofasys-multimodal-wlcb-3.oss-cn-wulanchabu.aliyuncs.com/yongqi.wyq/muchin_eval_audio_16k/s_0003_src.mp3?OSSAccessKeyId=LTAI5tNSZEaGH5to7YysTUXz&Expires=1749492017&Signature=TIUlEmVrDuLkB9s39%2FwvKcd1lsw%3D",
		// 				MIMEType: "audio/mpeg",
		// 			},
		// 		},
		// 	},
		// 	Role: "user",
		// },
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
			Role: "user",
		},
	}
	a1 := impl.StreamingGenerateContent(ctx, ppp)
	for a, err := range a1 {
		if err != nil {
			fmt.Println("aaaa", a, err)
			break
		}
		fmt.Println(a.Candidates[0].Content.Parts[0].Text)
	}
	// a, err := impl.NoStreamingGenerateContent(ctx, ppp)
	// if err != nil {
	// 	fmt.Println("cc", err)
	// 	panic(err)
	// }
	// fmt.Println(format.ToJSON(a))
}

func TestMain1(t *testing.T) {
	ppp := vertex.NewGemini2Config()
	ppp.ModelName = "gemini-2.5-flash-preview-tts"

	// creds1, _ := credentials.DetectDefault(&credentials.DetectOptions{
	// 	Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
	// 	// CredentialsJSON: data,
	// 	CredentialsFile: "/Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json",
	// })

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		//目前不支持，另外这个apikey，不能backend，projectid,location,credentials都填，否则会报错
		//https://console.cloud.google.com/apis/credentials?project=kade-poc (Vertex AI API/Generative Language API)
		//https://console.developers.google.com/apis/api/generativelanguage.googleapis.com/overview?project=499036589398
		APIKey:  "AIzaSyBih5Bmh8I5dK3QxiWF7Plqr35O_4NIvIY",
		// Backend: genai.BackendGeminiAPI,
		// Project: "kade-poc",
		// Project:     ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_5_Flash_Preview,
		// Location:    "us-central1",
		// Credentials: creds1,

		HTTPOptions: genai.HTTPOptions{APIVersion: "v1beta"},
	})
	if err != nil {
		fmt.Println("aaaa", err)
	}
	ppp.Contents = []*genai.Content{
		// {
		// 	Parts: []*genai.Part{
		// 		{
		// 			FileData: &genai.FileData{
		// 				FileURI:  "http://ofasys-multimodal-wlcb-3.oss-cn-wulanchabu.aliyuncs.com/yongqi.wyq/muchin_eval_audio_16k/s_0003_src.mp3?OSSAccessKeyId=LTAI5tNSZEaGH5to7YysTUXz&Expires=1749492017&Signature=TIUlEmVrDuLkB9s39%2FwvKcd1lsw%3D",
		// 				MIMEType: "audio/mpeg",
		// 			},
		// 		},
		// 	},
		// 	Role: "user",
		// },
		{
			Parts: []*genai.Part{
				{
					Text: "Say cheerfully: Have a wonderful day!",
				},
			},
			Role: "user",
		},
	}
	ppp.GenerateContentConfig.ResponseModalities = []string{"AUDIO"}
	ppp.GenerateContentConfig.SpeechConfig.MultiSpeakerVoiceConfig = &genai.MultiSpeakerVoiceConfig{
		SpeakerVoiceConfigs: []*genai.SpeakerVoiceConfig{
			{
				Speaker: "Joe",
				VoiceConfig: &genai.VoiceConfig{
					PrebuiltVoiceConfig: &genai.PrebuiltVoiceConfig{
						VoiceName: "Kore",
					},
				},
			},
			// &genai.SpeakerVoiceConfig{
			// 	Speaker: "Jane",
			// 	VoiceConfig: &genai.VoiceConfig{
			// 		PrebuiltVoiceConfig: &genai.PrebuiltVoiceConfig{
			// 			VoiceName: "Puck",
			// 		},
			// 	},
			// },
		},
	}
	a1, err := client.Models.GenerateContent(ctx, ppp.ModelName, ppp.Contents, ppp.GenerateContentConfig)
	if err != nil {
		fmt.Println("aaaa", err)
	}
	fmt.Println(a1.Candidates[0].Content.Parts[0].Text)
	// a, err := impl.NoStreamingGenerateContent(ctx, ppp)
	// if err != nil {
	// 	fmt.Println("cc", err)
	// 	panic(err)
	// }
	// fmt.Println(format.ToJSON(a))
}

func init() {
	// os.Setenv("DEBUG", "true") //this debug conflicts with another vs debug
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	if err := provider.Init(); err != nil {
		fmt.Println("eeer", err)
	}
	// err := registry.ProviderInit() //
	// if err != nil {
	// 	fmt.Println("1234", err)
	// 	// panic(err)
	// }
	impl = provider.GetModelIssuer(vertex.GRANT_MODEL_Gemini_2_5_Flash_Preview)
}
