package custom_plugin

import (
	"microAPro/channels"
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/calc"
	"microAPro/utils/logger"
	"strings"
)

type RecallSelf struct{}

func (r *RecallSelf) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {
	logger.Info("RecallSelf ContextFilter")
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
			[]string{"r", "recall", "recall_self", "撤回", "撤回自己", "撤回自己的消息"}) {
		return models.ContextFilterResult{}
	}

	// 撤回自己的消息
	messageId, ok := channels.BotMessageIdStack.Pop()

	if ok {
		bot_action.BotActionAPIInstance.RecallMessage(messageId)
		return models.ContextFilterResult{
			BreakFilter: true,
		}
	}
	return models.ContextFilterResult{}
}
