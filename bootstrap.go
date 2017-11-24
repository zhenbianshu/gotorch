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
	"flag"
	"errors"
	"monitor"
)

const Version = "0.9"

func main() {
	defer globalRecover()

	var signalOption = flag.String("s", "", `start: 		start the service
	end: 		stop the running service
	restart:	restart the running service`)
	var helpFlag = flag.Bool("h", false, "show options")
	var versionFlag = flag.Bool("v", false, "show service version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("CopyRight @zhenbianshu V" + Version)
		os.Exit(0)
	}
	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *signalOption != "" {
		if *signalOption == "start" {
			bootStrap(false)
		} else if *signalOption == "end" {
			end()
		} else if *signalOption == "restart" {
			reload()
		}
		return
	}

	// 无选项时默认启动服务
	bootStrap(false)
}

// bootstrap a daemon process.
func bootStrap(force bool) {
	if os.Getppid() != 1 || force {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.Start()
		return
	}
	logger.Debug("bootstrap", "action :start", "pid "+strconv.Itoa(os.Getpid()))

	savePid()
	listenSignal()
	go startMonitor()

	task.Init()
	for {
		task.Run()
		time.Sleep(time.Millisecond * task.CheckInterval)
	}
}

// reload process with a customize signal.
func reload() {
	pid, err := getRunningPid()
	checkErr(err)

	syscall.Kill(pid, syscall.SIGUSR2)
}

// end the process with SIGTERM signal.
func end() {
	pid, err := getRunningPid()
	checkErr(err)

	syscall.Kill(pid, syscall.SIGTERM)
}

func startMonitor() {
	for {
		time.Sleep(60 * time.Second)
		monitor.CheckStat()
	}
}

// save the process pid in file
func savePid() {
	pidFile := config.GetConfig("pid_file")
	if common.IsFileExist(pidFile) {
		fmt.Println("pid file already exist!")
		logger.Warning("bootstrap", "warning:pid file already exist", "pid"+strconv.Itoa(os.Getpid()))
		os.Exit(1)
	}

	file, _ := os.OpenFile(pidFile, os.O_WRONLY|os.O_CREATE, 0644)
	file.Write([]byte(strconv.Itoa(os.Getpid())))
}

// start a goroutine to listen the signal.
func listenSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGUSR2)

	go func() {
		s := <-c
		if s == syscall.SIGTERM {
			task.End()
			logger.Debug("bootstrap", "action: end", "pid "+strconv.Itoa(os.Getpid()), "signal "+fmt.Sprintf("%d", s))
			os.Exit(0)
		} else if s == syscall.SIGUSR2 {
			task.End()
			bootStrap(true)
		}
	}()
}

// get the running process's pid from file
func getRunningPid() (pid int, err error) {
	pidFile := config.GetConfig("pid_file")
	if !common.IsFileExist(pidFile) {
		err = errors.New("no service running")
		logger.Warning("bootstrap", "warning :no service running", "pid "+strconv.Itoa(os.Getpid()))
	}

	pidStr, err := ioutil.ReadFile(pidFile)
	pid, err = strconv.Atoi(string(pidStr))

	return pid, err
}

func globalRecover() {
	if p := recover(); p != nil {
		pidFile := config.GetConfig("pid_file")
		if common.IsFileExist(pidFile) {
			syscall.Unlink(pidFile)
		}
		logger.Error("unexpected quit: " + fmt.Sprintf("%s", p))
		fmt.Println("unexpected quit: " + fmt.Sprintf("%s", p))
		os.Exit(222)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
