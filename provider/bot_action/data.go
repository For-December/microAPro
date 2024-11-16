package bot_action

import (
	"github.com/lxzan/gws"
	"microAPro/constant/define"
)

var clients = make(map[int64]*gws.Conn)

var botActionResChannel = make(chan []byte, 1000)

var BotActionChannel = make(chan BotAction, define.ChannelBufferSize)
