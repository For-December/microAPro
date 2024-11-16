package bots_event

import (
	"github.com/lxzan/gws"
)

var clients map[int64]*gws.Conn

var botsEventChannels = map[int64]chan []byte{}
