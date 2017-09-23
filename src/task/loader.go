package task

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Task_list []*task

func Load() {
	tasks_file_name := config.GetConfig("tasks")

	file_handler, err := os.Open(tasks_file_name)

	if err != nil {
		fmt.Println("can't find the tasks file!")
		os.Exit(1)
	}

	file_data, _ := ioutil.ReadAll(file_handler)
	var task_configs []attr
	err = json.Unmarshal(file_data, task_configs)
	if err != nil {
		fmt.Println("task config parse error: " + err.Error())
		os.Exit(1)
	}

	for _, attr := range task_configs {
		task, err_desc := attr.buildTask()
		if task == nil {
			fmt.Print(err_desc)
			os.Exit(1)
		}
		task.init()
		AddTask(task)
	}
}

func AddTask(task *task) {
	Task_list = append(Task_list, task)
}
