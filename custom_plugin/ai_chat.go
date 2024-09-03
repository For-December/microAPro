package custom_plugin

import (
	"gorm.io/gorm/utils"
	"microAPro/constant/define"
	"microAPro/custom_plugin/ai"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

type AIChat struct{}

func (a *AIChat) GetPluginInfo() string {
	return "AIChat -> 内置ai"
}

func (a *AIChat) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {

	// 下一级处理
	if ctx.MessageType != "group" {
		return models.ContextFilterResult{}
	}

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
		return models.ContextFilterResult{}
	}
	logger.Info("AIChat ContextFilter")

	logger.InfoF("[%d] -> %s", ctx.GroupId, questionStr)

	bot_action.BotActionAPIInstance.SendGroupMessage(
		*(&models.MessageChain{
			GroupId: ctx.GroupId,
		}).At(utils.ToString(ctx.UserId)).Text(" ").Text(ai.ChatMsgWithHistory(questionStr)),
		func(messageId int) {
			// 将messageId保存
			global_data.BotMessageIdStack.Push(messageId)
			println("id---------------> ", messageId)
		})

	return models.ContextFilterResult{
		BreakFilter: false,
		ErrMsg:      nil,
	}
}
