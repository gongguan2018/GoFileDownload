package grpc_client

/*
  author: gongguan
  datetime: 20230728
  funcation: 请求服务端将下载链接切片传递过去
*/
import (
	"context"
	"download/config"
	"download/grpc_server"
	"download/pb"
	"download/pkg"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"time"
)

//请求服务端,让服务端将文件先下载到服务端
func DownRequest(cn *config.Config, ch chan string) {
	requ := &pb.DownloadRequest{
		DownloadFiles: cn.RequestURL,
	}
	ch1 := make(chan []string, 1)
	ch1 <- []string{}        //设置一个默认值
	go progress(ch, ch1, cn) //让progress和downloadFile并发运行
	go downloadFile(ch1, cn, requ)
}

//goroutinue连接服务端并获取正常的可以下载的URL
func downloadFile(ch1 chan []string, cn *config.Config, requ *pb.DownloadRequest) {
	logger := Loger()
	conn, err := pkg.ConnRpcServer(cn)
	if err != nil {
		logger.Println(err)
	}
	logger.Println("连接grpc server 正常...........")
	defer conn.Close()
	dc := pb.NewDownFileClient(conn)
	logger.Println("初始化客户端")
	//调用grpc server中的下载方法让服务端开始下载
	dr, err := dc.DownloadFile(context.Background(), requ)
	if err != nil {
		logger.Println(err)
	}
	logger.Printf("服务端已下载文件完毕，并返回了正常的URL切%v\n", dr.Downresponse)
	ch1 <- dr.Downresponse //服务端下载完成后会将正常的URL切片写入通道ch1中

}

//现在将服务端的文件下载回来
func downfile(okurl []string, cn *config.Config) error {
	//遍历这个已经下载到服务端的文件切片
	for _, v := range okurl {
		newstr := path.Base(v)
		newurl := fmt.Sprintf("%s", cn.HttpdServer.HttpdProtocol) + ":" + "//" + fmt.Sprintf("%s:%d", cn.RpcServer.ServerIP, cn.HttpdServer.HttpdPort) + "/" + "downloadurl" + "/" + newstr
		cmd := exec.Command("wget", "--no-use-server-timestamps", newurl)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}
		go func() {
			io.Copy(os.Stdout, stdout) //输出到控制台，可以看到下载过程
		}()
		go func() {
			io.Copy(os.Stderr, stderr)
		}()
		err = cmd.Start()
		if err != nil {
			return err
		}
		err = cmd.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}

//goroutine，显示下载状态并调用downurl
func progress(ch chan string, ch1 chan []string, cn *config.Config) {
	var resch []string
	logger := Loger()
	for {
		fmt.Print("服务端正在下载,请等待:")
		//循环打印50个>，主要用于提示用户服务端正在下载
		for i := 0; i < 50; i++ {
			fmt.Print(">")
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Print("\r")
		resch = <-ch1 //如果resch长度不为空，说明此时服务端已经下载完成切片中已经有数据
		if len(resch) != 0 {
			fmt.Println("服务端已下载完成.........................")
			fmt.Println("客户端将文件包下载到本地.................")
			goto downurl //通过goto跳转到downurl来进行下载，也就是本地从远程服务把文件下载回来
		}
	}
downurl:
	//将远程机器下载的文件再次下载到本地，此处是利用的wget+httpd
	er := downfile(resch, cn)
	if er != nil {
		ch <- er.Error()
	}
	logger.Println("服务端文件已下载到本地，现在删除服务端的文件.........")
	//连接grpc server用来删除远程机器的文件，节省空间
	conn, err := pkg.ConnRpcServer(cn)
	if err != nil {
		ch <- err.Error()
	}
	defer conn.Close()
	dc := pb.NewDownFileClient(conn)
	//接下来删除服务端的已下载的文件,否则空间会爆
	delrequ := &pb.DeleteRequest{
		DeleteFiles: resch,
	}
	_, err = dc.DeleteFile(context.Background(), delrequ)
	if err != nil {
		ch <- err.Error()
	}
	logger.Println("服务端文件已删除完毕................")
	errURL := grpc_server.DeleteSlice(cn.RequestURL, resch)
	if len(errURL) != 0 {
		fmt.Println()
		fmt.Printf("\033[1;33;41m[ERROR]:不能正常下载的URL为:%v\033[0m\n", errURL)
	}
	logger.Printf("不能正常下载的URL为:%v\n", errURL)
	os.Exit(0)
}
