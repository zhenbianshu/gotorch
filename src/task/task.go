package task

import (
	"time"
)

type taskItem struct {
	next  int64
	times map[string][]timeRange
	attr  attr
}

type timeRange struct {
	start int
	end   int
	every int
}

var timeOrder = []string{"month", "week", "day", "hour", "minute", "second"}

func (t *taskItem) init() {
	t.setNext()
}

func (t *taskItem) setNext() {
	currentTime := make(map[string]int)
	current := time.Now()
	currentTime["month"] = int(current.Month())
	currentTime["week"] = int(current.Weekday())
	currentTime["day"] = int(current.Day())
	currentTime["hour"] = int(current.Hour())
	currentTime["minute"] = int(current.Minute())
	currentTime["second"] = int(current.Second())

	for _, order := range timeOrder {
		for _, rangeItem := range t.times[order] {
			rangeItem.start++
		}
	}
}

func (t *taskItem) CheckTime() {
	current := time.Now()
	timestamp := current.Unix()
	if timestamp >= t.next {
		t.Exec()
		t.setNext()
	}
}

func (t *taskItem) Exec() {

}
