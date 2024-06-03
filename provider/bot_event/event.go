package bot_event

import (
	"fmt"
	"github.com/lxzan/gws"
	"microAPro/constant/define"
	"net/http"
	"os"
)

var client *gws.Conn

func Stop() {
	client.WriteClose(1000, nil)
	println("stop event")
}

func Start() {
	var err error
	client, _, err = gws.NewClient(&handler{}, &gws.ClientOption{
		Addr: define.BotEventAddr,
		RequestHeader: http.Header{
			"Authorization": []string{"Bearer test-114514"},
		},
		ParallelEnabled: false, // 禁用并发(内置并发实现频繁创建协程，不太合适)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client.ReadLoop()
}
