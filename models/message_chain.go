package models

import "microAPro/models/entity"

type MessageChain struct {
	Messages []entity.CommonMessage `json:"messages"`
	UserId   int                    `json:"user_id"`
	GroupId  int                    `json:"group_id"`
	FromId   int                    `json:"from_id"`
	TargetId int                    `json:"target_id"`
}

type JsonTypeMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func (receiver *MessageChain) ToString() string {
	resStr := ""
	for _, message := range receiver.Messages {
		switch message.MessageType {
		case "text":
			resStr += message.MessageContent["text"].(string)
		case "image":
			resStr += "<-image->[" + message.MessageContent["url"].(string) + "]"
		case "at":
			resStr += "@(" + message.MessageContent["qq"].(string) + ")"
		}
		resStr += "\n"
	}
	return resStr
}

func (receiver *MessageChain) ToJsonTypeMessage() []JsonTypeMessage {
	message := make([]JsonTypeMessage, 0)

	for _, commonMessage := range receiver.Messages {
		switch commonMessage.MessageType {
		case "text":
			message = append(message, JsonTypeMessage{
				Type: "text",
				Data: map[string]interface{}{"text": commonMessage.MessageContent["text"].(string)},
			})
		case "image":
			message = append(message, JsonTypeMessage{
				Type: "image",
				Data: map[string]interface{}{"url": commonMessage.MessageContent["url"].(string)},
			})
		case "at":
			message = append(message, JsonTypeMessage{
				Type: "at",
				Data: map[string]interface{}{"qq": commonMessage.MessageContent["qq"].(string)},
			})
		}
	}

	return message
}

func (receiver *MessageChain) Text(content string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "text",
		MessageContent: map[string]interface{}{"text": content},
	})
	return receiver
}

func (receiver *MessageChain) Image(url string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "image",
		MessageContent: map[string]interface{}{"url": url},
	})
	return receiver

}

func (receiver *MessageChain) At(qq string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "at",
		MessageContent: map[string]interface{}{"qq": qq},
	})
	return receiver
}
