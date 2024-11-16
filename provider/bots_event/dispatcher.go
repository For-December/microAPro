package bots_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/custom_plugin"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/utils/containers"
	"microAPro/utils/logger"
)

func init() {

}

var trie *containers.RouteTrie

// 注册插件
func registerCustomPlugins() {
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.BotInfos{})
	//global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.VoiceReply{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.RecallSelf{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.AIChat{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.GroupLogs{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.Echo{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.Translate{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.NaiLongCatcher{})
	global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.ColorPic{})

	// 树形路由匹配注册
	trie = containers.NewRouteTrie(models.CallbackFunc{})

	for _, plugin := range global_data.CustomPlugins {
		paths := plugin.GetPaths()
		for _, path := range paths {
			// 优先级 !! > 精确匹配 > **
			trie.Insert(path, plugin.GetPluginHandler())
		}

	}
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

	// 改成树形路由匹配
	trie.SearchAndExec(ctx)
}

func processMsg(msg message, msgType string, targetId, fromId int) {
	messageChain := &models.MessageChain{
		fromId:   groupMessage.UserId,
		TargetId: groupMessage.GroupId,
	}
	models.MessageChainBuilder.NewReceivedChain()
	for _, s := range gpMessage {

		switch s.Type {
		case "text":
			messageChain.Text(s.Data["text"].(string))
		case "image":
			messageChain.Image(s.Data["file"].(string))
		case "record":
			messageChain.Record(s.Data["file"].(string))
		case "at":
			messageChain.At(s.Data["qq"].(string))
		case "reply":
			messageChain.Reply(s.Data["id"].(string))
		case "face":
			messageChain.Face(s.Data["id"].(string))

		default:
			logger.Warning("no such message type: ", s.Type)
			continue
		}
	}

	executePlugins(&models.MessageContext{
		BotAccount:   define.BotQQ,
		MessageType:  msgType,
		MessageId:    groupMessage.MessageId,
		GroupId:      groupMessage.GroupId,
		UserId:       groupMessage.UserId,
		MessageChain: messageChain,
	})
}
func dispatcher(msg []byte) {
	event := botEvent{}
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
			groupMessage := groupMessageEvent{}
			if err := sonic.Unmarshal(msg, &groupMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}
			//logger.Info(string(msg))

			// 消息链

			gpMessage := groupMessage.Message

		case "private":
			privateMessage := privateMessageEvent{}
			if err := sonic.Unmarshal(msg, &privateMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}

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
