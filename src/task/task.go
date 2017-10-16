package task

type taskItem struct {
	next  int
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
}
