package task

type task struct {
	next  int
	each  map[string][]int
	every map[string]int
	attr  attr
}

func (task *task) init() {
	task.setNext()
}

func (task *task) setNext() {
}
