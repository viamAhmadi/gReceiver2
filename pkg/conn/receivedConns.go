package conn

import (
	"sync"
)

type ReceivedConns map[string]*ReceiveConn

var mutex = sync.Mutex{}

func (r *ReceivedConns) Add(conn *ReceiveConn) error {
	if cFound := r.Get(conn.Id); cFound != nil {
		return ErrConnExists
	}
	mutex.Lock()
	(*r)[conn.Id] = conn
	mutex.Unlock()
	return nil
}

func (r *ReceivedConns) Get(connId string) *ReceiveConn {
	mutex.Lock()
	conn := (*r)[connId]
	mutex.Unlock()
	return conn
}
