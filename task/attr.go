package task

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type attr struct {
	Command  string
	TaskType string
	Times    string
	Max      int
	Ips      []string
}

const TypeDaemon = "daemon"
const TypeCommon = "common"

// check if time config is valid
func (a attr) timeValid() (isValid bool, err error) {
	times := strings.Split(a.Times, " ")
	if len(times) != 6 {
		err = errors.New("time format error")
		return false, err
	}

	for index, argStr := range times {
		if argStr == "*" {
			continue
		}

		if ok, _ := regexp.MatchString(`[^*|\-|,|/|a-zA-Z0-9]`, argStr); ok {
			err = errors.New("undefined character")
			return false, err
		}

		rex, _ := regexp.Compile(`\d`)
		args := rex.FindAllString(argStr, -1)
		for argNum := range args {
			if !inRange(int(argNum), index) {
				err = errors.New("time num out of range")
				return false, err
			}
		}
	}
	return true, nil
}

// build a task item with a attr config
func (a attr) buildTask() (task *taskItem, err error) {
	if isValid, err := a.timeValid(); !isValid {
		return nil, err
	}

	times := make(map[int][]bool)
	timeConf := strings.Split(a.Times, " ")
	for index, argStr := range timeConf {
		limitList, err := parseTimeRange(argStr, index)
		if err != nil {
			return nil, err
		}
		times[index] = limitList
	}

	cmdArgs := strings.Split(a.Command, " ")
	cmd := cmdArgs[0]
	args := cmdArgs[1:]

	taskInstance := taskItem{times: times, attr: a, cmd: cmd, args: args}

	return &taskInstance, nil
}

func inRange(num int, index int) bool {
	valid := true
	start, end := getRange(index)
	if num > end || num < start {
		valid = false
	}

	return valid
}

// parse time config and get time point
func parseTimeRange(argStr string, index int) ([]bool, error) {
	parseErr := errors.New("time range parse error")

	_, max := getRange(index)
	points := make([]bool, max+1)
	argParts := strings.Split(argStr, ",")
	for _, argPart := range argParts {
		eachArg := strings.Split(argPart, "/")
		if len(eachArg) > 2 {
			return []bool{}, parseErr
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
				return []bool{}, parseErr
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
			points[start] = true
		}
	}

	return points, nil
}

// get time range by type
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
