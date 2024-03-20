package pkg

import (
	"fmt"
	"net"
	"time"
)

// tcp函数用于检测指定的IP地址和端口是否开放。
// 参数：
//
//	ip: 要检测的IP地址
//	port: 要检测的端口号
//	td: 超时时间
//	pl: 传入的协议TCP或者UDP
func Check(ip string, port string, td time.Duration, pl string) {
	str := net.JoinHostPort(ip, port)
	conn, err := net.DialTimeout(pl, str, td)
	if err != nil {
		fmt.Printf("protocol: %s , ip: %s , port: %s is closed\n", pl, ip, port)
	} else {
		fmt.Printf("ip: %s , port: %s is opened\n", ip, port)
		defer conn.Close()
	}
}
