package task

import (
	"regexp"
	"strings"
)

type attr struct {
	Script    string
	Task_type string
	Times     string
	Max       int
	Ips       []string
}

func (config attr) timeValid() (is_valid bool, err_desc string) {
	times := strings.Split(config.Times, " ")
	if len(times) != 6 {
		return false, "time format error!"
	}

	for index, arg_str := range times {
		if arg_str == "*" {
			continue
		}

		if regexp.MatchString("[^*|\\-|,|/|a-zA-Z0-9]", arg_str) {
			return false, "undefined character!"
		}

		rex, _ := regexp.Compile("\\d")
		args := rex.FindAllString(arg_str, -1)
		for arg_num := range args {
			if !inRange(int(arg_num), index) {
				return false, "time num out of range"
			}
		}

	}
	return true, ""
}

func (config attr) buildTask() (task *task, err_desc string) {
	if is_valid, err_desc := config.timeValid(); !is_valid {
		return nil, err_desc
	}

	var map_key string
	each := make(map[string][]int)
	every := make(map[string]int)
	times := strings.Split(config.Times, " ")
	for index, arg_str := range times {
		map_key = time_type[index]
		every[map_key] = getEvery(arg_str)
		each[map_key] = getEach(arg_str)
	}

	return &task{every: every, each: each}, ""
}

func inRange(num int, index int) bool {
	valid := true
	switch index {
	case 0, 1:
		if num > 59 || num < 0 {
			valid = false
		}
	case 2:
		if num > 23 || num < 0 {
			valid = false
		}
	case 3:
		if num > 31 || num < 1 {
			valid = false
		}
	case 4:
		if num > 12 || num < 1 {
			valid = false
		}
	case 5:
		if num > 7 || num < 1 {
			valid = false
		}
	}

	return valid
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
