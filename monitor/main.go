package monitor

import (
	"gotorch/logger"
)

func CheckStat() {
	stats := collectStat()
	for item, stat := range stats {
		logger.Debug("monitor", item+":"+stat)
	}
}

func collectStat() (stats map[string]string) {
	taskStat := getTaskStat()
	cpuStat := getCpuStat()
	memStat := getMemStat()

	stats = map[string]string{"task": taskStat, "cpu": cpuStat, "memory": memStat}
	return stats
}
