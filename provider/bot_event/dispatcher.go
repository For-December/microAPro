package bot_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/custom_plugin"
	"microAPro/models"
	"microAPro/models/entity"
	"microAPro/utils/logger"
)

var botEventChannel = make(chan []byte, define.ChannelBufferSize)

var customPlugins = make([]models.PluginBase, 0)

// 注册插件
func registerCustomPlugins() {
	customPlugins = append(customPlugins, &custom_plugin.AIChat{})
	customPlugins = append(customPlugins, &custom_plugin.GroupLog{})
}

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

func executePlugins(ctx *models.MessageContext) {
	for _, plugin := range customPlugins {

		result := plugin.ContextFilter(ctx)

		if result.ErrMsg != nil {
			logger.Warning("plugin.ContextFilter err: ", result.ErrMsg)
		}

		if result.BreakFilter {
			return
		}
	}
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

		// 群组消息
		case "group":
			groupMessage := GroupMessageEvent{}
			if err := sonic.Unmarshal(msg, &groupMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}
			logger.Info(string(msg))

			// 消息链
			messageChain := &models.MessageChain{
				Messages: make([]entity.CommonMessage, 0),
				FromId:   groupMessage.UserId,
				TargetId: groupMessage.GroupId,
			}

			for _, s := range groupMessage.Message {

				switch s.Type {
				case "text":
					messageChain.Messages = append(messageChain.Messages,
						entity.CommonMessage{
							MessageType: s.Type,
							MessageContent: map[string]interface{}{
								"text": s.Data.Text,
							},
						})
				case "image":
					messageChain.Messages = append(messageChain.Messages,
						entity.CommonMessage{
							MessageType: s.Type,
							MessageContent: map[string]interface{}{
								"url": s.Data.Url,
							},
						})
				case "at":
					messageChain.Messages = append(messageChain.Messages,
						entity.CommonMessage{
							MessageType: s.Type,
							MessageContent: map[string]interface{}{
								"qq": s.Data.QQ,
							},
						})
				default:
					logger.Warning("no such message type: ", s.Type)
					continue
				}
			}

			executePlugins(&models.MessageContext{
				BotAccount:   define.BotQQ,
				MessageType:  "group",
				MessageId:    groupMessage.MessageId,
				GroupId:      groupMessage.GroupId,
				UserId:       groupMessage.UserId,
				MessageChain: messageChain,
			})

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
