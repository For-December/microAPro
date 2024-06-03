package bot_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/utils/logger"
)

var botEventChannel = make(chan []byte, define.ChannelBufferSize)

func runDispatcher() {
	go func() {
		for {
			select {
			case msg := <-botEventChannel:
				dispatcher(msg)
			}
		}
	}()
}
func dispatcher(msg []byte) {
	event := BotEvent{}
	err := sonic.Unmarshal(msg, &event)
	if err != nil {
		logger.Error("解析事件json失败: ", string(msg), err)
		return
	}

	switch event.PostType {
	case "message":
		logger.Info("message")
	case "notice":
	case "request":
	case "meta_event":
	//events.HandleMateEvent(message)
	default:
		logger.Warning("unknown event type: ", event.PostType)

	}

}
