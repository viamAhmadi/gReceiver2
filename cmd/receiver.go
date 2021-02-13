package main

import (
	"github.com/viamAhmadi/gReceiver/pkg/conn"
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
		// TODO convert above params to rc conn struct
		go a.connectionHandler(from, &(*rc)[1])
		//a.infoLog.Println("new connection")
	} else if valStr == "m" {
		// TODO convert above params to msg struct
		go a.messageHandler(from, &(*rc)[1])
	} else {
		a.infoLog.Printf("there is unknown type, value: %v\n", valStr)
	}
}
