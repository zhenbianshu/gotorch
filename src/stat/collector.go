package stat

import (
	"strconv"
	"task"
)

// todo 解析linux系统命令结果

func getCpuStat() string {
	return ""
}

func getMemStat() string {
	return ""
}

func getTaskStat() string {
	return "Total task count:" + strconv.Itoa(len(task.TaskList)) + ", running task count:" + strconv.Itoa(task.WorkingCount)
}
