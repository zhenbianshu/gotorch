package logger

import (
	"config"
	"encoding/json"
	"fmt"
	"os"
)

type writer struct {
	file string
}

var log_writer *writer

func getLogWriter() *writer {
	if log_writer == nil {
		log_writer = &writer{}
	}

	return log_writer
}

func (logger *writer) setLogFile(level string) {
	log_dir_config := config.GetConfig("log_dir")
	log_dir := log_dir_config + getPkgName() + "/"
	if !isDirExist(log_dir) {
		os.MkdirAll(log_dir, 0777)
	}
	file_name := getFileName() + "_" + level + ".log"
	fmt.Println(logger.file)
	logger.file = log_dir + file_name
}

func (logger *writer) write(data map[string]string, level string) {
	logger.setLogFile(level)
	log_content, _ := json.Marshal(data)
	log_content = append(log_content, '\n')

	file, _ := os.OpenFile(logger.file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	file.Write(log_content)
}
