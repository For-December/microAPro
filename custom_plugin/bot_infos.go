package custom_plugin

import (
	"fmt"
	"microAPro/constant/define"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
)

type BotInfos struct{}

func (b *BotInfos) GetScope() uint32 {
	return define.GroupScope
}

var _ plugin_tree.PluginInterface = &BotInfos{}

func (b *BotInfos) GetPluginInfo() string {
	return "BotInfos -> bot相关信息展示\n${func | fn | f | 功能清单 | functions | 功能}"
}
func (b *BotInfos) GetPaths() []string {

	prefix := "@ " + define.BotQQ + " "
	return []string{
		prefix + "f",
		prefix + "func",
		prefix + "fn",
		prefix + "功能清单",
		prefix + "functions",
		prefix + "功能",
	}
}
func (b *BotInfos) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
		resStr := "功能清单：\n"
		for i, plugin := range plugin_tree.CustomPlugins {
			resStr += fmt.Sprintf("[%d]: %s\n\n", i+1, plugin.GetPluginInfo())
		}

		groupId := ctx.MessageChain.GetTargetId()
		api.SendGroupMessage(
			models.NewGroupChain(groupId).Text(resStr),
			func(messageId int64) {
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
			})
		return plugin_tree.ContextResult{}
	}
}
