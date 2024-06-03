package main

import (
	"microAPro/provider/bot_event"
	"os"
	"os/signal"
)

func main() {
	bot_event.Start()

}

func init() {
	interrupt := make(chan os.Signal, 1) // 一个用于接收中断信号的通道

	signal.Notify(interrupt, os.Interrupt) // 监听操作系统的中断信号，并将其发送到上面的 interrupt 通道。

	go func() {
		select {
		case <-interrupt:
			bot_event.Stop()
			println("interrupt")
		}
	}()
}
