package main

import (
	"download/config"
	"download/grpc_client"
)

func main() {
	logger := grpc_client.Loger()
	ch := make(chan string)
	cn, err := config.ReadClientConfig()
	if err != nil {
		logger.Println("客户端读取配置文件出错,请检查!!!")
	}
	go grpc_client.DownRequest(cn, ch)
	errs := <-ch
	if errs != "" {
		logger.Printf("错误信息为:%v\n", errs)
	}

}
