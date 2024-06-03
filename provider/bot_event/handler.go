package bot_event

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
	fmt.Println(string(payload))
}

func (c *handler) OnPong(socket *gws.Conn, payload []byte) {
	fmt.Println(string(payload))

}

func (c *handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	botEventChannel <- message.Bytes()
}
