package define

// MessageType 消息类型
type MessageType = string

const (
	GroupMsg   MessageType = "group"
	PrivateMsg MessageType = "private"
	TempMsg    MessageType = "temp"
)
