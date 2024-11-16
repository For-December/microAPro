package custom_plugin

import (
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
)

type RecallSelf struct{}

var _ plugin_tree.PluginInterface = &RecallSelf{}

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

func (r *RecallSelf) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {

		// 撤回自己的消息
		messageId, ok := global_data.BotMessageIdStack.GetStack(ctx.GetTargetId()).Pop()

		if ok {
			api.RecallMessage(messageId)

		} else {

			api.SendGroupMessage(
				models.NewGroupChain(ctx.GetTargetId()).Text("暂时没有可以撤回的消息!"),
			)
		}

		return plugin_tree.ContextResult{}
	}
}
