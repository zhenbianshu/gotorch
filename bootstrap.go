package main

import (
	"task"
	"time"
)

func main() {
	task.Init()
	for {
		task.Run()
		time.Sleep(time.Millisecond * 200)
	}
}
