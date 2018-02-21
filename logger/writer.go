package logger

import (
	"gotorch/common"
	"gotorch/config"
	"os"
	"time"
)

const SEP = " | "

type writer struct {
	file string
}

var logWriter *writer

func getLogWriter() *writer {
	if logWriter == nil {
		logWriter = &writer{}
	}

	return logWriter
}

func (l *writer) setLogFile(level, pkg string) {
	logDirConfig := config.GetConfig("log_dir")
	logDir := logDirConfig + pkg + "/"
	if !common.IsDirExist(logDir) {
		os.MkdirAll(logDir, 0777)
	}
	fileName := common.GetFileName() + "_" + level + ".log"
	l.file = logDir + fileName
}

func (l *writer) write(level, pkg string, data []string) {
	t := time.Now()

	data = append([]string{"time:" + t.Format(time.UnixDate)}, data...)
	l.setLogFile(level, pkg)

	logContent := common.Join(data, SEP)
	logContent += "\n"

	file, _ := os.OpenFile(l.file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	file.Write([]byte(logContent))
}
