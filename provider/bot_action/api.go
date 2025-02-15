package bot_action

import (
	"fmt"
	"github.com/bytedance/sonic"
	"microAPro/models"
	"microAPro/utils/logger"
	"time"
)

//var BotActionAPIInstance = &botActionAPI{}

type BotActionAPI struct {
	botAccount int64
}

func NewBotActionAPI(botAccount int64) *BotActionAPI {
	return &BotActionAPI{botAccount: botAccount}
}

func (receiver *BotActionAPI) GetBotAccount() int64 {
	return receiver.botAccount
}

func (receiver *BotActionAPI) SendGroupMessage(chain *models.MessageChain, callback ...func(messageId int64)) {

	type TParam struct {
		GroupId    int64                    `json:"group_id"`
		Message    []models.JsonTypeMessage `json:"message"`
		AutoEscape bool                     `json:"auto_escape"`
	}

	// 微秒时间戳，带上bot标识，理论上不会重复
	echoMsg := fmt.Sprintf("group_message_%d_%d", receiver.botAccount, time.Now().UnixMicro())
	botAction := NewBotAction(
		receiver.GetBotAccount(),
		"send_group_msg",
		TParam{
			GroupId:    chain.GetTargetId(),
			Message:    chain.ToJsonTypeMessage(),
			AutoEscape: false,
		},
		echoMsg)

	// 发送消息
	BotActionChannel <- botAction

	// 如果设置了回调则处理结果
	if len(callback) > 0 {
		receiver.solveSentRes(echoMsg, callback)
	}

}
func (receiver *BotActionAPI) solveSentRes(echoMsg string, callback []func(messageId int64)) {
	// 发完消息后处理结果
	go func() {
		for {
			select {
			case actionResult := <-botActionResChannel:
				event := BotActionResult{}

				if err := sonic.Unmarshal(actionResult, &event); err != nil {
					// 心跳包
					continue
				}

				if event.Echo == echoMsg {

					// 执行回调，结束协程
					logger.Info(event.Data)
					for _, f := range callback {
						f(event.Data.MessageId)
					}

					return
				}

				// 不是所需要的，重新放入channel
				botActionResChannel <- actionResult
				continue

			}
		}
	}()
}

func (receiver *BotActionAPI) RecallMessage(messageId int64) {

	echoMsg := fmt.Sprintf("recall_%d", messageId)

	BotActionChannel <- NewBotAction(
		receiver.botAccount,
		"delete_msg",
		map[string]int64{
			"message_id": messageId,
		},
		echoMsg)
}
