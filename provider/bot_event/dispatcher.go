package bot_event

import (
	"fmt"
	"github.com/bytedance/sonic"
	"microAPro/channels"
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
		switch event.MessageType {
		case "group":
			groupMessage := GroupMessageEvent{}
			if err := sonic.Unmarshal(msg, &groupMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}

			fmt.Println(groupMessage)
			questionStr := ""
			isAsk := false
			for i, s := range groupMessage.Message {
				if i == 0 {
					if s.Data.QQ == define.BotQQ {
						isAsk = true
					} else {
						isAsk = false
						break
					}
				}

				questionStr += s.Data.Text

			}
			if !isAsk {
				break
			}
			logger.Debug(questionStr)
			channels.AIChannel <- channels.AIAsk{
				AskerId:  groupMessage.UserId,
				GroupId:  groupMessage.GroupId,
				Question: questionStr,
			}
			//channels.MessageContextChannel <- models.MessageContext{
			//	BotAccount:   "",
			//	MessageType:  "group",
			//	GroupId:      groupMessage.GroupId,
			//	UserId:       groupMessage.UserId,
			//	MessageChain: (&models.MessageChain{}).Text("1"),
			//}
		default:
			logger.Warning("unknown message type: ", event.MessageType)

		}

	case "notice":
	case "request":
	case "meta_event":
	//events.HandleMateEvent(message)
	default:
		logger.Warning("unknown event type: ", event.PostType)

	}

}
