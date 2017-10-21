package task

import (
	"errors"
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
	}
	return true, ""
}

func (a attr) buildTask() (task *taskItem, errDesc string) {
	if isValid, errDesc := a.timeValid(); !isValid {
		return nil, errDesc
	}

	times := make(map[int][]int)
	timeConf := strings.Split(a.Times, " ")
	for index, argStr := range timeConf {
		limitList, err := parseTimeRange(argStr, index)
		if err != nil {
			return nil, err.Error()
		}
		times[index] = limitList
	}
	taskInstance := taskItem{times: times, attr: a}
	return &taskInstance, ""
}

func inRange(num int, index int) bool {
	valid := true
	start, end := getRange(index)
	if num > end || num < start {
		valid = false
	}

	return valid
}

func parseTimeRange(argStr string, index int) ([]int, error) {
	parseErr := errors.New("time range parse error")

	points := []int{}
	argParts := strings.Split(argStr, ",")
	for _, argPart := range argParts {
		eachArg := strings.Split(argPart, "/")
		if len(eachArg) > 2 {
			return []int{}, parseErr
		}

		var every int
		if len(eachArg) == 1 {
			every = 1
		} else {
			every, _ = strconv.Atoi(eachArg[1])
		}

		var start, end int
		if strings.IndexAny(eachArg[0], "-") >= 0 {
			numRange := strings.Split(argStr, "-")
			if len(numRange) != 2 {
				return []int{}, parseErr
			}

			start, _ = strconv.Atoi(numRange[0])
			end, _ = strconv.Atoi(numRange[0])
		} else if eachArg[0] == "*" {
			start, end = getRange(index)
		} else {
			num, _ := strconv.Atoi(eachArg[0])
			start = num
			end = num
		}

		for ; start < end; start += every {
			points = append(points, start)
		}
	}

	return points, nil
}

func getRange(index int) (start, end int) {
	switch index {
	case second, minute:
		start = 0
		end = 59
	case hour:
		start = 0
		end = 23
	case day:
		start = 1
		end = 31
	case week:
		start = 1
		end = 7
	case month:
		start = 1
		end = 12
	default:
		start = 0
		end = 0
	}

	return start, end
}
