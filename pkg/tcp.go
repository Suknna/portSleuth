package pkg

import (
	"net"
	"time"
)

// tcp函数用于检测指定的IP地址和端口是否开放。
// 参数：
//
//	ip: 要检测的IP地址
//	port: 要检测的端口号
//	td: 超时时间
func Check(ip string, ports []string, td time.Duration) (openPorts []string, closePorts []string) {

	for _, p := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, p), td)
		if err != nil {
			closePorts = append(closePorts, p)
		} else {
			defer conn.Close()
			openPorts = append(openPorts, p)
		}
	}
	return openPorts, closePorts
}
