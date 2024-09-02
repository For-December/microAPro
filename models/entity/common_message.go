package entity

type CommonMessage struct {
	MessageType    string                 `json:"message_type"`
	MessageContent map[string]interface{} `json:"message_content"`
}
