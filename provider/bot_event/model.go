package bot_event

type BotEvent struct {
	Time     int    `json:"time"`
	SelId    int64  `json:"sel_id"`
	PostType string `json:"post_type,omitempty"`
}
