package conn

import "errors"

const YES = 1
const NO = 0

var ErrConvertToModel = errors.New("convert error")
var ErrConnExists = errors.New("connection exists")
var ErrDealer = errors.New("dealer was nil")

type ReceiveConn struct {
	From            []byte
	Receiver        Receiver
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
}

// id = from rec conns get id from the connection
// for send  generate random connection id
func NewReceiveConn(des, id string, count int, from []byte) *ReceiveConn {
	return &ReceiveConn{
		From:            from,
		Destination:     des,
		Count:           count,
		Id:              id,
		MsgChan:         make(chan *Message),
		CloseCh:         make(chan struct{}),
		Messages:        &Messages{},
		MissingMessages: &[]string{},
		Successful:      NO,
		Counter:         0,
	}
}

func (c *ReceiveConn) AddMsg(msg *Message) error {
	c.Counter += 1
	return nil
}

func (c *ReceiveConn) SendPacketFactor(f *Factor) error {
	return nil
}
