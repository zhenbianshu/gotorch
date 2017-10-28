package task

import (
	"config"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
)

var TaskList map[string]*taskItem
var configMd5 [16]byte

func Run() {
	fileData := readFile()
	if !checkMd5(fileData) {
		load(fileData)
	}

	wg := sync.WaitGroup{}
	for _, taskItem := range TaskList {
		if taskItem.checkCond() {
			wg.Add(1)
			go taskItem.exec(&wg)
		}
	}

	wg.Wait()
}

func readFile() []byte {
	tasksFileName := config.GetConfig("tasks")

	fileHandler, err := os.Open(tasksFileName)

	if err != nil {
		fmt.Println("can't find the tasks file!")
		os.Exit(1)
	}
	fileData, _ := ioutil.ReadAll(fileHandler)

	return fileData
}

func checkMd5(fileData []byte) bool {
	sum := md5.Sum(fileData)
	if len(configMd5) > 0 && sum != configMd5 {
		return false
	}

	return true
}

func load(fileData []byte) {
	taskConfigs := make([]attr, 0)
	err := json.Unmarshal(fileData, &taskConfigs)
	if err != nil {
		fmt.Println("task config parse error: " + err.Error())
		os.Exit(1)
	}

	for _, attr := range taskConfigs {
		if TaskList[attr.Command] != nil && reflect.DeepEqual(TaskList[attr.Command].attr, attr) {
			continue
		}
		task, err := attr.buildTask()
		if task == nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		if TaskList == nil {
			TaskList = make(map[string]*taskItem)
		}
		TaskList[task.attr.Command] = task
	}
}
