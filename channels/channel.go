package channels

import (
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/utils/containers"
	"time"
)

var ReceivedMessageContextChannel = make(chan models.MessageContext, 100)

var BotActionChannel = make(chan string, define.ChannelBufferSize)

var AIChannel = make(chan AIAsk, 100)

var BotMessageIdStack = containers.NewCustomStack[int](20, 2*time.Minute)
