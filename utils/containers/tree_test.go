package containers

import (
	"fmt"
	"microAPro/models"
	"testing"
)

var trie = NewRouteTrie(models.CallbackFunc{
	AfterEach: []models.PluginHandler{
		func(ctx *models.MessageContext) models.ContextResult {
			println("After each handler")
			return models.ContextResult{}
		},
	},
	OnNotFound: []models.PluginHandler{
		func(ctx *models.MessageContext) models.ContextResult {
			println("Not found handler")
			return models.ContextResult{}
		},
	},
})

func TestTreeBase(t *testing.T) {

	// 插入路由和处理函数
	trie.Insert("home", func(ctx *models.MessageContext) models.ContextResult {
		fmt.Println("Home page handler")
		return models.ContextResult{}
	})

	trie.Insert("user profile", func(ctx *models.MessageContext) models.ContextResult {
		fmt.Println("User profile handler")
		return models.ContextResult{}
	})

	trie.Insert("ask $qq test", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Println("参数：", k, " -> ", v)
		}
		return models.ContextResult{}
	})

	trie.Insert("$qq test", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Println("参数：", k, " -> ", v)
		}
		return models.ContextResult{}
	})

	// 搜索路由并执行处理函数
	handlers := trie.Search("home")

	for _, handler := range handlers {
		handler(nil)
	}

	handlers = trie.Search("user profile")
	for _, handler := range handlers {
		handler(nil)
	}
	handlers = trie.Search("user")
	for _, handler := range handlers {
		handler(nil)
	}
	handlers = trie.Search("ask 测试 test")
	for _, handler := range handlers {
		handler(&models.MessageContext{})
	}

	handlers = trie.Search("123 test")
	for _, handler := range handlers {
		handler(&models.MessageContext{})
	}
}

func TestTreeSuper(t *testing.T) {

	trie.Insert("@ $qq test", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Printf("参数：[%v]->[%v]\n", k, v)
		}
		return models.ContextResult{}
	})

	trie.Insert("& 123 test $testParam", func(ctx *models.MessageContext) models.ContextResult {
		for k, v := range ctx.Params {
			fmt.Printf("参数：[%v]->[%v]\n", k, v)
		}
		return models.ContextResult{}
	})

	ctx1 := &models.MessageContext{
		MessageChain: (&models.MessageChain{}).At("123").Text("test"),
	}

	ctx2 := &models.MessageContext{
		MessageChain: (&models.MessageChain{}).Reply("123").Text("test").Text("你好"),
	}

	ctx3 := &models.MessageContext{
		MessageChain: (&models.MessageChain{}).Reply("123").Text("test 你好"),
	}

	trie.SearchAndExec(ctx1)
	trie.SearchAndExec(ctx2)
	trie.SearchAndExec(ctx3)

}
