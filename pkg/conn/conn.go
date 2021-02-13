package conn

import "errors"

const YES = byte(1)
const NO = byte(0)

var ErrConvertToModel = errors.New("convert error")
var ErrConnExists = errors.New("connection exists")
var ErrDealer = errors.New("dealer was nil")

type ReceiveConn struct {
	From        []byte
	Receiver    Receiver
	Destination string
	Count       int
	Counter     int
	Id          string // 20 char
	IsOpen      int    // open 1, closed 0
	Successful  int    //  1 - 0
	Messages    *Messages
	MsgChan     chan *Message
	ErrorCh     chan Error
}

// id = from rec conns get id from the connection
// for send  generate random connection id
func NewReceiveConn(des, id string, count int) *ReceiveConn {
	return &ReceiveConn{
		Destination: des,
		Count:       count,
		Counter:     0,
		Id:          id,
		Successful:  0,
		Messages:    &Messages{},
		MsgChan:     make(chan *Message),
		ErrorCh:     make(chan Error),
	}
}

func (c *ReceiveConn) AddMsg(msg *Message) error {
	c.Counter += 1
	return nil
}

func (c *ReceiveConn) SendPacketFactor(f *Factor) error {

}
