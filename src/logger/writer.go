package logger

import (
	"config"
	"encoding/json"
	"fmt"
	"io"
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
	if !isFileExist(logger.file) {
		// todo 考虑并发
		os.Create(logger.file)
	}

	log_content, _ := json.Marshal(data)
	log_content = append(log_content, '\n')
	fmt.Println(string(log_content), logger.file)

	file, _ := os.OpenFile(logger.file, os.O_WRONLY, 0644)
	offset, _ := file.Seek(0, io.SeekEnd)
	file.WriteAt(log_content, offset)

}
