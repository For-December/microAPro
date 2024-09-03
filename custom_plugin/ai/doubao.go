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
				Content: "你是一名杭州一中的高二女生，性格温柔但爱憎分明，下面是某个群友的发言，请给出适当回应：",
				//Content: "你是一名杭州一中的高二女生，也是一个有素质的网络喷子，很喜欢用😆😅🤗🤣这几个emoji表情。" +
				//	"现在你在一个群里，下面是某个群友的发言，面对对方的言论，请用阴阳怪气的方式反驳，减少语气词，话语尽可能简洁且直击要害，并加入emoji表情" +
				//	"但当别人想你求爱时，请认真温柔地用真实感情回应他，不要嘲笑和讽刺他",
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

var historyArray = []MessageMeta{
	{
		Role:    "system",
		Content: "你是一名杭州一中的高二女生，性格温柔但爱憎分明，下面是某个群友的发言，请给出适当回应：",

		//Content: "你是一名杭州一中的高二女生，也是一个有素质的网络喷子，很喜欢用😆😅🤗🤣这几个emoji表情。" +
		//	"现在你在一个群里，下面是某个群友的发，面对对方的言论，请用阴阳怪气的方式反驳，减少语气词，话语尽可能简洁且直击要害，并加入emoji表情" +
		//	"但当别人想你求爱时，请认真温柔地用真实感情回应他，不要嘲笑和讽刺他",
	},
}

func ChatMsgWithHistory(ask string) string {

	historyArray = append(historyArray, MessageMeta{
		Role:    "user",
		Content: ask,
	})

	logger.Info(historyArray)

	req, _ := sonic.MarshalString(Req{
		Model:    "ep-20240603053823-2z64f",
		Messages: historyArray,
		Stream:   false,
	})

	data := netman.FastPost(define.DouBaoEndPoint+define.DouBaoChat, req)
	res := Res{}
	err := sonic.Unmarshal(data, &res)
	logger.Info(string(data))
	if err != nil {
		return "出错了" + err.Error()
	}

	historyArray = append(historyArray, MessageMeta{
		Role:    "assistant",
		Content: res.Choices[0].Message.Content,
	})

	return res.Choices[0].Message.Content

}
