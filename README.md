## ip地址查询
ip数据库下载：https://raw.githubusercontent.com/lionsoul2014/ip2region/master/data/ip2region.db


## 编译linux执行文件
```c
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o iparea main.go
```

## 接口在线示例
```c
curl https://pj8.me/ip
{"code":200,"data":{"ipinfo":{"city":"郑州","country":"中国","isp":"联通","province":"河南"}},"msg":"OK"}
```

```c
curl https://pj8.me/ip/123.14.167.164

{"code":200,"data":{"ipinfo":{"city":"郑州","country":"中国","isp":"联通","province":"河南"}},"msg":"OK"}
```
