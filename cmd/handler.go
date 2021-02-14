package main

import (
	"fmt"
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"github.com/viamAhmadi/gReceiver2/pkg/util"
	"time"
)

func (a *application) connectionHandler(c *conn.ReceiveConn) {
	a.infoLog.Println("new connection")
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
		case _ = <-c.MsgChan:
			//fmt.Println("new message ", *m)
			if c.Counter == c.Count {
				c.Successful = conn.YES
				if err := c.SendPacketFactor(c.From, &conn.Factor{ConnId: c.Id, Successful: c.Successful, List: nil}); err != nil {
					a.errorLog.Println("error in send factor to ", c.Id)
				}
				a.infoLog.Printf("connId: %s received: %d\n", c.Id, c.Messages.Count())
				c.Close()
				return
			}
		case <-c.CloseCh:
			fmt.Println("close connection")
			return
		case <-time.After(util.CalculateTimeout(3, c.Count)):
			if c.IsOpen == 0 {
				return
			}
			missedCount := c.CalculateMissingMessages()
			if missedCount == 0 {
				c.Successful = conn.YES
			} else {
				c.Successful = conn.NO
			}
			if err := c.SendPacketFactor(c.From, &conn.Factor{ConnId: c.Id, Successful: c.Successful, List: c.MissingMessages}); err != nil {
				a.errorLog.Println("error in send factor to ", c.Id)
			}
			a.infoLog.Printf("connId: %s received: %d - connection closed, timeout\n", c.Id, c.Messages.Count())
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
		fmt.Println("connection is closed")
		return
	}
	rConn.MsgChan <- msg
}
