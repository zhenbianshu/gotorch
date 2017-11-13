package stat

import (
	"fmt"
	"time"
)

func ShowStat() {
	// 最好能依附到标准输出
	for {
		clear()
		printStat()
		time.Sleep(1 * time.Second)
	}
}

/**
 *清屏
 */
func clear() {
	for i := 0; i < 100; i++ {
		fmt.Printf("\n")
	}
	fmt.Printf("\033[0;0H")
}

func printStat() {
	stats := collectStat()
	for _, stat := range stats {
		fmt.Println(stat)
	}
	fmt.Printf("\033[0;0H")
}

func collectStat() (stats []string) {
	cpuStat := getCpuStat()
	memStat := getMemStat()
	taskStat := getTaskStat()

	stats = []string{cpuStat, memStat, taskStat}
	return stats
}
