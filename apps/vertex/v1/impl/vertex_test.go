package impl_test

import (
	"context"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
	"github.com/kade-chen/mcenter/apps/vertex"
	_ "github.com/kade-chen/mcenter/apps/vertex/v1/impl"
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
	// req.SystemInstruction.Parts = "必须中文回答"
	req.ModelName = "gemini-2.0-flash-001"
	req.Regional = vertex.LOCATION_US_Central1
	prompt := genai.Text("what is images")
	img := genai.FileData{
		MIMEType: "image/jpeg",
		FileURI:  "https://storage.googleapis.com/kadeccc/cat.jpg",
	}
	resp, err := impl.NoStreamingGenerateContent(ctx, req, prompt, img)
	a := format.ToJSON(resp)
	fmt.Println(a, err)
	// if err != nil {
	// 	fmt.Println("--------", err)
	// }
	// rb, err := json.MarshalIndent(resp, "", "  ")
	// // rb, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Println("--------", err)
	// }
	// t.Log(string(rb))
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

const InputText = `### 目标
分析视频/音频中的行为特征与能力表现，通过多维度分析，提取关键时刻，提供优化建议，并生成易于展示的报告。

### 输入
1. **素材背景**：包含完整背景信息、对话及目标特质（如领导力、团队协作等）。
2. **核心问题**：目标情境中的关键点（如压力下决策、资源分配、大局意识）。
3. **分析维度**：按基础维度和延展维度逐一剖析。

### 基础提示词设计
1. **初步分析**：从以下维度提取信息：
   - 性格特质（如果断、冷静）
   - 领导力思维（如团队激励、决策方式）
   - 快速反应能力（应对突发状况）
   - 临场发挥能力（即时判断与应变）
   - 大局意识（整体规划、资源分配）

2. **延展分析**：从以下角度补充优化：
   - 沟通能力（团队交流效率与条理性）
   - 时间管理（紧迫情况下的决策效率）
   - 情绪管理（在压力或不确定性下的表现）
   - 决策影响力（行动结果和效果）

### 输出要求
1. **关键时间戳**：标记相关行为对应的时间节点，便于对照分析（格式：XX:XX）。
2. **能力表现描述**：说明事例中体现的具体能力及行为。
3. **优化建议**：提出可提升空间与建议（如改善沟通策略、增强应变灵活性）。
4. **综合能力评级**：按 1-5 分为每项能力评分，并总结总体表现。`

func TestNoStreamingHeDE(t *testing.T) {
	req := vertex.NewGenaiSetting()
	req.SystemInstruction.Role = ""
	req.SystemInstruction.Parts = InputText
	req.ModelName = "gemini-1.5-flash-002"
	prompt := genai.Text("其中这是四个颜色的角色，帮我按照提示词进行输出")
	img := genai.FileData{
		MIMEType: mime.TypeByExtension(filepath.Ext("team1-c3.mp4")),
		FileURI:  "gs://hede-test-ai/team1-c3.mp4",
	}
	resp, err := impl.NoStreamingGenerateContent(ctx, req, img, prompt)
	if err != nil {
		// fmt.Println("--------", err)
		t.Fatal(err)
	}
	fmt.Println(resp.Candidates[0].Content.Parts)
	// rb, err := json.MarshalIndent(resp, "", "")
	// // rb, err := json.Marshal(resp)
	//
	//	if err != nil {
	//		fmt.Println("--------", err)
	//	}
	//
	// fmt.Println(string(rb))
}

func TestStreamingGenerateContentHEDE(t *testing.T) {
	req := vertex.NewGenaiSetting()
	// req.SystemInstruction.Parts = []genai.Part{
	// 	genai.Text("音频里说了什么"),
	// 	// genai.FileData{},
	// 	// genai.Blob{},
	// }
	req.ModelName = "gemini-1.5-flash-002"
	req.SystemInstruction.Parts = "目标：\n分析视频/音频中的行为特征与能力表现，通过多维度分析，提取关键时刻，提供优化建议，并生成易于展示的报告。\n\n输入：\n\t1.\t素材背景：包含完整背景信息、对话及目标特质（如领导力、团队协作等）。\n\t2.\t核心问题：目标情境中的关键点（如压力下决策、资源分配、大局意识）。\n\t3.\t分析维度：按基础维度和延展维度逐一剖析。\n\n基础提示词设计：\n\n\t1.\t初步分析：从以下维度提取信息：\n\t•\t性格特质（如果断、冷静）\n\t•\t领导力思维（如团队激励、决策方式）\n\t•\t快速反应能力（应对突发状况）\n\t•\t临场发挥能力（即时判断与应变）\n\t•\t大局意识（整体规划、资源分配）\n\t2.\t延展分析：从以下角度补充优化：\n\t•\t沟通能力（团队交流效率与条理性）\n\t•\t时间管理（紧迫情况下的决策效率）\n\t•\t情绪管理（在压力或不确定性下的表现）\n\t•\t决策影响力（行动结果和效果）\n\n\n输出要求：\n\n\t1.\t关键时间戳：标记相关行为对应的时间节点，便于对照分析（格式：XX:XX）。\n\t2.\t能力表现描述：说明事例中体现的具体能力及行为。\n\t3.\t优化建议：提出可提升空间与建议（如改善沟通策略、增强应变灵活性）。\n\t4.\t综合能力评级：按 1-5 分为每项能力评分，并总结总体表现。"

	prompt := genai.Text("其中这是四个颜色的角色，帮我按照提示词进行输出")
	img := genai.FileData{
		MIMEType: "video/mp4",
		FileURI:  "gs://hede-test-ai/team1-c3.mp4",
	}

	GenerateContentResponseIterator := impl.StreamingGenerateContent(ctx, req, img, prompt)
	for {
		resp, err := GenerateContentResponseIterator.Next()
		if err == iterator.Done {
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			fmt.Println("error: ", err)
			exception.NewInternalServerError("error: %v", err)
		}
		if err != nil {
			fmt.Println("error: ", err)
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

func TestNoStreamingGenerateContent1(t *testing.T) {
	// fmt.Println(ioc.Controller().List(), ioc.Config().List(), *new(int32))
	req := vertex.NewGenaiSetting()
	// req.SystemInstruction.Parts = []genai.Part{
	// 	genai.Text("必须中文回答"),
	// 	// genai.FileData{},
	// 	// genai.Blob{},
	// }
	req.SystemInstruction.Parts = "必须中文回答"
	req.ModelName = "gemini-2.0-flash-exp"
	prompt := genai.Text("西红柿炒鸡蛋怎么做")
	// img := genai.FileData{
	// 	MIMEType: "image/jpeg",
	// 	FileURI:  "https://storage.googleapis.com/kadeccc/cat.jpg",
	// }
	resp, err := impl.NoStreamingGenerateContent(ctx, req, prompt)
	fmt.Println(resp, err)
	// if err != nil {
	// 	fmt.Println("--------", err)
	// }
	// rb, err := json.MarshalIndent(resp, "", "  ")
	// // rb, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Println("--------", err)
	// }
	// t.Log(string(rb))
	// impl.GetLocationFromString()
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "../../../../etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName3).(vertex.Service)
}
