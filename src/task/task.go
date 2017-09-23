package task

import (
	"regexp"
	"strings"
)

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

func newTask(task_config []string) *task {
	var map_key string
	each := make(map[string][]int)
	every := make(map[string]int)
	for index, arg_str := range task_config[:5] {
		map_key = time_type[index]
		every[map_key] = getEvery(arg_str)
		each[map_key] = getEach(arg_str)
	}
	return &task{}
}

func getEvery(arg_str string) int {
	every_pattern, _ := regexp.Compile("/\\d")
	every_num := every_pattern.FindString(arg_str)
	if every_num {
		return int(every_num)
	} else {
		return 1
	}
}

func getEach(arg_str string) []int {
	var nums []int
	if strings.IndexAny(arg_str, "*") > -1 {
		nums = []int{}
	} else if strings.IndexAny(arg_str, ",") {
		num_list := strings.Split(arg_str, ",")
		for num := range num_list {
			nums = append(nums, int(num))
		}
	} else if strings.IndexAny(arg_str, "-") {
		num_range := strings.Split(arg_str, "-")
		for i := int(num_range[0]); i <= int(num_range[1]); i++ {
			nums = append(nums, i)
		}
	} else {
		nums = append(nums, int(arg_str))
	}

	return nums
}
