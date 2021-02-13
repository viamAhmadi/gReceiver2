package main

import (
	"fmt"
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"time"
)

func (a *application) connectionHandler(c *conn.ReceiveConn) {
	c.IsOpen = 1
	if err := a.ReceivedConns.Add(c); err != nil {
		a.errorLog.Println("connection already exist")
		return
	}
	go func() {
		if err := c.SendPacketSend(c.From, conn.Send{ConnId: c.Id}); err != nil {
			a.errorLog.Println("error in send packet send to ", c.Id)
		}
	}()
	for {
		select {
		case m := <-c.MsgChan:
			fmt.Println("new message ", *m)
			if c.Counter == c.Count {
				c.Successful = conn.YES
				if err := c.SendPacketFactor(c.From, &conn.Factor{ConnId: c.Id, Successful: conn.YES, List: nil}); err != nil {
					a.errorLog.Println("error in send factor to ", c.Id)
				}
				c.Close()
				return
			}
		case <-c.CloseCh:
			fmt.Println("close connection")
			return
		case <-time.After(3 * time.Second):
			// send factor
			// when sender get a factor so destroy her dealer
			a.infoLog.Println("connection closed, timeout 3 sec, id: ", c.Id)
			c.Close()
			return
		case e := <-c.ErrorCh:
			fmt.Println("err ch", e)
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
	if rConn.IsOpen == 0 {
		fmt.Println("connection is closed but you want to write")
		return
	}
	rConn.MsgChan <- msg
}
