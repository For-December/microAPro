package custom_plugin

import (
	"fmt"
	"microAPro/ai"
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

type AIChat struct {
}

func (a *AIChat) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {

	logger.Info("AIChat ContextFilter")
	// 下一级处理
	if ctx.MessageType != "group" {
		return models.ContextFilterResult{
			IsContinue: true,
			ErrMsg:     nil,
		}
	}

	fmt.Println(ctx.MessageChain.ToString())
	questionStr := ""
	isAsk := false
	for _, s := range ctx.MessageChain.Messages {

		if s.MessageType == "at" && s.MessageContent["qq"] == define.BotQQ {
			isAsk = true
		}

		if s.MessageType == "text" {
			questionStr += s.MessageContent["text"].(string)
		}

	}
	if !isAsk {
		return models.ContextFilterResult{
			IsContinue: true,
			ErrMsg:     nil,
		}
	}

	logger.InfoF("[%d] -> %s", ctx.GroupId, questionStr)

	bot_action.BotActionAPIInstance.SendGroupMessage(
		*(&models.MessageChain{
			GroupId: ctx.GroupId,
		}).Text(ai.ChatMsg(questionStr)),
		func(messageId int) {
			println("id---------------> ", messageId)
		})

	return models.ContextFilterResult{
		IsContinue: false,
		ErrMsg:     nil,
	}
}
