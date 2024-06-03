package bot

import (
	"fmt"
	"github.com/lxzan/gws"
	"microAPro/constant/define"
	"net/http"
	"os"
	"os/signal"
)

func Start() {
	client, _, err := gws.NewClient(&handler{}, &gws.ClientOption{
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

	interrupt := make(chan os.Signal, 1) // 一个用于接收中断信号的通道

	signal.Notify(interrupt, os.Interrupt) // 监听操作系统的中断信号，并将其发送到上面的 interrupt 通道。

	go func() {
		select {
		case <-interrupt:
			client.WriteClose(1000, nil)
			println("interrupt")
		}
	}()
	client.ReadLoop()
}
