package grpc_server

import (
	"log"
	"os"
)

func Loger() *log.Logger {
	f, err := os.OpenFile("server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "[info] ", log.LstdFlags)
	return logger
}
