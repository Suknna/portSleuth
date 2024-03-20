package cfg

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

// parseIP 解析传入的ip
func ParseIP(ipStr string) interface{} {
	switch ipStr {
	case 
	}
}

// scatteredIpToSlices 函数用于将不连续ip转换为切片
func scatteredIpToSlices(ips string) (netIpSlice []net.IP) {
	ipSlice := strings.Split(ips, ";")
	for _, v := range ipSlice {
		ok := IsIpv4(v)
		if !ok {
			panic("Invalid IPv4.")
		}
		netIpSlice = append(netIpSlice, net.ParseIP(v))
	}
	return netIpSlice
}

// scatteredPortToSlices 函数用于将不连续端口转换为切片
func scatteredPortToSlices(ps string) (portSlice []string) {
	portSlice = strings.Split(ps, ";")
	for _, v := range portSlice {
		ok := IsPort(v)
		if !ok {
			panic("Invalid Port.")
		}
	}
	return portSlice
}

// PortGeneration 函数用于生成指定端口范围内的所有端口号
func portGeneration(portDan string) (ports []string) {
	ok := strings.Contains(portDan, "-")
	if !ok {
		panic("If using a port range, please write it in the following way: a-b")
	}
	portSlice := strings.Split(portDan, "-")
	startP, _ := strconv.Atoi(portSlice[0])
	endP, _ := strconv.Atoi(portSlice[1])
	if startP >= endP {
		panic("Invalid Port range. Please confirm if your starting port is smaller than the ending port")
	}
	for i := startP; i <= endP; i++ {
		ok := IsPort(strconv.Itoa(i))
		if !ok {
			panic("Invalid Port.")
		}
		ports = append(ports, strconv.Itoa(i))
	}
	return ports
}

// IpGeneration 函数用于生成指定IP范围内的所有IP地址
func ipGeneration(ipDan string) (ips []net.IP) {
	ok := strings.Contains(ipDan, "-")
	if !ok {
		panic("If using a port range, please write it in the following way: a-b")
	}
	ipSlice := strings.Split(ipDan, "-")
	startIP := net.ParseIP(ipSlice[0])
	endIP := net.ParseIP(ipSlice[1])

	// 判断ip地址是否为空，如果为空表示输入错误的ipv4地址
	// 判断ip地址是否为IPV4
	for _, v := range ipSlice {
		ok := IsIpv4(v)
		if !ok {
			panic("Invalid IPv4.")
		}
	}
	// 判断ipv4地址起始是否大于终止
	if ipAtoI(startIP) >= ipAtoI(endIP) {
		panic("Invalid IP range. Please confirm if your starting port is smaller than the ending port")
	}
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

// isPort 函数用于判断传入的参数是否为端口
func isPort(s string) bool {
	patten := `^([0-9]{1,5})$|^([1-9][0-9]{4,4})$`
	matched, err := regexp.MatchString(s, patten)
	if err != nil {
		return false
	}
	return matched
}

// isIpv4 函数用于判断传入的参数是否为Ipv4地址
func isIpv4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil || ip.To4() == nil {
		return false
	}
	return true
}

// IsTcpUdp 函数用于盘传入的参数是否为TCP或者UDP
func isTcpUdp(s string) bool {
	lowerStr := strings.ToLower(s)
	return strings.Contains(lowerStr, "tcp") || strings.Contains(lowerStr, "udp")
}
