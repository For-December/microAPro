package custom_plugin

import (
	"microAPro/constant/define"
	"microAPro/custom_plugin/img"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
)

type Img2Img struct{}

var _ plugin_tree.PluginInterface = &Img2Img{}

func (i *Img2Img) GetPaths() []string {
	prefix := "@ " + define.BotQQ + " "
	return []string{
		prefix + "ii",
		prefix + "i2i",
		prefix + "img2img",
		prefix + "生图",
	}
}

func (i *Img2Img) GetPluginInfo() string {

	return "Img2Img -> 根据图片生成图片\n$@小A {ii | i2i | img2img | 生图}"
}

func (i *Img2Img) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
		groupId := ctx.MessageChain.GetTargetId()
		// 这里注册一个服务，等待用户发送图片
		// 服务的回调函数会将图片转发给一个图片处理服务，然后将处理后的图片发送给用户
		api.SendGroupMessage(
			models.NewGroupChain(groupId).Text("请发送图片"),
			func(messageId int64) {
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
			})

		logger.Info("??")

		// 获取该群下一个上下文
		content := global_data.GetNextContext(groupId)
		logger.Info(content)

		for _, message := range content.MessageChain.Messages {
			if message.MessageType == "image" {
				img.BDImg2ImgInChannel <- message.MessageContent["file"].(string)
				go func() {
					select {
					case imgPath := <-img.BDImg2ImgOutChannel:
						api.SendGroupMessage(
							models.NewGroupChain(groupId).Image(imgPath),
							func(messageId int64) {
								global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
							})
					}
				}()
				return plugin_tree.ContextResult{}
			}

			api.SendGroupMessage(
				models.NewGroupChain(groupId).Text("发的并不是图片，请重新启动"),
				func(messageId int64) {
					global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
				})
		}
		return plugin_tree.ContextResult{}
	}
}

func (i *Img2Img) GetScope() uint32 {
	return define.GroupScope
}
