package ai

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/netman"
	"microAPro/utils/logger"
)

func ChatMsg(ask string) string {

	req, _ := sonic.MarshalString(Req{
		Model: "ep-20240603053823-2z64f",
		Messages: []MessageMeta{
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
	logger.Info(string(data))
	if err != nil {
		return "出错了" + err.Error()
	}

	return res.Choices[0].Message.Content

}
