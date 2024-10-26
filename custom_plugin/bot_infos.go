package custom_plugin

import (
	"fmt"
	"microAPro/constant/define"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
)

type BotInfos struct{}

var _ models.PluginInterface = &BotInfos{}

func (b *BotInfos) GetPluginInfo() string {
	return "BotInfos -> bot相关信息展示"
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
func (b *BotInfos) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
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
		return models.ContextResult{}
	}
}
