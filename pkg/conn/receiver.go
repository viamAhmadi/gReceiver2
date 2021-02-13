package conn

import (
	"github.com/zeromq/goczmq"
)

type Receiver struct {
	sock *goczmq.Sock
}

func NewReceiver(sock *goczmq.Sock) *Receiver {
	b := Receiver{sock: sock}
	return &b
}
