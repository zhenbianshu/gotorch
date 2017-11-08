package main

import (
	"common"
	"config"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"task"
	"time"
	"logger"
)

const Version = "0.9"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("unknown parameter, use -h or --help to get help!")
		os.Exit(0)
	}

	bootType := os.Args[1]
	if bootType == "-s" || bootType == "--start" {
		bootStrap(false)
	} else if bootType == "-r" || bootType == "--restart" {
		reload()
	} else if bootType == "-e" || bootType == "--end" {
		end()
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

func bootStrap(force bool) {
	// 启动后台进程
	if os.Getppid() != 1 || force {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		return
	}
	logger.Debug(map[string]string{"action": "start", "pid": strconv.Itoa(os.Getpid())})

	savePid()
	listenSignal()

	task.Init()
	for {
		task.Run()
		time.Sleep(time.Millisecond * task.CheckInterval)
	}
}

func listenSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGINT)

	go func() {
		s := <-c
		if s == syscall.SIGTERM || s == syscall.SIGTSTP || s == syscall.SIGINT {
			task.End()
			logger.Debug(map[string]string{"action": "end", "pid": strconv.Itoa(os.Getpid()), "signal": fmt.Sprintf("%d", s)})
			os.Exit(0)
		} else if s == syscall.SIGUSR2 {
			task.End()
			bootStrap(true)
		}
	}()
}

func savePid() {
	pidFile := config.GetConfig("pid_file")
	if common.IsFileExist(pidFile) {
		fmt.Println("pid file already exist!")
		logger.Warning(map[string]string{"warning": "pid file already exist", "pid": strconv.Itoa(os.Getpid())})
		os.Exit(1)
	}

	file, _ := os.OpenFile(pidFile, os.O_WRONLY|os.O_CREATE, 0644)
	file.Write([]byte(strconv.Itoa(os.Getpid())))
}

func reload() {
	pid := getRunningPid()
	syscall.Kill(pid, syscall.SIGUSR2)
}

func end() {
	pid := getRunningPid()
	syscall.Kill(pid, syscall.SIGTERM)
}

func getRunningPid() int {
	pidFile := config.GetConfig("pid_file")
	if !common.IsFileExist(pidFile) {
		fmt.Println("no service running!")
		logger.Warning(map[string]string{"warning": "no service running", "pid": strconv.Itoa(os.Getpid())})
		os.Exit(1)
	}

	pidStr, _ := ioutil.ReadFile(pidFile)
	pid, _ := strconv.Atoi(string(pidStr))

	return pid
}
