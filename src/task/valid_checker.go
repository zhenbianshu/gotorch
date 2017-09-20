package task

import (
	"regexp"
)

func checkValid(task_config []string) bool {
	if len(task_config) < 7 {
		return false
	}

	for index, arg_str := range task_config[:5] {
		if arg_str == "*" {
			continue
		}

		if regexp.MatchString("[^*|\\-|,|/|a-zA-Z0-9]", arg_str) {
			return false
		}

		rex, _ := regexp.Compile("\\d")
		args := rex.FindAllString(arg_str, -1)
		for arg_num := range args {
			if !inRange(int(arg_num), index) {
				return false
			}
		}

	}
	return true
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
