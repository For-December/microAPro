package bot_action

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/lxzan/gws"
	"microAPro/constant/config"
	"microAPro/constant/define"
	"microAPro/utils/logger"
	"net/http"
	"os"
	"sync"
)

func Stop() {
	for act, conn := range clients {
		err := conn.WriteClose(1000, nil)
		if err != nil {
			logger.Error(err)
			return
		}
		println("stop action: ", act)
	}
}
func Start(wg *sync.WaitGroup) {

	initClients()

	// 启动所有客户端
	for _, client := range clients {
		go func(c *gws.Conn) {
			wg.Add(1)
			defer wg.Done()

			c.ReadLoop()
		}(client)
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case botAction := <-BotActionChannel: // bot 行为
				client := clients[botAction.GetBotAccount()]

				if marshalString, err := sonic.MarshalString(&botAction); err != nil {
					logger.Error(err)
					break
				} else if err = client.WriteString(marshalString); err != nil {
					logger.Error(err)
					break
				}
				//case <-botActionResChannel: // bot 行为结果
				// 其他地方接收
			}
		}

	}()

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
			Addr: define.BotActionAddr(config.EnvCfg.BotEndpoints[i]),
			RequestHeader: http.Header{
				"Authorization": []string{"Bearer test-114514"},
			},
			ParallelEnabled: false, // 禁用并发(内置并发实现频繁创建协程，不太合适)
			Logger:          nil,
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
