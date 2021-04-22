## ip地址查询
ip数据库下载：https://raw.githubusercontent.com/lionsoul2014/ip2region/master/data/ip2region.db


## 编译linux执行文件
```c
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o iparea main.go
```
