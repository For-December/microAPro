package custom_plugin

import (
	"microAPro/models"
	"microAPro/utils/logger"
)

type GroupLog struct {
}

func (g *GroupLog) ContextFilter(
	ctx *models.MessageContext,
) models.ContextFilterResult {
	logger.DebugF("[group:%d]-[user:%d] => %s", ctx.GroupId, ctx.UserId, ctx.MessageChain.ToString())

	return models.ContextFilterResult{}
}
