package custom_plugin

import (
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
)

type ForwardMsg struct{}

var _ plugin_tree.PluginInterface = &ForwardMsg{}

func (f *ForwardMsg) GetScope() uint32 {
	return define.GroupScope
}

func (f *ForwardMsg) GetPluginInfo() string {
	return "ForwardMsg -> 转发消息"
}

func (f *ForwardMsg) GetPaths() []string {
	return []string{"**"}
}

func (f *ForwardMsg) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {

		msg := models.NewGroupChain(902907141)
		msg.Messages = ctx.MessageChain.Messages
		if ctx.GetFromId() == 1921567337 {
			api.SendGroupMessage(msg)
			return plugin_tree.ContextResult{}
		}

		return plugin_tree.ContextResult{
			IsContinue: true,
		}
	}
}
