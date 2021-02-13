package main

import (
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"github.com/zeromq/goczmq"
)

func (a *application) startReceiving(endpoints string) error {
	router, err := goczmq.NewRouter(endpoints)
	if err != nil {
		return err
	}

	a.receiver = conn.NewReceiver(router)

	for {
		rc, err := router.RecvMessage()
		if err != nil {
			a.errorLog.Println(err)
			continue
		}
		go a.router(&rc)
	}
}

func (a *application) router(rc *[][]byte) {
	valStr := string((*rc)[1][0])
	from := (*rc)[0]
	if valStr == "c" {
		c, err := conn.ConvertToReceiveConn(from, (*rc)[1])
		if err != nil {
			a.errorLog.Println(err)
			return
		}
		c.Receiver = a.receiver
		go a.connectionHandler(c)
	} else if valStr == "m" {
		msg, err := conn.ConvertToMessage(&(*rc)[1])
		if err != nil {
			a.errorLog.Println(err)
		}
		go a.messageHandler(msg)
	} else {
		a.infoLog.Printf("there is unknown type, value: %v\n", valStr)
	}
}
