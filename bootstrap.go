package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"task"
	"time"
)

const Version = "0.9"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("unknown parameter, use -h or --help to get help!")
		os.Exit(0)
	}

	bootType := os.Args[1]
	if bootType == "-s" || bootType == "--start" {
		bootStrap()
	} else if bootType == "-r" || bootType == "--restart" {
		// task.Reload()
	} else if bootType == "-e" || bootType == "--end" {
		task.End()
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

func bootStrap() {
	// 启动后台进程
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		return
	}

	// todo 注册信号量处理函数
	task.Init()
	for {
		task.Run()
		time.Sleep(time.Millisecond * 200)
	}
}
