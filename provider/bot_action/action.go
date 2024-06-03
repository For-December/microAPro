package bot_action

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/lxzan/gws"
	"microAPro/channels"
	"microAPro/constant/define"
	"microAPro/utils/logger"
	"net/http"
	"os"
)

var client *gws.Conn
var err error

func Stop() {
	client.WriteClose(1000, nil)
	println("stop action")
}
func Start() {
	client, _, err = gws.NewClient(&handler{}, &gws.ClientOption{
		Addr: define.BotActionAddr,
		RequestHeader: http.Header{
			"Authorization": []string{"Bearer test-114514"},
		},
		ParallelEnabled: false, // 禁用并发(内置并发实现频繁创建协程，不太合适)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type BotAction struct {
		Action string      `json:"action"`
		Params interface{} `json:"params"`
		Echo   string      `json:"echo"`
	}

	go func() {
		type T struct {
			GroupId    int    `json:"group_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}

		for {
			select {
			case botAction := <-channels.BotActionChannel: // bot 行为
				if err := client.WriteString(botAction); err != nil {
					logger.Error(err)
					break
				}
			case aiAsk := <-channels.AIChannel: // ai 问答
				logger.Info(1)
				marshalString, err := sonic.MarshalString(&BotAction{
					Action: "send_group_msg",
					Params: T{
						GroupId:    aiAsk.GroupId,
						Message:    "功能开发中~",
						AutoEscape: false,
					},
					Echo: "chat_gpt_msg",
				})
				if err != nil {
					logger.Error(err)
					break
				}
				channels.BotActionChannel <- marshalString

			}
		}
	}()

	client.ReadLoop()
}
