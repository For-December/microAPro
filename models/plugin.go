package models

type ContextFilterResult struct {
	IsContinue bool
	ErrMsg     error
}
type PluginBase interface {
	ContextFilter(ctx *MessageContext) ContextFilterResult
}
