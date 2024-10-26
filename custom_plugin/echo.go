package custom_plugin

import (
	"fmt"
	"gorm.io/gorm/utils"
	"microAPro/constant/define"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
)

type Echo struct {
}

var _ models.PluginInterface = &Echo{}

func (e *Echo) GetPluginInfo() string {
	return "Echo -> 回显\n${echo [文本]}"
}

func (e *Echo) GetPaths() []string {
	cmdStrArr := []string{
		"echo",
		"ec",
		"ech",
		"回显",
		"回应",
		"复读",
	}
	res := make([]string, 0)

	for _, str := range cmdStrArr {
		res = append(res,
			fmt.Sprintf("@ %v %v $text", define.BotQQ, str))
	}

	return res
}

func (e *Echo) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		// 从消息链中提取文本消息
		text := ctx.Params["text"]
		// 发送文本消息
		bot_action.BotActionAPIInstance.SendGroupMessage(
			*(&models.MessageChain{
				GroupId: ctx.GroupId,
			}).At(utils.ToString(ctx.UserId)).Text(" ").Text(text),
			func(messageId int) {
				// 将messageId保存
				global_data.BotMessageIdStack.Push(messageId)
			})
		return models.ContextResult{}
	}
}
