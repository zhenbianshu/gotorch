package task

type taskItem struct {
	next  int
	each  map[string][]int
	every map[string]int
	attr  attr
}

func (task *taskItem) init() {
	task.setNext()
}

func (task *taskItem) setNext() {
}
