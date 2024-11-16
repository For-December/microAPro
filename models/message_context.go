package models

type MessageContext struct {
	BotAccount  int64  `json:"bot_account"`
	MessageType string `json:"message_type"`
	MessageId   int64  `json:"message_id"`
	//GroupId      int           `json:"group_id"`
	//UserId       int           `json:"user_id"`
	MessageChain *MessageChain `json:"message_chain"`

	// 内部路由使用
	Params map[string]string
}

func (receiver *MessageContext) GetTargetId() int64 {
	return receiver.MessageChain.GetTargetId()
}

func (receiver *MessageContext) GetFromId() int64 {
	return receiver.MessageChain.GetFromId()
}
