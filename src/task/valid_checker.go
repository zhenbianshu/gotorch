package task

func checkValid(task_config []string) bool {
	if len(task_config) < 7 {
		return false
	}

	return true
}
