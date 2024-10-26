package containers

import (
	"strings"
)

// RouteTrieNode 路由基数树节点
type RouteTrieNode struct {
	children map[string]*RouteTrieNode
	handler  func()
}

// RouteTrie 路由基数树
type RouteTrie struct {
	root *RouteTrieNode
}

// NewRouteTrie 创建一个新的路由基数树
func NewRouteTrie() *RouteTrie {
	return &RouteTrie{
		root: &RouteTrieNode{
			children: make(map[string]*RouteTrieNode),
		},
	}
}

// Insert 在路由基数树中插入路由和对应的处理函数
func (t *RouteTrie) Insert(path string, handler func()) {
	parts := strings.Split(path, "/")
	node := t.root
	for _, part := range parts {
		if part == "" {
			continue
		}
		if _, ok := node.children[part]; !ok {
			node.children[part] = &RouteTrieNode{
				children: make(map[string]*RouteTrieNode),
			}
		}
		node = node.children[part]
	}
	node.handler = handler
}

// Search 在路由基数树中查找路由对应的处理函数
func (t *RouteTrie) Search(path string) func() {
	parts := strings.Split(path, "/")
	node := t.root
	for _, part := range parts {
		if part == "" {
			continue
		}
		if _, ok := node.children[part]; !ok {
			return nil
		}
		node = node.children[part]
	}
	return node.handler
}
