package task

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
		task, errDesc := attr.buildTask()
		if task == nil {
			fmt.Print(errDesc)
			os.Exit(1)
		}
		TaskList = append(TaskList, task)
	}

	for {
		taskCount := len(TaskList)
		c := make(chan bool, taskCount)
		for _, taskItem := range TaskList {
			go goTask(taskItem, c)
		}

		for i := 0; i < taskCount; i++ {
			<-c
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func goTask(t *taskItem, c chan bool) {
	on := false
	if t.checkTime() {
		t.exec()
		on = true
	}

	c <- on
}
