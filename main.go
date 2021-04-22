package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"iparea/ip2region"
	"log"
	"os"
)

var region *ip2region.Ip2Region

func init() {
	region, _ = ip2region.New("ip2region.db")
}

func InitLogging(filename string) {
	logFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter) // You may need this
}

func main() {
	InitLogging("iparea.log")
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ip", func(c *gin.Context) {
		cip := c.ClientIP()
		cipinfo, _ := region.MemorySearch(cip)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "OK",
			"data": gin.H{
				"ip": cip,
				"ipinfo": gin.H{
					"country":  cipinfo.Country,
					"province": cipinfo.Province,
					"city":     cipinfo.City,
					"isp":      cipinfo.ISP,
				},
			},
		})
	})

	r.GET("/ip/:ip", func(c *gin.Context) {
		ip := c.Params.ByName("ip")
		ipInfo, _ := region.MemorySearch(ip)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "OK",
			"data": gin.H{
				"ipinfo": gin.H{
					"country":  ipInfo.Country,
					"province": ipInfo.Province,
					"city":     ipInfo.City,
					"isp":      ipInfo.ISP,
				},
			},
		})

	})
	defer region.Close()
	r.Run(":7000")
}
