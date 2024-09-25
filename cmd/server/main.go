package main

import (
	"fmt"
	"iot-go/internal/config"
	"iot-go/internal/router"
	"iot-go/pkg/net"
	"log"
)

func main() {
	config.Init()
	ip, err := net.GetLocalIP()
	if err != nil {
		log.Fatalf("未能获取到网卡信息:%s", err)
	}
	port := config.AppConfig.Server.Port
	log.Printf("启动服务 http://%s:%d", ip, port)
	e := router.SetupRouter()
	// 指定端口运行服务
	if err := e.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("初始化错误:%s", err)
	}
}
