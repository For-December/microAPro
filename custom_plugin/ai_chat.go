package custom_plugin

import (
	"gorm.io/gorm/utils"
	"microAPro/constant/define"
	"microAPro/custom_plugin/ai"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

var _ plugin_tree.PluginInterface = &AIChat{}

type AIChat struct{}

func (a *AIChat) GetScope() uint32 {
	return define.GroupScope
}

func (a *AIChat) GetPluginInfo() string {
	return "AIChat -> 内置ai"
}

func (a *AIChat) GetPaths() []string {
	return []string{
		"@ " + define.BotQQ + " **",
	}
}

var DevMode = false

func (a *AIChat) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {

		groupId := ctx.MessageChain.GetTargetId()
		userId := ctx.MessageChain.GetFromId()

		questionStr := ctx.Params["**"]

		//logger.InfoF("[%d] -> %s", ctx.GroupId, questionStr)

		//prompt := "你是一名杭州一中的高二女生，性格温柔但爱憎分明，下面是某个群友的发言，请给出适当回应："
		prompt := ""

		if DevMode {
			logger.Debug("启用开发者模式")
			prompt = "你是一名杭州一中的高二女生，也是一个有素质的网络喷子，很喜欢用😆😅🤗🤣这几个emoji表情。" +
				"现在你在一个群里，下面是某个群友的发，面对对方的言论，请用阴阳怪气的方式反驳，减少语气词，话语尽可能简洁且直击要害，并加入emoji表情" +
				"但当别人想你求爱时，请认真温柔地用真实感情回应他，不要嘲笑和讽刺他。" +
				"如果有人问你为什么在杭州一中上学等类似问题时，请围绕【可能某些人对杭州的女孩子存在不切实际的性幻想吧】来酌情适当回复"
		} else {
			logger.Debug("禁用开发者模式")
			prompt = "你是一名杭州一中的高二女生，性格温柔但爱憎分明，下面是某个群友的发言，请给出适当回应："
		}

		answerStr := "默认回复"
		if questionStr == "devOn" {
			DevMode = true
			answerStr = "AI开发者模式已开启"
		} else if questionStr == "devOff" {
			DevMode = false
			answerStr = "AI开发者模式已关闭"
		} else {
			answerStr = ai.ChatMsgWithHistory(groupId, questionStr, prompt)
		}

		api.SendGroupMessage(
			models.NewGroupChain(groupId).
				At(utils.ToString(userId)).Text(" ").Text(answerStr),
			func(messageId int64) {
				// 将messageId保存
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
				println("id---------------> ", messageId)
			})

		return plugin_tree.ContextResult{}
	}
}
