package bots_event

import (
	"github.com/lxzan/gws"
	"microAPro/constant/config"
)

var clients = make(map[int64]*gws.Conn)

var botsEventChannels = make(map[int64]chan []byte)

func init() {
	for _, account := range config.EnvCfg.BotAccounts {
		botsEventChannels[account] = make(chan []byte, 100)
	}
}
