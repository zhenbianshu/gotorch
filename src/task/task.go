package task

import (
	"time"
)

type taskItem struct {
	times map[int][]int
	attr  attr
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

func (t *taskItem) exec() {

}
