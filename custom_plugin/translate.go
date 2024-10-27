package custom_plugin

import (
	"fmt"
	"gorm.io/gorm/utils"
	"microAPro/custom_plugin/translate"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/provider/bot_action"
)

type Translate struct {
}

var _ models.PluginInterface = &Translate{}

func (t *Translate) GetPluginInfo() string {
	return "Translate -> 翻译到指定语言\n${translateTo|tr2 [zh|en|jp] [文本]}"
}

func (t *Translate) GetPaths() []string {
	cmdStrArr := []string{
		"translateTo",
		"tr2",
	}
	res := make([]string, 0)

	for _, str := range cmdStrArr {
		res = append(res,
			fmt.Sprintf("%v $lang **", str))
	}

	return res
}

func (t *Translate) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		// 从消息链中提取文本消息
		lang := ctx.Params["lang"]
		text := ctx.Params["**"]

		var toLangFunc func(string) string

		switch lang {
		case "zh":
			toLangFunc = translate.ToZh
		case "en":
			toLangFunc = translate.ToEn
		case "jp":
			toLangFunc = translate.ToJp
		default:
			toLangFunc = translate.ToZh
		}

		bot_action.BotActionAPIInstance.SendGroupMessage(
			*(&models.MessageChain{
				GroupId: ctx.GroupId,
			}).
				At(utils.ToString(ctx.UserId)).
				Text(" ").Text(toLangFunc(text)),
			func(messageId int) {
				// 将messageId保存
				global_data.BotMessageIdStack.GetStack(ctx.GroupId).Push(messageId)
			})
		return models.ContextResult{}
	}
}
