package bots_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/define"
	"microAPro/custom_plugin"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/utils/containers"
	"microAPro/utils/logger"
	"sync"
)

func init() {

}

var trie *containers.RouteTrie

// 注册插件，将处理函数放到基数树中
func registerCustomPlugins() {
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.BotInfos{})
	//global_data.CustomPlugins = append(global_data.CustomPlugins, &custom_plugin.VoiceReply{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.RecallSelf{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.AIChat{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.GroupLogs{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.Echo{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.Translate{})
	//plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.NaiLongCatcher{})
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.ColorPic{})

	// 树形路由匹配注册
	trie = containers.NewRouteTrie(plugin_tree.CallbackFunc{})

	for _, plugin := range plugin_tree.CustomPlugins {
		paths := plugin.GetPaths()
		for _, path := range paths {
			// 优先级 !! > 精确匹配 > **
			trie.Insert(path, plugin.GetPluginHandler())
		}

	}
}

func runDispatcher(wg *sync.WaitGroup) {
	// 多个调度器处理对应channel数据
	for _, channel := range botsEventChannels {
		// 通道传参，实际上是指针传递，因为通道本事是指针
		go func(ch chan []byte) {
			wg.Add(1)
			defer wg.Done()
			for {
				select {
				case msg := <-ch:
					// msg有个self id字段，因此可以不必传botAccount
					dispatcher(msg)
				}
			}
		}(channel)
	}

}

func executePlugins(api *bot_action.BotActionAPI, ctx *models.MessageContext) {

	// 改成树形路由匹配
	trie.SearchAndExec(api, ctx)
}

func processMsg(msg []message,
	msgType define.MessageType,
	fromId, targetId, messageId int64,
	botAccount int64) {

	messageChain := models.NewReceivedChain(fromId, targetId)

	for _, s := range msg {

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

	executePlugins(bot_action.NewBotActionAPI(botAccount),
		&models.MessageContext{
			BotAccount:   botAccount,
			MessageType:  msgType,
			MessageId:    messageId,
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
		//logger.Info("message")
		switch event.MessageType {

		// 群组消息
		case define.GroupMsg:
			groupMessage := groupMessageEvent{}
			if err := sonic.Unmarshal(msg, &groupMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}
			logger.Info("群消息")

			// 消息链
			processMsg(
				groupMessage.Message,
				define.GroupMsg,
				groupMessage.UserId,
				groupMessage.GroupId,
				groupMessage.MessageId,
				groupMessage.SelfId,
			)

		case define.PrivateMsg:
			privateMessage := privateMessageEvent{}
			if err := sonic.Unmarshal(msg, &privateMessage); err != nil {
				logger.Error("groupMessage err: ", string(msg), err)
				return
			}

			logger.Info("私信")
			// 消息链

			switch privateMessage.SubType {
			case "friend":
				// 好友对话

			case "group":
				// 临时会话

			}

			processMsg(
				privateMessage.Message,
				define.PrivateMsg,
				privateMessage.UserId,
				0, // targetId
				privateMessage.MessageId,
				privateMessage.SelfId,
			)

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
