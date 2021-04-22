https://raw.githubusercontent.com/lionsoul2014/ip2region/master/data/ip2region.db


 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o iparea main.go