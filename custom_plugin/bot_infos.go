package custom_plugin

import (
	"fmt"
	"microAPro/constant/define"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/calc"
	"strings"
)

type BotInfos struct{}

func (b *BotInfos) GetPluginInfo() string {
	return "BotInfos -> bot相关信息展示"
}

func (b *BotInfos) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {

	if ctx.MessageType != "group" || len(ctx.MessageChain.Messages) != 2 {
		return models.ContextFilterResult{}
	}

	if ctx.MessageChain.Messages[0].MessageType != "at" ||
		ctx.MessageChain.Messages[0].MessageContent["qq"] != define.BotQQ {
		return models.ContextFilterResult{}
	}

	if ctx.MessageChain.Messages[1].MessageType != "text" ||
		!calc.IsTargetInArray[string](
			strings.TrimSpace(ctx.MessageChain.Messages[1].MessageContent["text"].(string)),
			[]string{"f", "func", "fn", "功能清单", "functions", "功能"}) {
		return models.ContextFilterResult{}
	}

	resStr := "功能清单：\n"
	for i, plugin := range global_data.CustomPlugins {
		resStr += fmt.Sprintf("[%d]: %s\n", i+1, plugin.GetPluginInfo())
	}

	bot_action.BotActionAPIInstance.SendGroupMessage(
		*((&models.MessageChain{
			GroupId: ctx.GroupId,
		}).Text(resStr)),
		func(messageId int) {
			global_data.BotMessageIdStack.Push(messageId)
		})

	return models.ContextFilterResult{
		BreakFilter: true,
	}
}
