package main

import (
	"task"
	"time"
)

func main() {
	for {
		task.Run()
		time.Sleep(time.Millisecond * 200)
	}
}
