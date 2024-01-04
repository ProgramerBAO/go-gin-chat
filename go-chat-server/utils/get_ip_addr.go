package utils

import (
	"fmt"
	"net"
)

func GetLocalIP() (string, error) {
	// 连接到一个公共服务（Google的DNS服务器）
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// 获取连接的本地地址
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

func GetLocalIP2() []string {
	addrs, _ := net.InterfaceAddrs()
	res := make([]string, 2, 4)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				res = append(res, ipnet.IP.String())
			}
		}
	}
	return res
}
