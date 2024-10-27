package global_data

import (
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/utils/containers"
	"time"
)

var BotActionChannel = make(chan string, define.ChannelBufferSize)

var BotMessageIdStack = containers.NewStackGroup[int](20, 2*time.Minute)

var CustomPlugins = make([]models.PluginInterface, 0)
