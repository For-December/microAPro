package ai

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/netman"
)

type T2 struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type T struct {
	Model    string `json:"model"`
	Messages []T2   `json:"messages"`
	Stream   bool   `json:"stream"`
}
type Res struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Content interface{} `json:"content"`
		} `json:"logprobs"`
		Message struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int    `json:"created"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func ChatMsg(ask string) string {

	req, _ := sonic.MarshalString(T{
		Model: "ep-20240603053823-2z64f",
		Messages: []T2{
			{
				Role:    "system",
				Content: "你是一名杭州一中的高二女生，现在在一个群里，下面是某个群友的发言，请给出适当回应：",
			},
			{
				Role:    "user",
				Content: ask,
			},
		},
		Stream: false,
	})

	data := netman.FastPost(define.DouBaoEndPoint+define.DouBaoChat, req)
	res := Res{}
	err := sonic.Unmarshal(data, &res)
	//logger.Error(string(data))
	if err != nil {
		return "出错了" + err.Error()
	}

	return res.Choices[0].Message.Content

}
