package containers

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	trie := NewRouteTrie()

	// 插入路由和处理函数
	trie.Insert("/home", func() {
		fmt.Println("Home page handler")
	})
	trie.Insert("/user/profile", func() {
		fmt.Println("User profile handler")
	})

	// 搜索路由并执行处理函数
	handler := trie.Search("/home")
	if handler != nil {
		handler()
	}
	handler = trie.Search("/user/profile")
	if handler != nil {
		handler()
	}
	handler = trie.Search("/user")
	if handler != nil {
		handler()
	}
}
