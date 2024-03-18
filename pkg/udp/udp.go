package udp
<<<<<<< HEAD
=======

import (
	"fmt"
	"net"
	"time"
)

// udp函数用于检测指定的IP地址和端口是否开放。
// 参数：
//
//	ip: 要检测的IP地址
//	port: 要检测的端口号
func UdpCheck(ip net.IP, port string, td time.Duration) {
	str := net.JoinHostPort(ip.String(), port)
	conn, err := net.DialTimeout("tcp", str, td)
	if err != nil {
		fmt.Printf("ip: %s , port: %s is closed\n", ip, port)
	} else {
		fmt.Printf("ip: %s , port: %s is opened\n", ip, port)
		defer conn.Close()
	}
}
>>>>>>> master
