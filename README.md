# GoFileDownload
通过Golang实现的命令行下载工具，只需要在部署一套服务端，那么就可以用多个客户端去连接此服务端下载文件

软件架构

通过grpc的方式实现客户端和服务端的通信，同时服务端和客户端需要安装软件如下： 1、服务端需要安装httpd、wget 2、客户端机器需要安装wget

工作原理

1、首先客户端根据client.yml中定义的要URL生成URL切片

2、将URL切片通过grpc发送到服务端

3、服务端获取后对这些URL进行检测，如果是正常的可下载的就进行下载，并将正常的URL切片返回给客户端

4、服务端将下载的文件保存在httpd的路劲下，这样客户端就可以通过wget来从httpd下载这些文件

5、客户端根据返回的正常URL切片来与起始的URL切片进行截取，从而获取非正常的URL

6、客户端根据返回的正常URL切片来从httpd中下载文件

7、文件下载完成后，客户端再次将正常的URL切片发送给服务端，服务端将这些URL对应的文件删除，从而保证服务端空间不会爆满

8、打印整个下载过程，并最后打开出哪些URL不能下载，如果都是可以下载就不会打印异常URL

安装教程

服务端安装httpd，可直接执行命令yum -y install httpd即可，当然你想编译安装也可以。

在httpd默认目录下创建下载存储文件夹 mkdir -p /var/www/html/downloadurl

在服务端安装wget，执行命令如下：yum -y install wget

4、 编译服务端，执行命令：go build -o file-server server.go

5、 编译客户端，执行命令：go build -o file-agent agent.go

6、 将file-server和config目录下的server.go、server.yml 上传到服务端服务器,注意：需要在file-server同级目录下创建config文件夹并将server.go、server.yml放进去

7、 将file-agent和config目录下的client.go、client.yml 上传到客户端服务器, 注意：需要在file-agent同级目录下创建config文件夹并将client.go、client.yml放进去

使用说明

1、在客户端服务器编辑client.yml文件，添加要下载的文件连接，可添加一个或者多个，如图：
![image](https://github.com/gongguan2018/GoFileDownload/assets/40058594/5485628d-ce58-4b92-bc5b-19458fd29c7d)

2、配置grpc server的IP和端口以及http的端口和协议，如图：
![image](https://github.com/gongguan2018/GoFileDownload/assets/40058594/ba86ed38-0937-4d16-b958-7393d1e9f41c)

3、然后执行file-agent即可，会显示下载过程，如图：
![image](https://github.com/gongguan2018/GoFileDownload/assets/40058594/07de6f4e-36dc-4ec8-ae2d-3d5afe24f7e6)

4、查看下载文件，已经下载到本地，如图：
![image](https://github.com/gongguan2018/GoFileDownload/assets/40058594/66e04253-fb7c-4a16-8989-b4f4353d827e)

注意： 有些url可能国内的服务器下载会提示失败，但是国外的服务器下载就可以，原因就在那一堵墙，如果想准确的下载全部的包，那么可将file-server部署在国外的服务器上 原则上，本工具并无设置下载文件大小上限，因此只要你的磁盘空间大小够用，网络不会断，那么就可以下载任何大文件，但如果服务器配置不足或者网络不稳定，那么下载大文件可能会出现异常，因此建议用此工具下载中小文件

有任何疑问可以微信搜索公众号"运维Devops"，关注并联系我
