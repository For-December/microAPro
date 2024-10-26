package containers

import (
	"microAPro/models"
	"strings"
)

// RouteTrieNode 路由基数树节点
type RouteTrieNode struct {
	children map[string]*RouteTrieNode
	//isParam   bool
	paramName string
	handler   models.PluginHandler
}

// RouteTrie 路由基数树
type RouteTrie struct {
	callbackFunc models.CallbackFunc
	root         *RouteTrieNode

	depth int
}

// NewRouteTrie 创建一个新的路由基数树
func NewRouteTrie(callbackFunc models.CallbackFunc) *RouteTrie {

	// 如果没有设置回调函数，使用默认的回调函数
	if callbackFunc.OnNotFound == nil {
		callbackFunc.OnNotFound = func(ctx *models.MessageContext) models.ContextResult {
			return models.ContextResult{}
		}
	}

	if callbackFunc.AfterEach == nil {
		callbackFunc.AfterEach = func(ctx *models.MessageContext) models.ContextResult {
			return models.ContextResult{}
		}
	}

	return &RouteTrie{
		root: &RouteTrieNode{
			children: make(map[string]*RouteTrieNode),
		},
		callbackFunc: callbackFunc,
	}
}

// Insert 在路由基数树中插入路由和对应的处理函数
func (t *RouteTrie) Insert(path string, handler models.PluginHandler) {
	parts := strings.Split(path, " ")
	node := t.root
	for _, part := range parts {
		if part == "" {
			continue
		}
		if part[0] == '$' { // 参数节点，用空字符串表示该特殊节点
			if _, ok := node.children[""]; !ok {
				node.children[""] = &RouteTrieNode{
					children: make(map[string]*RouteTrieNode),
					//isParam:   true,
					paramName: part[1:],
				}
			}

			// 将子节点作为当前节点
			node = node.children[""]
		} else {
			if _, ok := node.children[part]; !ok {
				node.children[part] = &RouteTrieNode{
					children: make(map[string]*RouteTrieNode),
				}
			}
			node = node.children[part]
		}

		t.depth++

	}
	node.handler = func(ctx *models.MessageContext) models.ContextResult {
		if handler(ctx).IsContinue {
			return models.ContextResult{
				IsContinue: true,
			}
		}
		return t.callbackFunc.AfterEach(ctx)
	}
}

// Search 在路由基数树中查找路由对应的处理函数
func (t *RouteTrie) Search(path string) models.PluginHandler {
	parts := strings.Split(path, " ")
	node := t.root
	params := make(map[string]string)

	for _, part := range parts {
		if part == "" {
			continue
		}

		// 如果当前是路径的一部分，继续查找
		if _, ok := node.children[part]; ok {
			node = node.children[part]

			// 否则当前可能是参数，查找此路径下面的参数节点
		} else if _, ok = node.children[""]; ok {
			params[node.children[""].paramName] = part
			node = node.children[""]
		} else {
			// 不是参数节点，也找不到这部分对应的节点，返回404
			return t.callbackFunc.OnNotFound
		}
	}

	if node.handler == nil {
		return t.callbackFunc.OnNotFound
	}

	return func(ctx *models.MessageContext) models.ContextResult {
		// 代理函数
		if ctx != nil {
			ctx.Params = params
		}

		return node.handler(ctx)
	}
}
