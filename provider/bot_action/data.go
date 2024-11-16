package bot_action

import "github.com/lxzan/gws"

var clients map[int64]*gws.Conn

var botActionResChannel = make(chan []byte, 1000)
