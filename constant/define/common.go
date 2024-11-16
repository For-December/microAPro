package define

const ChannelBufferSize = 100

const providerBaseUrl = "ws://localhost:8081"

const suffix = "/onebot/v11/ws"

// BotActionAddr = providerBaseUrl + suffix + "/api"
func BotActionAddr(endpoint string) string {
	return endpoint + suffix + "/api"
}

// BotEventAddr = providerBaseUrl + suffix + "/event"
func BotEventAddr(endpoint string) string {
	return endpoint + suffix + "/event"
}

const BotQQ = "3090807650"

const DouBaoEndPoint = "https://ark.cn-beijing.volces.com/api/v3"

const DouBaoChat = "/chat/completions"

const VolImageEndPoint = "https://visual.volcengineapi.com"
