package pkg

import (
	"download/config"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func ConnRpcServer(cn *config.Config) (*grpc.ClientConn, error) {
	//连接grpc-server，如果10s中没有连接成功,就视为连接失败
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cn.RpcServer.ServerIP, cn.RpcServer.RPCPort), grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
