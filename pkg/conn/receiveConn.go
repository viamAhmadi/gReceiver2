package conn

import (
	"errors"
	"github.com/zeromq/goczmq"
)

const YES = 1
const NO = 0

var ErrConvertToModel = errors.New("convert error")
var ErrConnExists = errors.New("connection exists")
var ErrDealer = errors.New("dealer was nil")

type ReceiveConn struct {
	From            []byte
	Receiver        *Receiver
	Destination     string
	Count           int
	Counter         int
	Id              string // 20 char
	IsOpen          int    // open 1, closed 0
	Successful      int    //  1 - 0
	Messages        *Messages
	MissingMessages *[]string
	MsgChan         chan *Message
	CloseCh         chan struct{}
	ErrorCh         chan Error
	MissedCount     int
}

// id = from rec conns get id from the connection
// for send  generate random connection id
func NewReceiveConn(des, id string, count int, from []byte) *ReceiveConn {
	return &ReceiveConn{
		From:            from,
		Destination:     des,
		Count:           count,
		IsOpen:          NO,
		Id:              id,
		MsgChan:         make(chan *Message),
		ErrorCh:         make(chan Error),
		CloseCh:         make(chan struct{}),
		Messages:        &Messages{},
		MissingMessages: &[]string{},
		Successful:      NO,
		Counter:         0,
		MissedCount:     0,
	}
}

func (c *ReceiveConn) AddMsg(msg *Message) error {
	c.Counter += 1
	return c.Messages.Add(msg)
}

func (c *ReceiveConn) CalculateMissingMessages() int {
	var missed int
	for i := 1; i <= c.Count; i++ {
		if m := c.Messages.Get(i); m == nil {
			*c.MissingMessages = append(*c.MissingMessages, string(i))
			missed += 1
		}
	}
	c.MissedCount = missed
	return missed
}

func (c *ReceiveConn) SendPacketFactor(to []byte, f *Factor) error {
	err := c.Receiver.sock.SendFrame(to, goczmq.FlagMore)
	if err != nil {
		return err
	}
	return c.Receiver.sock.SendFrame(*SerializeFactor(f), goczmq.FlagNone)
}

func (c *ReceiveConn) SendPacketSend(to []byte, f Send) error {
	err := c.Receiver.sock.SendFrame(to, goczmq.FlagMore)
	if err != nil {
		return err
	}
	return c.Receiver.sock.SendFrame(SerializeSend(f.ConnId), goczmq.FlagNone)
}

func (c *ReceiveConn) Close() {
	c.IsOpen = NO
	close(c.MsgChan)
	close(c.ErrorCh)
	close(c.CloseCh)
}
