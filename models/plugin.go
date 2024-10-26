package models

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}

type PluginBaseInterface interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult

	GetPluginInfo() string
}

// 树形匹配

type ContextResult struct {
	IsContinue bool
	Error      error
}

type PluginHandler func(ctx *MessageContext) ContextResult

type CallbackFunc struct {
	AfterEach  PluginHandler
	OnNotFound PluginHandler
}

type PluginInterface interface {
	GetPaths() []string // ban [user] [duration]
	GetPluginInfo() string
	GetPluginHandler() PluginHandler
}
