package task

type taskItem struct {
	next  int
	each  map[string][]int
	every map[string]int
	attr  attr
}

func (t *taskItem) init() {
	t.setNext()
}

func (t *taskItem) setNext() {
}
