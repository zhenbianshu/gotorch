package task

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var TaskList []*taskItem

func Load() {
	tasksFileName := config.GetConfig("tasks")

	fileHandler, err := os.Open(tasksFileName)

	if err != nil {
		fmt.Println("can't find the tasks file!")
		os.Exit(1)
	}

	fileData, _ := ioutil.ReadAll(fileHandler)
	taskConfigs := make([]attr, 0)
	err = json.Unmarshal(fileData, &taskConfigs)
	if err != nil {
		fmt.Println("task config parse error: " + err.Error())
		os.Exit(1)
	}

	for _, attr := range taskConfigs {
		task, err := attr.buildTask()
		if task == nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		TaskList = append(TaskList, task)
	}

	for {
		wg := sync.WaitGroup{}
		for _, taskItem := range TaskList {
			if taskItem.checkCond() {
				wg.Add(1)
				go taskItem.exec(&wg)
			}
		}

		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
