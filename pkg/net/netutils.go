package net

import (
	"fmt"
	"log"
	"net"
)

func GetLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("未能获取到本机ip地址")
	}
	for _, iface := range interfaces {
		// 排除回环接口 (lo) 和未开启的接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && ip.To4() != nil {
				return ip.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no network interface found")

}
