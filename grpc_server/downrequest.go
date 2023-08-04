package grpc_server

import (
	"context"
	"download/pb"
	"os"
	"path"
)

/*
  为什么要嵌入结构体UnimplementedDownFileServer？
  因为嵌入了此结构体后才相当于实现了方法DownloadFile和DeleteFile，那么就可以重写这两个方法了
*/
type RequestServer struct {
	*pb.UnimplementedDownFileServer
}

//下载文件
func (*RequestServer) DownloadFile(ctx context.Context, dr *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	logger := Loger()
	errURL := checkURL(dr.DownloadFiles)
	logger.Printf("无法正常下载的URL为:%v\n", errURL)
	okURL := DeleteSlice(dr.DownloadFiles, errURL)
	logger.Printf("可以正常下载的URL为:", okURL)
	if len(okURL) != 0 {
		_, err := downURL(okURL)
		if err != nil {
			logger.Println(err)
		}
		logger.Println("开始下载文件..................")
	}
	logger.Println("文件下载完成，将正常的URL返回给客户端..........")
	return &pb.DownloadResponse{Downresponse: okURL}, nil
}

//删除文件
func (*RequestServer) DeleteFile(ctx context.Context, dr *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	for _, v := range dr.DeleteFiles {
		urlsuffix := path.Base(v)
		//拼接到httpd的路径/var/www/html/downloadurl
		newpath := "/var/www/html/downloadurl" + "/" + urlsuffix
		err := os.Remove(newpath)
		if err != nil {
			return nil, err
		}
	}
	return &pb.DeleteResponse{DeleteResult: 1}, nil
}
