package main

import (
	"download/config"
	"download/grpc_server"
	"download/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	scn, err := config.ReadServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	//创建监听
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", scn.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建一个grpc服务对象
	service_object := grpc.NewServer()
	//将服务对象注册到grpc的内部注册中心
	/*
		           为什么注册的时候要用grpc_server下的结构体RequestServer呢？
			   通过查看注册函数的用法：func RegisterDownFileServer(s grpc.ServiceRegistrar, srv DownFileServer)
			   可以看出，此注册函数的第一个参数为grpc服务对象，第二个参数其实是一个接口，此接口中包含的方法正是
			   我们在proto中定义的两个方法，因此就可以理解为只要实现了这个接口的结构体都可以作为参数传递给此接口
		           因此RequestServer结构体要先内嵌pb.UnimplementedDownFileServer结构体，因为pb.UnimplementedDownFileServer结构体实现了这个接口，因此内嵌后相当于RequestServer也实现了此接口，此时我们可重写接口中的方法，来根据自己的需求实现某些目的
	*/
	pb.RegisterDownFileServer(service_object, &grpc_server.RequestServer{})
	//启动grpc
	if err := service_object.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
