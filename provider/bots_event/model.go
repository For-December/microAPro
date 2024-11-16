package bots_event

type botEvent struct {
	MessageType string `json:"message_type,omitempty"`
	Time        int    `json:"time"`
	SelId       int64  `json:"sel_id"`
	PostType    string `json:"post_type,omitempty"`
}

type message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	//File    string `json:"file"`
	//Url     string `json:"url"`
	//Summary string `json:"summary"`
	//QQ      string `json:"qq"`
	//Text    string `json:"text"`
	//SubType int    `json:"subType"`
}

type groupMessageEvent struct {
	MessageType string      `json:"message_type"`
	SubType     string      `json:"sub_type"`
	MessageId   int         `json:"message_id"`
	GroupId     int         `json:"group_id"`
	UserId      int         `json:"user_id"`
	Anonymous   interface{} `json:"anonymous"`
	Message     []message   `json:"message"`
	RawMessage  string      `json:"raw_message"`
	Font        int         `json:"font"`
	Sender      struct {
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

type privateMessageEvent struct {
	MessageType string    `json:"message_type"`
	SubType     string    `json:"sub_type"`
	MessageId   int       `json:"message_id"`
	UserId      int       `json:"user_id"`
	Message     []message `json:"message"`
	RawMessage  string    `json:"raw_message"`
	Font        int       `json:"font"`
	Sender      struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
	} `json:"sender"`
	TargetId int    `json:"target_id"`
	Time     int    `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}
