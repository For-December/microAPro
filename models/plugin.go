package models

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}

type PluginBaseInterface interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult

	GetPluginInfo() string
}

type ContextResult struct {
	IsFinished bool
	Error      error
}

type PluginHandler func(ctx *MessageContext) ContextResult

type PluginInterface interface {
	GetPath() string // ban [user] [duration]
	GetPluginInfo() string
	PluginHandler() func(ctx *MessageContext) ContextResult
}
