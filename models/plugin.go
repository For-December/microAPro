package models

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}
type PluginBase interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult
}
