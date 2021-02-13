package main

import (
	"flag"
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"log"
	"os"
)

type application struct {
	receiver      *conn.Receiver
	errorLog      *log.Logger
	infoLog       *log.Logger
	ReceivedConns conn.ReceivedConns
}

func main() {
	addr := flag.String("addr", "tcp://127.0.0.3:5555", "Receiver address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		ReceivedConns: conn.ReceivedConns{},
	}

	app.infoLog.Println("starting receiver...")

	if err := app.startReceiving(*addr); err != nil {
		panic(err)
	}
}
