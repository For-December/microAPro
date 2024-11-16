package bots_event

import (
	"github.com/lxzan/gws"
)

var clients map[string]*gws.Conn

var botsEventChannels = map[int]chan []byte{}
