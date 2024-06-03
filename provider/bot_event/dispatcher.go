package bot_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
)

var botEventChannel = make(chan []byte, define.ChannelBufferSize)

func EventDispatcher() {
	for {
		select {
		case msg := <-botEventChannel:
			event := BotEvent{}
			err := sonic.Unmarshal(msg, &event)
			if err != nil {

				return
			}
		}
	}
}
