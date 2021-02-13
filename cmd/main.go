package main

import (
	"github.com/viamAhmadi/gReceiver2/pkg/conn"
	"github.com/viamAhmadi/gReceiver2/pkg/models/storage"
	"log"
	"os"
)

type application struct {
	receiver      *conn.Receiver
	errorLog      *log.Logger
	infoLog       *log.Logger
	ReceivedConns conn.ReceivedConns
	storage       storage.Storage
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		ReceivedConns: conn.ReceivedConns{},
		storage:       storage.New(""), // todo path of storage
	}

	app.infoLog.Println("starting receiver...")

	if err := app.startReceiving("tcp://127.0.0.1:5555"); err != nil {
		panic(err)
	}
}
