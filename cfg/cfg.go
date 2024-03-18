package cfg

import (
	"net"
	"strconv"
	"strings"
)

// PortGeneration 函数用于生成指定端口范围内的所有端口号
func PortGeneration(portDan string) []string {
	ok := strings.Contains(portDan, "-")
	if !ok {
		//fmt.Println("If using a port range, please write it in the following way: a-b")
		panic("If using a port range, please write it in the following way: a-b")
	}
	portSlice := strings.Split(portDan, "-")
	startP, _ := strconv.Atoi(portSlice[0])
	endP, _ := strconv.Atoi(portSlice[1])
	if startP >= endP {
		//fmt.Println("Please confirm if your starting port is smaller than the ending port")
		panic("Invalid Port range. Please confirm if your starting port is smaller than the ending port")
	}
	var ports []string
	for i := startP; i <= endP; i++ {
		ports = append(ports, strconv.Itoa(i))
	}
	return ports
}

// IpGeneration 函数用于生成指定IP范围内的所有IP地址
func IpGeneration(ipDan string) []net.IP {
	ok := strings.Contains(ipDan, "-")
	if !ok {
		panic("If using a port range, please write it in the following way: a-b")
	}
	ipSlice := strings.Split(ipDan, "-")
	startIP := net.ParseIP(ipSlice[0])
	endIP := net.ParseIP(ipSlice[1])

	// 判断ip地址是否为空，如果为空表示输入错误的ipv4地址
	if startIP == nil || endIP == nil {
		panic("Invalid IP address.")
	}
	// 判断ip地址是否为IPV4
	if startIP.To4() == nil || endIP.To4() == nil {
		panic("Only IPv4 addresses are supported.")
	}
	// 判断ipv4地址起始是否大于终止
	if ipAtoI(startIP) >= ipAtoI(endIP) {
		panic("Invalid IP range. Please confirm if your starting port is smaller than the ending port")
	}

	var ips []net.IP

	// 生成ip
	for {
		ips = append(ips, startIP)
		nextIP := incrementIP(startIP)
		if nextIP.Equal(endIP) {
			ips = append(ips, startIP)
			break
		}
		startIP = nextIP
	}
	return ips
}

// ipAtoI 函数将IP地址转换为整数
func ipAtoI(ip net.IP) int {
	bits := strings.Split(ip.String(), ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64
	sum += int64(b0)<<24 + int64(b1)<<16 + int64(b2)<<8 + int64(b3)
	return int(sum)
}

// incrementIP 函数用于递增IP地址
func incrementIP(ip net.IP) net.IP {
	ip = ip.To4()
	if ip == nil {
		return nil
	}
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
	return ip
}
