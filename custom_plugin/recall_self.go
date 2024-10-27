package custom_plugin

import (
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
)

type RecallSelf struct{}

var _ models.PluginInterface = &RecallSelf{}

func (r *RecallSelf) GetPluginInfo() string {
	return "RecallSelf -> 撤回当前群小A的消息\n${r|recall|recall_self|撤回}"
}

func (r *RecallSelf) GetPaths() []string {
	return []string{
		"r",
		"recall",
		"recall_self",
		"撤回",
	}
}

func (r *RecallSelf) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		// 撤回自己的消息
		messageId, ok := global_data.BotMessageIdStack.GetStack(ctx.GroupId).Pop()

		if ok {
			bot_action.BotActionAPIInstance.RecallMessage(messageId)

		} else {

			bot_action.BotActionAPIInstance.SendGroupMessage(
				*((&models.MessageChain{
					GroupId: ctx.GroupId,
				}).Text("暂时没有可以撤回的消息!")),
				func(messageId int) {
				})
		}

		return models.ContextResult{}
	}
}
