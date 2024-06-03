package models

import "microAPro/models/entity"

type MessageChain struct {
	Messages []entity.CommonMessage `json:"messages"`
	FromId   int                    `json:"from_id"`
	TargetId int                    `json:"target_id"`
}

func (receiver *MessageChain) Text(content string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "text",
		MessageContent: content,
	})
	return receiver
}
