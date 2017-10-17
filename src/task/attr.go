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
	}
	return true, ""
}

func (a attr) buildTask() (task *taskItem, errDesc string) {
	if isValid, errDesc := a.timeValid(); !isValid {
		return nil, errDesc
	}
	// todo check command executable

	var mapKey string
	times := make(map[string][]timeRange)
	timeConf := strings.Split(a.Times, " ")
	for index, argStr := range timeConf {
		mapKey = timeType[index]
		limitList, err := parseTimeRange(argStr, index)
		if err != nil {
			return nil, err.Error()
		}
		times[mapKey] = limitList
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

func parseTimeRange(argStr string, index int) ([]timeRange, error) {
	parseErr := errors.New("time range parse error")
	argParts := strings.Split(argStr, ",")
	limitList := make([]timeRange, 0)
	for _, argPart := range argParts {
		eachArg := strings.Split(argPart, "/")
		if len(eachArg) > 2 {
			return []timeRange{}, parseErr
		}

		limit := timeRange{}
		if len(eachArg) == 1 {
			limit.every = 1
		} else {
			every, err := strconv.Atoi(eachArg[1])
			if err != nil {
				return []timeRange{}, parseErr
			}
			limit.every = every
		}

		if strings.IndexAny(eachArg[0], "-") >= 0 {
			numRange := strings.Split(argStr, "-")
			if len(numRange) != 2 {
				return []timeRange{}, parseErr
			}

			start, tErr := strconv.Atoi(numRange[0])
			end, dErr := strconv.Atoi(numRange[0])
			if tErr != nil || dErr != nil {
				return []timeRange{}, parseErr
			}
			limit.start = start
			limit.end = end
		} else if eachArg[0] == "*" {
			start, end := getRange(index)
			limit.start = start
			limit.end = end
		} else {
			num, err := strconv.Atoi(eachArg[0])
			if err != nil {
				return []timeRange{}, parseErr
			}
			limit.start = num
			limit.end = num
		}
		limitList = append(limitList, limit)
	}

	return limitList, nil
}

func getRange(index int) (start, end int) {
	switch index {
	case 0, 1:
		start = 0
		end = 59
	case 2:
		start = 0
		end = 23

	case 3:
		start = 1
		end = 31
	case 4:
		start = 1
		end = 12
	case 5:
		start = 1
		end = 7
	default:
		start = 0
		end = 0
	}

	return start, end
}
