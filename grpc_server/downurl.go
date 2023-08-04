package grpc_server

import (
	"download/config"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//检查要下载的URL是否正常，并返回异常的URL切片
func checkURL(urls []string) []string {
	var errURL []string
	for _, urlValue := range urls {
		cmd := exec.Command("wget", "--spider", "--timeout=5", "--tries=1", urlValue)
		_, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(err.Error(), "exit") {
				errURL = append(errURL, urlValue)
			}
		}
	}
	return errURL

}

//删除切片中异常的url,返回正常的URL切片
func DeleteSlice(url, delslice []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range url {
		if !m[v] {
			m[v] = true
		}
	}
	for _, v1 := range delslice {
		delete(m, v1)
	}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

//服务端执行命令下载文件
func downURL(url []string) (string, error) {
	sc, err := config.ReadServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range url {
		cmd := exec.Command("wget", "-c", "-P", fmt.Sprintf("%s", sc.Server.DownPath), v)
		//	stdout, _ := cmd.StdoutPipe()
		//	stderr, _ := cmd.StderrPipe()
		//	go func() {
		//		io.Copy(os.Stdout, stdout)
		//	}()
		//	go func() {
		//		io.Copy(os.Stderr, stderr)
		//	}()
		err := cmd.Start()
		if err != nil {
			return "", err
		}
		err = cmd.Wait()
		if err != nil {
			return "", err
		}
	}
	return sc.Server.DownPath, nil
}
