package grpc_client

import (
	"log"
	"os"
)

func Loger() *log.Logger {
	f, err := os.OpenFile("agent.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "[info] ", log.LstdFlags)
	return logger
}
