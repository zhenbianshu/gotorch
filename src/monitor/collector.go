package monitor

import (
	"strconv"
	"task"
	"os/exec"
	"runtime"
	"strings"
)

func getCpuStat() string {
	cmd := exec.Command("/usr/bin/uptime")
	res, _ := cmd.Output()
	segments := strings.Split(string(res), " ")
	return "system cpu load: " + segments[10]
}

func getMemStat() string {
	var memStat string
	if runtime.GOOS == "darwin" {
		memStat = "memory:execute 'free' yourself!"
	} else {
		cmd := exec.Command("/usr/bin/free -m")
		res, _ := cmd.Output()
		memStat = string(res)
	}
	return memStat
}

func getTaskStat() string {
	return "Total task count:" + strconv.Itoa(len(task.TaskList)) + ", running task count:" + strconv.Itoa(task.WorkingCount)
}
