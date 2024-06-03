package channels

type AIAsk struct {
	AskerId  int    `json:"askerId"`
	GroupId  int    `json:"groupId"`
	Question string `json:"question"`
}
