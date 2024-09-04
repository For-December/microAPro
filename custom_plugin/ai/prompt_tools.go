package ai

import (
	"fmt"
	"time"
)

func ProcessPromptWithData(prompt string) string {

	now := time.Now()
	info := fmt.Sprintf(`## 
日常信息
当前日期：{%d/%d/%d}
当前时间：{%d:%d:%d}
当前季节：{夏季}
当前天气：{晴朗}
`, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	return info + prompt
}
