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

func (t *taskItem) init() {
	t.setNext()
}

func (t *taskItem) setNext() {
	// todo method to get next time
}

func (t *taskItem) CheckTime() {
	current := time.Now()
	timestamp := current.Unix()

	if timestamp > t.next {
		t.Exec()
		t.setNext()
	}
}

func (t *taskItem) Exec() {

}
