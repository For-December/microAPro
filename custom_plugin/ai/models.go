package ai

type MessageMeta struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Req struct {
	Model    string        `json:"model"`
	Messages []MessageMeta `json:"messages"`
	Stream   bool          `json:"stream"`
}
type Res struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Content interface{} `json:"content"`
		} `json:"logprobs"`
		Message struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int    `json:"created"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
