package custom_plugin

import (
	"fmt"
	"gorm.io/gorm/utils"
	"microAPro/constant/define"
	"microAPro/custom_plugin/translate"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
)

type Translate struct {
}

func (t *Translate) GetScope() uint32 {
	return define.GroupScope

}

var _ plugin_tree.PluginInterface = &Translate{}

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

func (t *Translate) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
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

		api.SendGroupMessage(
			models.NewGroupChain(ctx.GetTargetId()).
				At(utils.ToString(ctx.GetFromId())).
				Text(" ").Text(toLangFunc(text)),
			func(messageId int64) {
				// 将messageId保存
				global_data.BotMessageIdStack.GetStack(ctx.GetTargetId()).Push(messageId)
			})
		return plugin_tree.ContextResult{}
	}
}
