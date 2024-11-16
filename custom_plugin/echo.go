package custom_plugin

import (
	"fmt"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
)

type Echo struct {
}

var _ plugin_tree.PluginInterface = &Echo{}

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

func (e *Echo) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
		// 从消息链中提取文本消息
		text := ctx.Params["text"]
		qq := ctx.Params["qq"]
		image := ctx.Params["image"]

		groupId := ctx.MessageChain.GetTargetId()

		messageChain := models.NewGroupChain(groupId)

		if qq != "" {
			messageChain.At(qq)
		}

		if image != "" {
			messageChain.Image(image)
		}

		if text != "" {
			messageChain.Text(text)
		}

		api.SendGroupMessage(
			messageChain,
			func(messageId int64) {
				// 将messageId保存
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
			})
		return plugin_tree.ContextResult{}
	}
}
