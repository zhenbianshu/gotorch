package stat

import (
	"fmt"
	"os/exec"
	"os"
)

func ShowStat() {
	clear()
	printStat()
}

/**
 * 清屏并将光标移到最前
 */
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf("\033[0;0H")
	// todo 将控制权还给终端
}

func printStat() {
	stats := collectStat()
	for _, stat := range stats {
		fmt.Println(stat)
	}
	fmt.Printf("\n")
}

func collectStat() (stats []string) {
	taskStat := getTaskStat()
	cpuStat := getCpuStat()
	memStat := getMemStat()

	stats = []string{taskStat, cpuStat, memStat}
	return stats
}
