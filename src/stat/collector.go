package stat

import (
	"strconv"
	"task"
)

func getCpuStat() string {
	return ""
}

func getMemStat() string {
	return ""
}

func getTaskStat() string {
	return "Total task count:" + strconv.Itoa(len(task.TaskList)) + ", running task count:" + strconv.Itoa(task.WorkingCount)
}
