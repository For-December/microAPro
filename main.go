package main

import (
	"microAPro/dbmodels"
	"microAPro/provider/bot_action"
	"microAPro/provider/bot_event"
	"microAPro/storage/database"
	"microAPro/utils/logger"

	"os"
	"os/signal"
)

func main() {

	if err := database.Client.AutoMigrate(&dbmodels.GroupAskAnswer{}, &dbmodels.GroupLog{}); err != nil {
		logger.Error(err)
		return
	}
	//return
	go func() {
		bot_event.Start()

	}()
	bot_action.Start()

}

func init() {
	interrupt := make(chan os.Signal, 1) // 一个用于接收中断信号的通道

	signal.Notify(interrupt, os.Interrupt) // 监听操作系统的中断信号，并将其发送到上面的 interrupt 通道。

	go func() {
		select {
		case <-interrupt:
			bot_event.Stop()
			bot_action.Stop()
			println("interrupt")
		}
	}()
}
