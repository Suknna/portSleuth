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
	"os"
	"portsleuth/cfg"
	"portsleuth/pkg"

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
		pl, _ := cmd.Flags().GetString("protocol")
		td, _ := cmd.Flags().GetString("timeout")
		if ipStr != "" && pStr != "" {
			portSleuthRun(ipStr, pStr, pl, td)
		} else {
			panic("Enter at least one IP address and port")
		}
	},
}

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
	rootCmd.Flags().StringP("protocol", "P", "tcp", "Enter protocol. Like: {tcp|udp}")
	rootCmd.Flags().StringP("timeout", "s", "3", "Enter the timeout in seconds.")
}

func portSleuthRun(ipStr string, pStr string, plStr string, to string) {
	ipSlice := cfg.ParseIP(ipStr)
	portSlice := cfg.ParsePort(pStr)
	pl := cfg.ParseProtocol(plStr)
	td := cfg.ParseTime(to)
	for _, ip := range ipSlice {
		for _, p := range portSlice {
			pkg.Check(ip, p, td, pl)
		}
	}
}
