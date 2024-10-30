package custom_plugin

import (
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

type ColorPic struct{}

var _ models.PluginInterface = &ColorPic{}

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
func (c *ColorPic) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		logger.Info("ColorPic -> 随机色图")
		bot_action.BotActionAPIInstance.SendGroupMessage(
			*((&models.MessageChain{
				GroupId: ctx.GroupId,
			}).Image("https://image.anosu.top/pixiv")),
			func(messageId int) {
				global_data.BotMessageIdStack.GetStack(ctx.GroupId).Push(messageId)
			})

		return models.ContextResult{}
	}
}
