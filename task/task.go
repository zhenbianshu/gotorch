package task

import (
	"os/exec"
	"sync"
	"syscall"
	"time"

	"gotorch/logger"
)

type taskItem struct {
	times map[int][]bool
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

// check exec cond
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

// check if task is a daemon type
func (t *taskItem) isDaemon() bool {
	if t.attr.TaskType == TypeDaemon {
		return true
	}

	return false
}

// check running task count
func (t *taskItem) checkMax() bool {
	if t.attr.Max > 0 && len(t.pids) >= t.attr.Max {
		return false
	}
	return true

}

// check if task can exec this second
func (t *taskItem) checkTime() bool {
	current := time.Now()
	curTimestamp := current.Unix()

	if t.last >= curTimestamp {
		return false
	}

	curSec := current.Second()
	curMin := current.Minute()
	curHour := current.Hour()
	curDay := current.Day()
	curWeek := int(current.Weekday())
	curMonth := int(current.Month())

	if t.times[second][curSec] && t.times[minute][curMin] && t.times[hour][curHour] && t.times[day][curDay] && t.times[week][curWeek] && t.times[month][curMonth] {
		return true
	}

	return false
}

// check ip limit
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

// run the task, save pid and wait it ends
func (t *taskItem) exec(wg *sync.WaitGroup) {
	t.last = time.Now().Unix()

	cmd := exec.Command(t.cmd, t.args...)
	err := cmd.Start()
	wg.Done()
	if err != nil {
		logger.Error(t.attr.Command + " : " + err.Error())
		t.drop()
		return
	}
	t.pids = append(t.pids, cmd.Process.Pid)

	cmd.Wait()
	WorkingCount--
	for index, pid := range t.pids {
		if pid == cmd.Process.Pid {
			t.pids = append(t.pids[:index], t.pids[:index]...)
		}
	}
}

// kill a process by sending a SIGKILL signal
func (t *taskItem) clearTask() {
	for _, pid := range t.pids {
		syscall.Kill(pid, syscall.SIGKILL)
	}
}
