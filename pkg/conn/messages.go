package conn

import (
	"errors"
	"sync"
)

var (
	ErrMsgExist = errors.New("message exist")
)

type Messages map[int]*Message

var mutexM = sync.Mutex{}

func (m *Messages) Add(msg *Message) error {
	if mFound := m.Get(msg.Id); mFound != nil {
		return ErrMsgExist
	}
	mutexM.Lock()
	(*m)[msg.Id] = msg
	mutexM.Unlock()
	return nil
}

func (m *Messages) Get(id int) *Message {
	mutexM.Lock()
	msg := (*m)[id]
	mutexM.Unlock()
	return msg
}

func (m *Messages) Count() int {
	return len(*m)
}
