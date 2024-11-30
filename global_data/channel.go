package global_data

import (
	"microAPro/constant/config"
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/utils/containers"
	"time"
)

var BotMessageIdStack = containers.NewStackGroup[int64](20, 2*time.Minute)

// GroupChannels 用于存储群聊的消息
// 这里的map不会在运行时扩展，所以用Map 而非 sync.Map
var GroupChannels = make(map[int64]chan *models.MessageContext)

func init() {
	for _, grp := range config.EnvCfg.GroupWhitelist {
		GroupChannels[grp] = make(chan *models.MessageContext, define.ChannelBufferSize)
	}

	// 机器人账号
	for _, account := range config.EnvCfg.BotAccounts {
		GroupChannels[account] = make(chan *models.MessageContext, define.ChannelBufferSize)
	}

}

func GetNextContext(groupId int64) *models.MessageContext {
	select {
	case ctx := <-GroupChannels[groupId]:
		return ctx
	}

}
