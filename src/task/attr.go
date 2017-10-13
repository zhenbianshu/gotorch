package task

import (
	"regexp"
	"strconv"
	"strings"
)

type attr struct {
	Script   string
	TaskType string
	Times    string
	Max      int
	Ips      []string
}

var timeType = []string{"second", "minute", "hour", "day", "month", "week"}

func (a attr) timeValid() (isValid bool, errDesc string) {
	times := strings.Split(a.Times, " ")
	if len(times) != 6 {
		return false, "time format error!"
	}

	for index, argStr := range times {
		if argStr == "*" {
			continue
		}

		if ok, _ := regexp.MatchString(`[^*|\-|,|/|a-zA-Z0-9]`, argStr); ok {
			return false, "undefined character!"
		}

		rex, _ := regexp.Compile(`\d`)
		args := rex.FindAllString(argStr, -1)
		for argNum := range args {
			if !inRange(int(argNum), index) {
				return false, "time num out of range"
			}
		}
		// todo check ,/ 并用等异常配置
	}
	return true, ""
}

func (a attr) buildTask() (task *taskItem, errDesc string) {
	if isValid, errDesc := a.timeValid(); !isValid {
		return nil, errDesc
	}

	var mapKey string
	each := make(map[string][]int)
	every := make(map[string]int)
	times := strings.Split(a.Times, " ")
	for index, argStr := range times {
		mapKey = timeType[index]
		every[mapKey] = getEvery(argStr)
		each[mapKey] = getEach(argStr)
	}
	taskInstance := taskItem{every: every, each: each, attr: a}
	return &taskInstance, ""
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

func getEvery(argStr string) int {
	everyPattern, _ := regexp.Compile(`/\d`)
	everyNum := everyPattern.FindString(argStr)
	if everyNum != "" {
		num, _ := strconv.Atoi(everyNum)
		return num
	} else {
		return 1
	}
}

func getEach(argStr string) []int {
	var nums []int
	if strings.IndexAny(argStr, "*") > -1 {
		nums = []int{}
	} else if strings.IndexAny(argStr, ",") >= 0 {
		numList := strings.Split(argStr, ",")
		for num := range numList {
			nums = append(nums, int(num))
		}
	} else if strings.IndexAny(argStr, "-") >= 0 {
		numRange := strings.Split(argStr, "-")
		limit, _ := strconv.Atoi(numRange[1])
		for i, _ := strconv.Atoi(numRange[0]); i <= limit; i++ {
			nums = append(nums, i)
		}
	} else {
		num, _ := strconv.Atoi(argStr)
		nums = append(nums, num)
	}

	return nums
}
