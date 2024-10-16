package ai

import (
	"fmt"
	"microAPro/utils/logger"
	"time"
)

func ProcessPromptWithData(prompt string) string {

	now := time.Now()
	logger.Debug(now)
	info := fmt.Sprintf(`已知信息（你的回复需要结合下面这些信息）
当前日期：{%d/%d/%d}
当前时间：{%d:%d:%d}
当前季节：{夏季}
当前天气：{晴朗}
结合已知信息：`, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	return info + prompt
}
