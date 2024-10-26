package containers

import (
	"fmt"
	"microAPro/models"
	"testing"
)

func TestTree(t *testing.T) {
	trie := NewRouteTrie(models.CallbackFunc{
		AfterEach: func(ctx *models.MessageContext) models.ContextResult {
			println("After each handler")
			return models.ContextResult{}
		},
		OnNotFound: func(ctx *models.MessageContext) models.ContextResult {
			println("Not found handler")
			return models.ContextResult{}
		},
	})

	// 插入路由和处理函数
	trie.Insert("/home", func(ctx *models.MessageContext) models.ContextResult {
		fmt.Println("Home page handler")
		return models.ContextResult{}
	})

	trie.Insert("/user/profile", func(ctx *models.MessageContext) models.ContextResult {
		fmt.Println("User profile handler")
		return models.ContextResult{}
	})

	trie.Insert("/ask/:qq/test", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Println("参数：", k, " -> ", v)
		}
		return models.ContextResult{}
	})

	trie.Insert("/:qq/test", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Println("参数：", k, " -> ", v)
		}
		return models.ContextResult{}
	})

	// 搜索路由并执行处理函数
	handler := trie.Search("/home")
	if handler != nil {
		handler(nil)
	}
	handler = trie.Search("/user/profile")
	if handler != nil {
		handler(nil)
	}
	handler = trie.Search("/user")
	if handler != nil {
		handler(nil)
	}
	handler = trie.Search("/ask/测试/test")
	if handler != nil {
		handler(&models.MessageContext{})
	}

	handler = trie.Search("/123/test")
	if handler != nil {
		handler(&models.MessageContext{})
	}
}
