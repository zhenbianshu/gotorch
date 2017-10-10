package logger

import (
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

func (logger *writer) setLogFile(level string) {
	logDirConfig := config.GetConfig("log_dir")
	logDir := logDirConfig + getPkgName() + "/"
	if !isDirExist(logDir) {
		os.MkdirAll(logDir, 0777)
	}
	fileName := getFileName() + "_" + level + ".log"
	logger.file = logDir + fileName
}

func (logger *writer) write(data map[string]string, level string) {
	t := time.Now()
	data["time"] = t.Format(time.UnixDate)
	logger.setLogFile(level)
	logContent, _ := json.Marshal(data)
	logContent = append(logContent, '\n')

	file, _ := os.OpenFile(logger.file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	file.Write(logContent)
}
