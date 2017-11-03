package logger

import (
	"common"
	"config"
	"encoding/json"
	"os"
	"time"
)

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

func (l *writer) setLogFile(level string) {
	logDirConfig := config.GetConfig("log_dir")
	logDir := logDirConfig + common.GetPkgName() + "/"
	if !common.IsDirExist(logDir) {
		os.MkdirAll(logDir, 0777)
	}
	fileName := common.GetFileName() + "_" + level + ".log"
	l.file = logDir + fileName
}

func (l *writer) write(data map[string]string, level string) {
	t := time.Now()
	data["time"] = t.Format(time.UnixDate)
	l.setLogFile(level)
	logContent, _ := json.Marshal(data)
	logContent = append(logContent, '\n')

	file, _ := os.OpenFile(l.file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	file.Write(logContent)
}
