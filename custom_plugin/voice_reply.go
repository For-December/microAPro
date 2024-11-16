package custom_plugin

import (
	"fmt"
	"microAPro/constant/define"
	"microAPro/models"
	"microAPro/models/plugin_tree"
	"microAPro/utils/logger"
	"net"
	"strings"
)

type VoiceReply struct{}

func (v *VoiceReply) GetPluginInfo() string {
	return "VoiceReply -> 语音回复"
}

func (v *VoiceReply) ContextFilter(
	ctx *models.MessageContext,
) plugin_tree.ContextFilterResult {
	logger.Info("VoiceReply ContextFilter")
	if ctx.MessageType != "group" || len(ctx.MessageChain.Messages) != 2 {
		return plugin_tree.ContextFilterResult{}
	}

	if ctx.MessageChain.Messages[0].MessageType != "at" ||
		ctx.MessageChain.Messages[0].MessageContent["qq"] != define.BotQQ {
		return plugin_tree.ContextFilterResult{}
	}

	if ctx.MessageChain.Messages[1].MessageType != "text" {
		return plugin_tree.ContextFilterResult{}
	}

	splits := strings.Split(
		strings.TrimSpace(ctx.MessageChain.Messages[1].MessageContent["text"].(string)),
		" ")

	if len(splits) != 2 || splits[0] != "语音" {
		logger.Info(splits)
		return plugin_tree.ContextFilterResult{}
	}

	text2sound(splits[1])

	//bot_action.BotActionAPIInstance.SendGroupMessage(
	//	*((&models.MessageChain{
	//		GroupId: ctx.GroupId,
	//	}).Record("C:\\Users\\forDece\\PycharmProjects\\MoeGoe\\res\\1.mp3")),
	//	func(messageId int) {
	//	})

	return plugin_tree.ContextFilterResult{
		BreakFilter: true,
	}
}

func text2sound(message string) {
	// 创建一个 TCP 连接到本地的 50007 端口
	conn, err := net.Dial("tcp", "127.0.0.1:50007")
	if err != nil {
		fmt.Println("连接服务器失败：", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭连接失败：", err)
		}
	}(conn)

	// 要发送的消息
	// message := "こんにちは"

	// 发送消息
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送消息失败：", err)
		return
	}

	// 接收服务器的响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("接收响应失败：", err)
		return
	}

	// 打印响应
	fmt.Println("收到服务器响应：", string(buffer[:n]))
}
