package ai

import (
	"fmt"
	"microAPro/utils/logger"
	"sync"
	"time"
)

var historyArrayMap = make(map[int][]MessageMeta)
var mutex = new(sync.Mutex)

func GetMsgMetaWithHistory(groupId int, prompt string, meta MessageMeta) []MessageMeta {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := historyArrayMap[groupId]; ok {
		if prompt != "" {
			historyArrayMap[groupId][0] = MessageMeta{
				Role:    "system",
				Content: ProcessPromptWithData(prompt),
			}

			historyArrayMap[groupId] = append(historyArrayMap[groupId], meta)
		}
		return historyArrayMap[groupId]
	}

	historyArrayMap[groupId] = []MessageMeta{
		{
			Role:    "system",
			Content: ProcessPromptWithData(prompt),
		},
		meta,
	}

	return historyArrayMap[groupId]
}

func ProcessPromptWithData(prompt string) string {

	now := time.Now()
	logger.Debug(now)
	info := fmt.Sprintf(`已知信息（你的回复需要结合下面这些信息）
当前日期：{%d/%d/%d}
当前时间：{%d:%d:%d}
当前季节：{夏季}
当前天气：{晴朗}
//群里的梗：1500 -> 某群友开盒花费1500元
结合已知信息：`, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	return info + prompt
}
