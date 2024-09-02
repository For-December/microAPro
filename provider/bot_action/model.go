package bot_action

type BotAction struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

type BotActionResult struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
		MessageId int `json:"message_id"`
	} `json:"data"`
	Echo string `json:"echo"`
}

type HeartBeat struct {
	Interval int `json:"interval"`
	Status   struct {
		AppInitialized bool `json:"app_initialized"`
		AppEnabled     bool `json:"app_enabled"`
		AppGood        bool `json:"app_good"`
		Online         bool `json:"online"`
		Good           bool `json:"good"`
	} `json:"status"`
	MetaEventType string `json:"meta_event_type"`
	Time          int    `json:"time"`
	SelfId        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
}
