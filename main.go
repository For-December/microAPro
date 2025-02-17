package main

import (
	"microAPro/dbmodels"
	"microAPro/provider/bot_action"
	"microAPro/provider/bots_event"
	"microAPro/storage/database"
	"microAPro/utils/logger"
	"sync"

	"os"
	"os/signal"
)

func main() {

	if err := database.Client.AutoMigrate(&dbmodels.GroupAskAnswer{}, &dbmodels.GroupLog{}); err != nil {
		logger.Error(err)
		return
	}

	//return
	wg := sync.WaitGroup{}
	bots_event.Start(&wg)
	bot_action.Start(&wg)

	wg.Wait()

}

func init() {
	interrupt := make(chan os.Signal, 1) // 一个用于接收中断信号的通道

	signal.Notify(interrupt, os.Interrupt) // 监听操作系统的中断信号，并将其发送到上面的 interrupt 通道。

	go func() {
		select {
		case <-interrupt:
			bots_event.Stop()
			bot_action.Stop()
			println("interrupt")
		}
	}()
}
