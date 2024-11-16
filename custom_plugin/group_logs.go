package custom_plugin

import (
	"github.com/bytedance/sonic"
	"gorm.io/gorm/utils"
	"microAPro/dbmodels"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/storage/database"
	"microAPro/utils/logger"
	"time"
)

type GroupLogs struct{}

var _ plugin_tree.PluginInterface = &GroupLogs{}

func (g *GroupLogs) GetPluginInfo() string {
	return "GroupLogs -> 记录和群聊相关的运行日志"
}
func (g *GroupLogs) GetPaths() []string {
	return []string{
		"!!",
	}
}

func (g *GroupLogs) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {

		groupId := ctx.MessageChain.GetTargetId()
		userId := ctx.MessageChain.GetFromId()
		logger.DebugF("[group:%d]-[user:%d] => %s",
			groupId, userId, ctx.MessageChain.ToString())

		marshalString, err := sonic.MarshalString(ctx.MessageChain.ToJsonTypeMessage())
		if err != nil {
			logger.Error(err)
			return plugin_tree.ContextResult{}
		}

		if err := database.Client.Save(&dbmodels.GroupLog{
			ID:        0,
			GroupID:   utils.ToString(groupId),
			UserID:    utils.ToString(groupId),
			Message:   marshalString,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}).Error; err != nil {
			logger.Error(err)
		}
		return plugin_tree.ContextResult{}
	}
}
