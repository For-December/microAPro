package models

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}

type PluginBaseInterface interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult

	GetPluginInfo() string
}

type PluginInterface interface {
	GetPath() string // ban [user] [duration]
	GetPluginInfo() string

	ContextFilter(ctx *MessageContext) ContextFilterResult
}
