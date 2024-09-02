package models

type MessageContext struct {
	BotAccount   string        `json:"bot_account"`
	MessageType  string        `json:"message_type"`
	MessageId    int           `json:"message_id"`
	GroupId      int           `json:"group_id"`
	UserId       int           `json:"user_id"`
	MessageChain *MessageChain `json:"message_chain"`
}

func (receiver *MessageContext) SendGroupMessage() {

}

func (receiver *MessageContext) SendPrivateMessage() {

}
