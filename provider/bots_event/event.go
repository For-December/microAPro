package bots_event

import (
	"fmt"
	"github.com/lxzan/gws"
	"microAPro/constant/config"
	"microAPro/constant/define"
	"net/http"
	"os"
	"sync"
)

func Start(wg *sync.WaitGroup) {
	initClients()
	registerCustomPlugins()
	runDispatcher()

	for _, client := range clients {
		go func(c *gws.Conn) {
			wg.Add(1)
			defer func() {
				wg.Done()
			}()

			c.ReadLoop()
		}(client)
	}

}
func Stop() {

	for act, conn := range clients {
		conn.WriteClose(1000, nil)
		println("stop event: ", act)
	}
}

func initClients() {

	size := len(config.EnvCfg.BotAccounts)
	if len(config.EnvCfg.BotEndpoints) != size {
		panic("BotAccounts cannot match BotEndpoints!")
	}

	for i := 0; i < size; i++ {
		client, _, err := gws.NewClient(&handler{
			config.EnvCfg.BotAccounts[i],
		}, &gws.ClientOption{
			Addr: define.BotEventAddr(config.EnvCfg.BotEndpoints[i]),
			RequestHeader: http.Header{
				"Authorization": []string{"Bearer test-114514"},
			},
			ParallelEnabled: false, // 禁用并发(内置并发实现频繁创建协程，不太合适)
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, ok := clients[config.EnvCfg.BotAccounts[i]]; ok {
			panic("duplicated accounts!")
		}
		clients[config.EnvCfg.BotAccounts[i]] = client
	}

}
