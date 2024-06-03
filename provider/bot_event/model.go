package bot_event

type BotEvent struct {
	MessageType string `json:"message_type,omitempty"`
	Time        int    `json:"time"`
	SelId       int64  `json:"sel_id"`
	PostType    string `json:"post_type,omitempty"`
}

type GroupMessageEvent struct {
	MessageType string      `json:"message_type"`
	SubType     string      `json:"sub_type"`
	MessageId   int         `json:"message_id"`
	GroupId     int         `json:"group_id"`
	UserId      int         `json:"user_id"`
	Anonymous   interface{} `json:"anonymous"`
	Message     []struct {
		Type string            `json:"type"`
		Data map[string]string `json:"data"`
	} `json:"message"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	Sender     struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Sex      string `json:"sex"`
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Level    string `json:"level"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	} `json:"sender"`
	Time     int    `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}
