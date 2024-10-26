package containers

import (
	"microAPro/models"
	"strings"
)

// RouteTrieNode 路由基数树节点
type RouteTrieNode struct {
	children  map[string]*RouteTrieNode
	isParam   bool
	paramName string
	handler   models.PluginHandler
}

// RouteTrie 路由基数树
type RouteTrie struct {
	afterFunc models.PluginHandler
	root      *RouteTrieNode
}

// NewRouteTrie 创建一个新的路由基数树
func NewRouteTrie(afterFunc models.PluginHandler) *RouteTrie {
	return &RouteTrie{
		root: &RouteTrieNode{
			children: make(map[string]*RouteTrieNode),
		},
		afterFunc: afterFunc,
	}
}

// Insert 在路由基数树中插入路由和对应的处理函数
func (t *RouteTrie) Insert(path string, handler models.PluginHandler) {
	parts := strings.Split(path, "/")
	node := t.root
	for _, part := range parts {
		if part == "" {
			continue
		}
		if part[0] == ':' {
			node.children[part[1:]] = &RouteTrieNode{
				children:  make(map[string]*RouteTrieNode),
				isParam:   true,
				paramName: part[1:],
			}
			node = node.children[part[1:]]
		} else {
			if _, ok := node.children[part]; !ok {
				node.children[part] = &RouteTrieNode{
					children: make(map[string]*RouteTrieNode),
				}
			}
			node = node.children[part]
		}

	}
	node.handler = func(ctx *models.MessageContext) models.ContextResult {
		if handler(ctx).IsFinished {
			return models.ContextResult{
				IsFinished: true,
			}
		}

		return t.afterFunc(ctx)
	}
}

// Search 在路由基数树中查找路由对应的处理函数
func (t *RouteTrie) Search(path string) models.PluginHandler {
	parts := strings.Split(path, "/")
	node := t.root
	params := make(map[string]string)

	for _, part := range parts {
		if part == "" {
			continue
		}

		if node.isParam {
			// 如果是参数节点，将参数值存入params
			params[node.paramName] = part

			// 继续查找下一个节点
			node = node.children[node.paramName]

		} else if _, ok := node.children[part]; ok {
			node = node.children[part]
		} else {
			// 不是参数节点，也找不到这部分对应的节点，返回nil
			return nil
		}
	}
	return node.handler
}
