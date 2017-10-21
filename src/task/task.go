package task

import (
	"time"
)

type taskItem struct {
	times map[int][]int
	attr  attr
	last  int64
}

const (
	second = iota
	minute
	hour
	day
	week
	month
)

func (t *taskItem) isOn(order int, timePoint int) bool {
	for _, num := range t.times[order] {
		if num == timePoint {
			return true
		}
	}
	return false
}

func (t *taskItem) checkTime() bool {
	current := time.Now()
	curTimestamp := current.Unix()

	// 检查是否在当前时间被执行过
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

	t.last = curTimestamp
	return true
}

func (t *taskItem) exec() {
	// exec.Command(t.attr.Script)
}
