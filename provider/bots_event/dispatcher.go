package bots_event

import (
	"github.com/bytedance/sonic"
	"microAPro/constant/config"
	"microAPro/constant/define"
	"microAPro/custom_plugin"
	"microAPro/global_data"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/provider/bot_action"
	"microAPro/utils/calc"
	"microAPro/utils/containers"
	"microAPro/utils/logger"
	"sync"
)

func init() {

}

var groupTrie *containers.RouteTrie
var privateTrie *containers.RouteTrie

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
	plugin_tree.CustomPlugins = append(plugin_tree.CustomPlugins, &custom_plugin.Img2Img{})

	// 树形路由匹配注册
	groupTrie = containers.NewRouteTrie(plugin_tree.CallbackFunc{})
	privateTrie = containers.NewRouteTrie(plugin_tree.CallbackFunc{})

	for _, plugin := range plugin_tree.CustomPlugins {
		paths := plugin.GetPaths()
		for _, path := range paths {
			// 优先级 !! > 精确匹配 > **
			if plugin.GetScope()&define.GroupScope != 0 {
				groupTrie.Insert(path, plugin.GetPluginHandler())
			}
			if plugin.GetScope()&define.PrivateScope != 0 {
				privateTrie.Insert(path, plugin.GetPluginHandler())
			}
		}

	}
}

func runDispatcher(wg *sync.WaitGroup) {
	// 多个调度器处理对应channel数据，每个bot一个channel
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

	// 群聊消息的channel，bots作为集群，每个群聊一个channel
	for _, grp := range config.EnvCfg.GroupWhitelist {
		go func(group int64) {
			wg.Add(1)
			defer wg.Done()
			for {
				ctx := global_data.GetNextContext(group)
				executePlugins(bot_action.NewBotActionAPI(3090807650), ctx)
			}
		}(grp)
	}

}

func executePlugins(api *bot_action.BotActionAPI, ctx *models.MessageContext) {
	//

	// 改成树形路由匹配

	switch ctx.MessageType {
	case define.GroupMsg:
		groupTrie.SearchAndExec(api, ctx)
	case define.PrivateMsg:
		privateTrie.SearchAndExec(api, ctx)
	case define.TempMsg:
	default:
		logger.Warning("unknown message type: ", ctx.MessageType)

	}
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

	if _, ok := global_data.GroupChannels[targetId]; !ok {
		panic("group channel not found!")
	}

	global_data.GroupChannels[targetId] <- &models.MessageContext{
		BotAccount:   botAccount,
		MessageType:  msgType,
		MessageId:    messageId,
		MessageChain: messageChain,
	}

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

			if !calc.IsTargetInArray(groupMessage.GroupId,
				config.EnvCfg.GroupWhitelist) {
				// 该群不在白名单中
				return
			}

			logger.Info("群消息👇")

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
				privateMessage.SelfId, // targetId，这里目标是bot本身
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
