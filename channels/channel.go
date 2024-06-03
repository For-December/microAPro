package channels

import (
	"microAPro/constant/define"
	"microAPro/models"
)

var MessageContextChannel = make(chan models.MessageContext, 100)

var BotActionChannel = make(chan string, define.ChannelBufferSize)

var AIChannel = make(chan AIAsk, 100)
