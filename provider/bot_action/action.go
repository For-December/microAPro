package bot_action

import (
	"fmt"
	"github.com/lxzan/gws"
	"microAPro/constant/define"
	"microAPro/global_data"
	"microAPro/utils/logger"
	"net/http"
	"os"
)

var client *gws.Conn
var botActionResChannel = make(chan []byte, define.ChannelBufferSize)

func Stop() {
	client.WriteClose(1000, nil)
	println("stop action")
}
func Start() {
	var err error
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

	go func() {
		type T struct {
			GroupId    int    `json:"group_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}

		for {
			select {
			case botAction := <-global_data.BotActionChannel: // bot 行为
				if err := client.WriteString(botAction); err != nil {
					logger.Error(err)
					break
				}
			}
		}
	}()

	client.ReadLoop()
}
