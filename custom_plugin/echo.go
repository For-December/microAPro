package custom_plugin

import (
	"fmt"
	"gorm.io/gorm/utils"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
)

type Echo struct {
}

var _ models.PluginInterface = &Echo{}

func (e *Echo) GetPluginInfo() string {
	return "Echo -> 回显\n${echo|ec [文本 | 图片 | @群友 文本]}"
}

func (e *Echo) GetPaths() []string {
	cmdStrArr := []string{
		"echo",
		"ec",
	}
	res := make([]string, 0)

	for _, str := range cmdStrArr {
		res = append(res,
			fmt.Sprintf("%v $text", str))

		res = append(res,
			fmt.Sprintf("%v @ $qq $text", str))

		res = append(res,
			fmt.Sprintf("%v # $image", str))
	}

	return res
}

func (e *Echo) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		// 从消息链中提取文本消息
		text := ctx.Params["text"]
		qq := ctx.Params["qq"]
		image := ctx.Params["image"]
		if image != "" {
			bot_action.BotActionAPIInstance.SendGroupMessage(
				*(&models.MessageChain{
					GroupId: ctx.GroupId,
				}).At(utils.ToString(ctx.UserId)).Image(image),
				func(messageId int) {
					// 将messageId保存
					global_data.BotMessageIdStack.GetStack(ctx.GroupId).Push(messageId)
				})
			return models.ContextResult{}
		}
		// 发送文本消息
		if qq != "" {
			bot_action.BotActionAPIInstance.SendGroupMessage(
				*(&models.MessageChain{
					GroupId: ctx.GroupId,
				}).At(qq).Text(" ").Text(text),
				func(messageId int) {
					// 将messageId保存
					global_data.BotMessageIdStack.GetStack(ctx.GroupId).Push(messageId)
				})
			return models.ContextResult{}
		}

		bot_action.BotActionAPIInstance.SendGroupMessage(
			*(&models.MessageChain{
				GroupId: ctx.GroupId,
			}).At(utils.ToString(ctx.UserId)).Text(" ").Text(text),
			func(messageId int) {
				// 将messageId保存
				global_data.BotMessageIdStack.GetStack(ctx.GroupId).Push(messageId)
			})
		return models.ContextResult{}
	}
}
