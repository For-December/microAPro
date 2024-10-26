package custom_plugin

import (
	"github.com/bytedance/sonic"
	"gorm.io/gorm/utils"
	"microAPro/dbmodels"
	"microAPro/models"
	"microAPro/storage/database"
	"microAPro/utils/logger"
	"time"
)

type GroupLogs struct{}

var _ models.PluginInterface = &GroupLogs{}

func (g *GroupLogs) GetPluginInfo() string {
	return "GroupLogs -> 记录和群聊相关的运行日志"
}
func (g *GroupLogs) GetPaths() []string {
	return []string{
		"**",
	}
}

func (g *GroupLogs) GetPluginHandler() models.PluginHandler {
	return func(ctx *models.MessageContext) models.ContextResult {
		logger.DebugF("[group:%d]-[user:%d] => %s", ctx.GroupId, ctx.UserId, ctx.MessageChain.ToString())

		marshalString, err := sonic.MarshalString(ctx.MessageChain.ToJsonTypeMessage())
		if err != nil {
			logger.Error(err)
			return models.ContextResult{}
		}

		if err := database.Client.Save(&dbmodels.GroupLog{
			ID:        0,
			GroupID:   utils.ToString(ctx.GroupId),
			UserID:    utils.ToString(ctx.UserId),
			Message:   marshalString,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}).Error; err != nil {
			logger.Error(err)
		}
		return models.ContextResult{}
	}
}
