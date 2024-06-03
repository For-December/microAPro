package bot_action

import (
	"fmt"
	"github.com/lxzan/gws"
)

type handler struct{}

func (c *handler) OnOpen(socket *gws.Conn) {
	//_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
}

func (c *handler) OnClose(socket *gws.Conn, err error) {
	fmt.Println(err.Error())
}

func (c *handler) OnPing(socket *gws.Conn, payload []byte) {
	fmt.Println("ping", string(payload))
}

func (c *handler) OnPong(socket *gws.Conn, payload []byte) {
	fmt.Println("pong", string(payload))

}

func (c *handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	//fmt.Println(string(message.Bytes()))
}
