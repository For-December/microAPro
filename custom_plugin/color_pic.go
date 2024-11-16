package custom_plugin

import (
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

type ColorPic struct{}

var _ plugin_tree.PluginInterface = &ColorPic{}

func (c *ColorPic) GetPluginInfo() string {
	return "ColorPic -> 随机色图\n${st|color_pic|涩图}"
}
func (c *ColorPic) GetPaths() []string {

	return []string{
		"st",
		"color_pic",
		"涩图",
	}
}
func (c *ColorPic) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
		logger.Info("ColorPic -> 随机色图")
		groupId := ctx.MessageChain.GetTargetId()
		api.SendGroupMessage(
			models.NewGroupChain(groupId).Image("https://image.anosu.top/pixiv"),
			func(messageId int64) {
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
			})

		return plugin_tree.ContextResult{}
	}
}
