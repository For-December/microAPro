package global_data

import (
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/provider/bot_action"
	"microAPro/utils/containers"
	"time"
)

var BotActionChannel = make(chan bot_action.BotAction, define.ChannelBufferSize)

var BotMessageIdStack = containers.NewStackGroup[int](20, 2*time.Minute)

var CustomPlugins = make([]models.PluginInterface, 0)
