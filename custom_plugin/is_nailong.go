package custom_plugin

import (
	"encoding/json"
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/netman"
	"microAPro/provider/bot_action"
	"microAPro/utils/logger"
	"strconv"
)

type NaiLongCatcher struct{}

func (n *NaiLongCatcher) GetScope() uint32 {
	return define.GroupScope

}

var _ plugin_tree.PluginInterface = &NaiLongCatcher{}

func (n *NaiLongCatcher) GetPluginInfo() string {
	return "NaiLongCatcher -> 奶龙狩猎者"
}
func (n *NaiLongCatcher) GetPaths() []string {
	return []string{
		"# $image",
	}
}

func (n *NaiLongCatcher) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
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
			groupId := ctx.MessageChain.GetTargetId()
			if string(res) == "true" {
				api.SendGroupMessage(
					models.NewGroupChain(groupId).
						Reply(strconv.Itoa(int(ctx.MessageId))).Text("别发奶龙！"))
			}
		}()
		return plugin_tree.ContextResult{}
	}
}
