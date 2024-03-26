/*
Copyright © 2024 Suknna

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"portsleuth/cfg"
	"portsleuth/pkg"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portsleuth",
	Short: "Port detection tool",
	Long:  `Detect the port connection between the current host and the target host.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		ipStr, _ := cmd.Flags().GetString("ip")
		pStr, _ := cmd.Flags().GetString("port")
		td, _ := cmd.Flags().GetString("timeout")
		goNum, _ := cmd.Flags().GetInt("goroutine")
		if ipStr != "" && pStr != "" {
			portSleuthRun(ipStr, pStr, td, goNum)
		} else {
			panic("Enter at least one port and IPv4 address.")
		}
	},
}

var wg sync.WaitGroup

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("ip", "i", "", "Enter  Ipv4 address. Like: {192.168.1.2|192.168.1.2-192.168.1.222|192.168.1.3,192.168.3.2,192.168.4.5}")
	rootCmd.Flags().StringP("port", "p", "", "Enter port. Like: {80|80-8080|80,22,39,60}")
	rootCmd.Flags().StringP("timeout", "s", "0.2", "Enter the timeout in seconds.")
	rootCmd.Flags().IntP("goroutine", "g", 10, "Enter the number of goroutines.")
}

// portSleuthRun 函数用于并发地检查给定的 IP 地址和端口号，以确定它们是否可用。
// 参数：
//
//	ipStr: 一个包含多个 IP 地址的字符串
//	pStr: 一个包含多个端口号的字符串
//	to: 超时时间，格式为 "1s"、"2m"、"3h" 等。
//	goNum: 同时运行的最大协程数。
//	fmtS: 日志格式化字符串
func portSleuthRun(ipStr string, pStr string, to string, goNum int) {
	ipSlice := cfg.ParseIP(ipStr)
	portSlice := cfg.ParsePort(pStr)
	td := cfg.ParseTime(to)
	ch := make(chan struct{}, goNum)
	for _, ip := range ipSlice {
		wg.Add(1)
		ch <- struct{}{}
		go func(ip string, ps []string, td time.Duration) {
			defer wg.Done()
			op, cp := pkg.Check(ip, ps, td)
			fmt.Printf(`
Ipv4 %s :
    opened:
		%s
    closed:
		%s`, ip, op, cp)
			<-ch
		}(ip, portSlice, td)

	}
	wg.Wait()
}
