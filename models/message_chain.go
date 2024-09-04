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
			resStr += "<-image->[" + message.MessageContent["file"].(string) + "]"
		case "record":
			resStr += "<-record->[" + message.MessageContent["file"].(string) + "]"
		case "at":
			resStr += "@(" + message.MessageContent["qq"].(string) + ")"
		case "reply":
			resStr += "reply(" + message.MessageContent["id"].(string) + ")"
		case "face":
			resStr += "face(" + message.MessageContent["id"].(string) + ")"
		}
		//resStr += "\n"
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
				Data: map[string]interface{}{"file": commonMessage.MessageContent["file"].(string)},
			})
		case "record":
			message = append(message, JsonTypeMessage{
				Type: "record",
				Data: map[string]interface{}{"file": commonMessage.MessageContent["file"].(string)},
			})

		case "at":
			message = append(message, JsonTypeMessage{
				Type: "at",
				Data: map[string]interface{}{"qq": commonMessage.MessageContent["qq"].(string)},
			})
		case "reply":
			message = append(message, JsonTypeMessage{
				Type: "reply",
				Data: map[string]interface{}{"id": commonMessage.MessageContent["id"].(string)},
			})
		case "face":
			message = append(message, JsonTypeMessage{
				Type: "face",
				Data: map[string]interface{}{"id": commonMessage.MessageContent["id"].(string)},
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

func (receiver *MessageChain) Image(file string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "image",
		MessageContent: map[string]interface{}{"file": file},
	})
	return receiver

}

func (receiver *MessageChain) Record(file string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "record",
		MessageContent: map[string]interface{}{"file": file},
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

// Reply 通过消息id回复
func (receiver *MessageChain) Reply(id string) *MessageChain {
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "reply",
		MessageContent: map[string]interface{}{"id": id},
	})
	return receiver
}

func (receiver *MessageChain) Face(id string) *MessageChain {

	// 关于id和表情的对应
	// https://github.com/kyubotics/coolq-http-api/wiki/%E8%A1%A8%E6%83%85-CQ-%E7%A0%81-ID-%E8%A1%A8
	receiver.Messages = append(receiver.Messages, entity.CommonMessage{
		MessageType:    "face",
		MessageContent: map[string]interface{}{"id": id},
	})
	return receiver
}
