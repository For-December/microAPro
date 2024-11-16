package global_data

import (
	"microAPro/utils/containers"
	"time"
)

var BotMessageIdStack = containers.NewStackGroup[int64](20, 2*time.Minute)
