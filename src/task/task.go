package task

import (
	"logger"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

type taskItem struct {
	times map[int][]int
	attr  attr
	last  int64
	cmd   string
	args  []string
	pids  []int
}

const (
	second = iota
	minute
	hour
	day
	week
	month
)

/**
 * 判断某个时间单位是否符合
 */
func (t *taskItem) isOn(order int, timePoint int) bool {
	for _, num := range t.times[order] {
		if num == timePoint {
			return true
		}
	}
	return false
}

/**
 * 检查执行条件
 */
func (t *taskItem) checkCond() bool {
	if !t.checkMax() {
		return false
	}

	if t.isDaemon() {
		return true
	}

	if !t.checkTime() {
		return false
	}

	return true
}

func (t *taskItem) isDaemon() bool {
	if t.attr.TaskType == TypeDaemon {
		return true
	}

	return false
}

/**
 * 检查当前执行最大进程数
 */
func (t *taskItem) checkMax() bool {
	if t.attr.Max > 0 && len(t.pids) >= t.attr.Max {
		return false
	}
	return true

}

/**
 * 检查执行时间
 */
func (t *taskItem) checkTime() bool {
	current := time.Now()
	curTimestamp := current.Unix()

	if t.last >= curTimestamp {
		return false
	}

	curTime := make(map[int]int)
	curTime[second] = current.Second()
	curTime[minute] = current.Minute()
	curTime[hour] = current.Hour()
	curTime[day] = current.Day()
	curTime[week] = int(current.Weekday())
	curTime[month] = int(current.Month())

	for i := second; i <= month; i++ {
		if !t.isOn(i, curTime[i]) {
			return false
		}
	}

	return true
}

func (t *taskItem) checkIp() bool {
	if len(t.attr.Ips) < 1 {
		return true
	}

	for _, ip := range t.attr.Ips {
		if ip == localIp {
			return true
		}
	}

	return false
}

/**
 * 执行命令
 */
func (t *taskItem) exec(wg *sync.WaitGroup) {
	t.last = time.Now().Unix()

	cmd := exec.Command(t.cmd, t.args...)
	err := cmd.Start()
	wg.Done()
	if err != nil {
		logger.Error(t.attr.Command + " : " + err.Error())
		return
	}
	t.pids = append(t.pids, cmd.Process.Pid)

	cmd.Wait()
	for index, pid := range t.pids {
		if pid == cmd.Process.Pid {
			t.pids = append(t.pids[:index], t.pids[:index]...)
		}
	}
}

/**
 * 清理进程
 */
func (t *taskItem) clearTask() {
	for _, pid := range t.pids {
		syscall.Kill(pid, syscall.SIGKILL)
	}
}
