package task

type task struct {
	time_next int
	config    string
}

func newTask() *task {
	return &task{}
}
