## 关于 goc_code的使用方法

# ReadHtml.go

判断网页字符串中是否包含某个关键字，如果包含则保存到log.txt文件中

```bash
me.exe https://www.zusms.com/messages/601fe5b6b9ff681aa1bb3fb6 科技
```


# golang_js

Golang解析JS，并执行js函数

```bash
#安装两个库
go get -v -u github.com/robertkrimen/otto
go get -v -u github.com/yihubaikai/gopublic
```


# http3.go

http简易服务器,默认开启9566端口
```bash
nohup ./http 80>/device/null 2>&1 & 
```

# docker.go

docker容器使用的,主要为了测试docker环境的文件读写函数 默认使用8080
```bash
docker
```
