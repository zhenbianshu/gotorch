package stat

import (
	"fmt"
	"time"
)

func ShowStat() {
	for {
		clear()
		printStat()
		time.Sleep(1 * time.Second)
	}

}

func clear() {

}

func printStat() {
	stats := collectStat()
	for _, stat := range stats {
		fmt.Println(stat)
	}
}

func collectStat() (stats []string) {
	cpuStat := getCpuStat()
	memStat := getMemStat()
	taskStat := getTaskStat()

	stats = []string{cpuStat, memStat, taskStat}
	return stats
}
