package custom_plugin

import (
	"encoding/json"
	"microAPro/models"
	"microAPro/netman"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
	"strconv"
)

type NaiLongCatcher struct{}

var _ models.PluginInterface = &NaiLongCatcher{}

func (n *NaiLongCatcher) GetPluginInfo() string {
	return "NaiLongCatcher -> 奶龙狩猎者"
}
func (n *NaiLongCatcher) GetPaths() []string {
	return []string{
		"# $image",
	}
}

func (n *NaiLongCatcher) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		logger.Info("NaiLongCatcher")
		go func() {
			imageUrl := ctx.Params["image"]

			marshal, err := json.Marshal(struct {
				Url string `json:"url"`
			}{
				Url: imageUrl,
			})
			if err != nil {
				logger.Error(err)
				return
			}

			res := netman.FastPost("http://127.0.0.1:8000/long", string(marshal))
			if string(res) == "true" {
				bot_action.BotActionAPIInstance.SendGroupMessage(
					*((&models.MessageChain{
						GroupId: ctx.GroupId,
					}).Reply(strconv.Itoa(ctx.MessageId)).Text("别发奶龙！")),
					func(messageId int) {
					})
			}
		}()
		return models.ContextResult{}
	}
}
