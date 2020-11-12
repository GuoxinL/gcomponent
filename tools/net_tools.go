/*
   Created by guoxin in 2020/6/8 3:46 下午
*/
package tools

import (
	"errors"
	"net"
)

/**
获得IP地址
*/
func GetLocalIP() (string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Get Local IP error")
}
