package task

var time_type = []string{"second", "minute", "hour", "day", "month", "week"}

const (
	SECOND = iota
	MINUTE
	HOUR
	DAY
	MONTH
	WEEK
)

type task struct {
	time_next int
	each      map[string][]int
	every     map[string]int
}
