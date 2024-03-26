package cfg

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var AddrObj struct {
	ip        string
	openPort  []string
	closePort []string
}

func ParseFmtString(fmtS string) string {
	fs := strings.ToLower(fmtS)
	if fs != "closed" && fs != "opened" {
		fmt.Println("Invalid format string.")
		os.Exit(1)
	}
	return fs
}

// parseIP 解析传入的ip
func ParseIP(ipStr string) []string {
	if isIpv4(ipStr) {
		return []string{ipStr}
	} else if strings.Contains(ipStr, "-") {
		return ipGeneration(ipStr)
	} else if strings.Contains(ipStr, ",") {
		return scatteredIpToSlices(ipStr)
	} else {
		fmt.Println("Invalid ip flag.")
		os.Exit(1)
		return nil
	}
}

// parsePort 解析传入的port
func ParsePort(portStr string) []string {
	if isPort(portStr) {
		return []string{portStr}
	} else if strings.Contains(portStr, "-") {
		return portGeneration(portStr)
	} else if strings.Contains(portStr, ",") {
		return scatteredPortToSlices(portStr)
	} else {
		fmt.Println("Invalid port flag.")
		os.Exit(1)
		return nil
	}

}

// parseProtocol 解析传入的协议
func ParseProtocol(pl string) string {
	lowerStr := strings.ToLower(pl)
	if strings.Contains(lowerStr, "tcp") || strings.Contains(lowerStr, "udp") {
		return lowerStr
	} else {
		fmt.Println("Invalid protocol flag.")
		os.Exit(1)
		return ""
	}
}

func ParseTime(t string) time.Duration {
	td, err := time.ParseDuration(t + "s")
	if err != nil {
		fmt.Printf("Invalid timeout value: %s\n", err)
		os.Exit(1)
	}
	return td
}

// scatteredIpToSlices 函数用于将不连续ip转换为切片
func scatteredIpToSlices(ips string) []string {
	ipSlice := strings.Split(ips, ",")
	for _, v := range ipSlice {
		ok := isIpv4(v)
		if !ok {
			fmt.Printf("Invalid discontinuous address. IP: %s \n", v)
			os.Exit(1)
		}
	}
	return ipSlice
}

// scatteredPortToSlices 函数用于将不连续端口转换为切片
func scatteredPortToSlices(ps string) []string {
	portSlice := strings.Split(ps, ",")
	for _, v := range portSlice {
		ok := isPort(v)
		if !ok {
			fmt.Printf("Invalid discontinuous port. Port: %s \n", v)
			os.Exit(1)
		}
	}
	return portSlice
}

// PortGeneration 函数用于生成指定端口范围内的所有端口号
func portGeneration(portDan string) (ports []string) {
	ok := strings.Contains(portDan, "-")
	if !ok {
		fmt.Printf("Invalid discontinuous port.  %s\n", portDan)
		os.Exit(1)
	}
	portSlice := strings.Split(portDan, "-")
	startP, _ := strconv.Atoi(portSlice[0])
	endP, _ := strconv.Atoi(portSlice[1])
	if startP >= endP {
		fmt.Printf("Invalid discontinuous port. %s\n", portDan)
		os.Exit(1)
	}
	for i := startP; i <= endP; i++ {
		ok := isPort(strconv.Itoa(i))
		if !ok {
			fmt.Printf("Invalid Port. port: %d\n", i)
		}
		ports = append(ports, strconv.Itoa(i))
	}
	return ports
}

// IpGeneration 函数用于生成指定IP范围内的所有IP地址
func ipGeneration(ipDan string) (ips []string) {
	ok := strings.Contains(ipDan, "-")
	if !ok {
		fmt.Printf("Invalid discontinuous IPv4. %s\n", ipDan)
		os.Exit(1)
	}
	ipSlice := strings.Split(ipDan, "-")
	startIP := net.ParseIP(ipSlice[0])
	endIP := net.ParseIP(ipSlice[1])

	// 判断ip地址是否为空，如果为空表示输入错误的ipv4地址
	// 判断ip地址是否为IPV4
	for _, v := range ipSlice {
		ok := isIpv4(v)
		if !ok {
			fmt.Printf("Invalid IPv4. %s\n", v)
			os.Exit(1)
		}
	}
	// 判断ipv4地址起始是否大于终止
	if ipAtoI(startIP) >= ipAtoI(endIP) {
		fmt.Printf("Invalid discontinuous IPv4. %s\n", ipDan)
		os.Exit(1)
	}
	// 生成ip
	for {
		ips = append(ips, startIP.String())
		nextIP := incrementIP(startIP)
		if nextIP.Equal(endIP) {
			ips = append(ips, endIP.String())
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
	if port, err := strconv.Atoi(s); err == nil && port >= 0 && port <= 655335 {
		return true
	}
	return false
}

// isIpv4 函数用于判断传入的参数是否为Ipv4地址
func isIpv4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil || ip.To4() == nil {
		return false
	}
	return true
}

// // IsTcpUdp 函数用于盘传入的参数是否为TCP或者UDP
// func isTcpUdp(s string) bool {
// 	lowerStr := strings.ToLower(s)
// 	return strings.Contains(lowerStr, "tcp") || strings.Contains(lowerStr, "udp")
// }
