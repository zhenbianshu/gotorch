package stat

import (
	"time"
	tm "goterm"
)

func ShowStat() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println("Current Time:", time.Now().Format(time.RFC1123))
	printStat()
	tm.Flush()
}

func printStat() {
	stats := collectStat()
	for _, stat := range stats {
		tm.Println(stat)
	}
}

func collectStat() (stats []string) {
	taskStat := getTaskStat()
	cpuStat := getCpuStat()
	memStat := getMemStat()

	stats = []string{cpuStat, memStat, taskStat}
	return stats
}
