package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/shirou/gopsutil/process"
)

var wg sync.WaitGroup

// setupCloseHandler 接受终止信号的函数
func setupCloseHandler(c chan (os.Signal)) {
	<-c
	os.Exit(1)
	defer wg.Done()
}

func main() {
	// 创建接受终止信号的通道
	c := make(chan os.Signal, 2)
	// 将终止信号传输到通道中
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// 等待一个goroutine
	wg.Add(1)
	// 创建一个goroutine
	go setupCloseHandler(c)

	p, err := process.Processes()
	if err != nil {
		log.Println("获取进程列表失败")
		panic(err)
	}
	for _, proc := range p {
		procName, err := proc.Name()
		if err != nil {
			log.Printf("获取pid为: %d进程名称失败\n", proc.Pid)
		}
		if procName == "1" {
			err = proc.Kill()
			if err != nil {
				log.Printf("杀死%s进程失败,pid为:%d,请手动结束!!!\n", procName, proc.Pid)
			}
		} else {
			fmt.Printf("获取到进程名称为：%s\n", procName)
		}
	}
	fmt.Println("请输入ctrl+c来结束程序")
	// 等待goroutine结束
	wg.Wait()
}
