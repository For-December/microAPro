package plugin_tree

import (
	"microAPro/models"
	"microAPro/provider/bot_action"
)

var CustomPlugins = make([]PluginInterface, 0)

type ContextFilterResult struct {
	BreakFilter bool
	ErrMsg      error
}

type PluginBaseInterface interface {
	ContextFilter(ctx *models.MessageContext) ContextFilterResult

	GetPluginInfo() string
}

// 树形匹配

type ContextResult struct {
	IsContinue bool
	Error      error
}

type PluginHandler func(
	api *bot_action.BotActionAPI,
	ctx *models.MessageContext,
) ContextResult

type CallbackFunc struct {
	AfterEach  []PluginHandler
	OnNotFound []PluginHandler
}

type PluginInterface interface {
	GetPaths() []string // ban [user] [duration]
	GetPluginInfo() string
	GetPluginHandler() PluginHandler

	GetScope() uint32
}
