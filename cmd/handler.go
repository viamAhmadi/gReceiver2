package main

import (
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"time"
)

func (a *application) connectionHandler(c *conn.ReceiveConn) {
	if err := a.ReceivedConns.Add(c); err != nil {
		a.errorLog.Println("connection already exist")
		return
	}

	for {
		select {
		case m := <-c.MsgChan:
			a.infoLog.Println("new message ", m.Id)
			if c.Counter == c.Count {
				if err := c.SendPacketFactor(&conn.Factor{ConnId: c.Id, Successful: conn.YES, List: nil}); err != nil {
					a.errorLog.Println("error in send factor to ", c.Id)
				}
				return
			}
		case <-time.After(5 * time.Second):
			// send factor
			// when sender get a factor so destroy her dealer
			a.infoLog.Println("connection closed, timeout, id: ", c.Id)
			return
		}
	}
}

func (a *application) messageHandler(msg *conn.Message) {
	rConn := a.ReceivedConns.Get(msg.ConnId)
	if rConn == nil {
		a.errorLog.Println("connection not found")
		return
	}

	if err := rConn.AddMsg(msg); err != nil {
		a.errorLog.Println("message already exist")
		return
	}

	rConn.MsgChan <- msg
}
