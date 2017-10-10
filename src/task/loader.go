package task

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	var taskConfigs []attr
	err = json.Unmarshal(fileData, taskConfigs)
	if err != nil {
		fmt.Println("task config parse error: " + err.Error())
		os.Exit(1)
	}

	for _, attr := range taskConfigs {
		task, errDesc := attr.buildTask()
		if task == nil {
			fmt.Print(errDesc)
			os.Exit(1)
		}
		task.init()
		AddTask(task)
	}
}

func AddTask(task *taskItem) {
	TaskList = append(TaskList, task)
}
