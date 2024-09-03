package bot_action

import (
	"fmt"
	"github.com/bytedance/sonic"
	"microAPro/channels"
	"microAPro/models"
	"microAPro/utils/logger"
	"time"
)

var BotActionAPIInstance = &botActionAPI{}

type botActionAPI struct {
}

func (receiver *botActionAPI) SendGroupMessage(chain models.MessageChain, callback func(messageId int)) {

	type TParam struct {
		GroupId    int                      `json:"group_id"`
		Message    []models.JsonTypeMessage `json:"message"`
		AutoEscape bool                     `json:"auto_escape"`
	}

	echoMsg := fmt.Sprintf("group_message_%d", time.Now().Unix())
	marshalString, err := sonic.MarshalString(&BotAction{
		Action: "send_group_msg",
		Params: TParam{
			GroupId:    chain.GroupId,
			Message:    chain.ToJsonTypeMessage(),
			AutoEscape: false,
		},
		Echo: echoMsg,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	channels.BotActionChannel <- marshalString
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

					// 执行回调，结束携程
					logger.Info(event.Data)
					callback(event.Data.MessageId)

					return
				}

				// 不是所需要的，重新放入channel
				botActionResChannel <- actionResult
				continue

			}
		}
	}()
}

func (receiver *botActionAPI) RecallMessage(messageId int) {

	echoMsg := fmt.Sprintf("recall_%d", messageId)
	marshalString, err := sonic.MarshalString(&BotAction{
		Action: "delete_msg",
		Params: map[string]int{
			"message_id": messageId,
		},
		Echo: echoMsg,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	channels.BotActionChannel <- marshalString
}
