package task

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gotorch/config"
	"gotorch/logger"
	"io/ioutil"
	"net"
	"os"
	"reflect"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var TaskList map[string]*taskItem
var CheckInterval time.Duration
var WorkingCount int
var configMd5 [16]byte
var localIp string

// package init
func Init() {
	var err error
	CheckInterval, err = time.ParseDuration(config.GetConfig("interval"))
	if err != nil {
		logger.Warning("loader", "warning : no interval config, use default")
		CheckInterval = 100
	}

	TaskList = make(map[string]*taskItem)

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localIp = ipNet.IP.String()
			}
		}
	}
}

// run : check and exec tasks
func Run() {
	tasksFile := config.GetConfig("tasks")
	fileData := readFile(tasksFile)
	if !checkMd5(fileData) {
		load(fileData)
	}

	wg := sync.WaitGroup{}
	for _, taskItem := range TaskList {
		if taskItem.checkCond() {
			wg.Add(1)
			WorkingCount++
			go taskItem.exec(&wg)
		}
	}

	wg.Wait()
}

// end and clear tasks
func End() {
	for _, taskItem := range TaskList {
		taskItem.clearTask()
	}

	pidFile := config.GetConfig("pid_file")
	syscall.Unlink(pidFile)
	logger.Debug("loader", "service end, pid "+strconv.Itoa(os.Getpid()))
}

// read the task file
func readFile(fileName string) []byte {
	fileHandler, err := os.Open(fileName)
	defer fileHandler.Close()

	if err != nil {
		panic(err.Error())
	}
	fileData, _ := ioutil.ReadAll(fileHandler)

	return fileData
}

// check task file md5, if changed then reload
func checkMd5(fileData []byte) bool {
	sum := md5.Sum(fileData)
	if sum != configMd5 {
		return false
	}

	return true
}

// load the config
func load(fileData []byte) {
	taskConfigs := make([]attr, 0)
	err := json.Unmarshal(fileData, &taskConfigs)
	if err != nil {
		panic("task config parse error: " + err.Error())
	}

	for _, oldTask := range TaskList {
		for _, newAttr := range taskConfigs {
			if oldTask.attr.Command == newAttr.Command {
				goto GoOn
			}
		}

		oldTask.clearTask()
		delete(TaskList, oldTask.attr.Command)
	GoOn:
	}

	for _, attr := range taskConfigs {
		if TaskList[attr.Command] != nil {
			if reflect.DeepEqual(TaskList[attr.Command].attr, attr) {
				continue
			} else {
				TaskList[attr.Command].clearTask()
			}
		}

		task, err := attr.buildTask()
		if task == nil {
			fmt.Print(err.Error())
			logger.Error(err.Error())
			os.Exit(1)
		}

		if !task.checkIp() {
			continue
		}

		TaskList[task.attr.Command] = task
	}
}

// delete a task item
func (t *taskItem) drop() {
	for index, item := range TaskList {
		if item.attr.Command == t.attr.Command {
			delete(TaskList, index)
		}
	}
}
