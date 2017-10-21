package task

import (
	"time"
)

type taskItem struct {
	next         int64
	times        map[int][]int
	attr         attr
	bootstrapped bool
}

const (
	second = iota
	minute
	hour
	day
	week
	month
)

func (t *taskItem) init() {
	t.getNext(false)
}

func (t *taskItem) getNext(forceNextDay bool) int64 {
	current := time.Now()
	curMonth := int(current.Month())
	curWeek := int(current.Weekday())
	curDay := int(current.Day())
	curHour := int(current.Hour())
	curMinute := int(current.Minute())
	curSecond := int(current.Second())
	curTimestamp := current.Unix()

	// 获取当前执行天的执行时间
	if !forceNextDay && t.isOn(month, curMonth) && t.isOn(week, curWeek) && t.isOn(day, curDay) {
		if curHour >= t.times[hour][len(t.times[hour])-1] && curMinute >= t.times[minute][len(t.times[minute])-1] && curSecond >= t.times[second][len(t.times[second])-1] {
			return t.getNext(true)
		}

		return 0
	} else {
		// 获取下一个执行天
		dayTimestamp := curTimestamp - int64(curHour*3600) - int64(curMinute*60) - int64(curSecond)
		for {
			dayTimestamp += 24 * 3600
			dayTime := time.Unix(dayTimestamp, 0)
			if t.isOn(month, int(dayTime.Month())) && t.isOn(week, int(dayTime.Weekday())) && t.isOn(day, int(dayTime.Day())) {
				break
			}
		}
		return dayTimestamp + int64(t.times[hour][0]) + int64(t.times[minute][0]) + int64(t.times[second][0])
	}

}

func (t *taskItem) isOn(order int, timePoint int) bool {
	for _, num := range t.times[order] {
		if num == timePoint {
			return true
		}
	}
	return false
}

func (t *taskItem) getNextPoint(order int, timePoint int) (point int, outRange bool) {
	for index, num := range t.times[order] {
		if num > timePoint {
			return num, false
		} else if num == timePoint {
			if index+1 > len(t.times[order]) {
				return 0, true
			} else {
				return t.times[order][index+1], false
			}
		}
	}

	return 0, true
}

func (t *taskItem) getNextTime(order string) {

}

func (t *taskItem) CheckTime() {
	current := time.Now()
	timestamp := current.Unix()
	if timestamp >= t.next {
		t.Exec()
		t.getNext(false)
	}
}

func (t *taskItem) Exec() {

}
