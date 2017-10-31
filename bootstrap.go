package main

import (
	"fmt"
	"os"
	"task"
	"time"
)

const Version = "0.9"

func main() {
	dealCliArgs()
	task.Init()
	for {
		task.Run()
		time.Sleep(time.Millisecond * 200)
	}
}

func dealCliArgs() {
	bootType := os.Args[1]
	if bootType == "-s" || bootType == "--start" {
		return
	} else if bootType == "-r" || bootType == "--restart" {
		// task.Reload()
	} else if bootType == "-e" || bootType == "--end" {
		// task.End()
	} else if bootType == "-v" || bootType == "--version" {
		fmt.Println("CopyRight @zhenbianshu V" + Version)
	} else if bootType == "-h" || bootType == "--help" {
		fmt.Println("-s --start 启动服务")
		fmt.Println("-e --end 关闭服务")
		fmt.Println("-r --restart 平滑重启服务")
		fmt.Println("-v --version 查看服务版本")
		fmt.Println("-h --help 查看帮助")
	} else {
		fmt.Println("unknown parameter, use -h or --help to get help!")
	}
	os.Exit(0)
}
