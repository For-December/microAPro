package models

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}

type PluginBaseInterface interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult

	GetPluginInfo() string
}
