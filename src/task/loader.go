package task

import (
	"bufio"
	"config"
	"io"
	"logger"
	"os"
	"strings"
)

var Task_list []*task

func Load() {
	tasks_file_name := config.GetConfig("tasks")

	file_handler, err := os.OpenFile(tasks_file_name, os.O_RDONLY, 0644)
	if err != nil {
		logger.Error("can't find the tasks file!")
	}

	reader := bufio.NewReader(file_handler)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		addTask(string(line))
	}
}

func addTask(line string) {
	task_config := strings.Split(line, " ")
	if !checkValid(task_config) {
		return
	}

	task := newTask(task_config)
	AddTask(task)

}

func AddTask(task *task) {
	Task_list = append(Task_list, task)
}
