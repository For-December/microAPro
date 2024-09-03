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

type GroupLog struct {
}

func (g *GroupLog) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {
	logger.DebugF("[group:%d]-[user:%d] => %s", ctx.GroupId, ctx.UserId, ctx.MessageChain.ToString())

	marshalString, err := sonic.MarshalString(ctx.MessageChain.ToJsonTypeMessage())
	if err != nil {
		logger.Error(err)
		return models.ContextFilterResult{}
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
	return models.ContextFilterResult{}
}
