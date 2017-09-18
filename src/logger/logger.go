package logger

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strings"
)

const LOG_LEVEL_DEBUG = "debug"
const LOG_LEVEL_WARNING = "warning"
const LOG_LEVEL_ERROR = "error"

func Debug(info map[string]string) {
	logger := getLogWriter()
	logger.write(info, LOG_LEVEL_DEBUG)
}

func Warning(info map[string]string) {
	logger := getLogWriter()
	logger.write(info, LOG_LEVEL_WARNING)
}

func Error(info map[string]string) {
	logger := getLogWriter()
	logger.write(info, LOG_LEVEL_ERROR)
}

func getFileName() string {
	_, _source_path, _, _ := runtime.Caller(4)
	root_dir, _ := os.Getwd()

	filename := strings.TrimSuffix(strings.TrimPrefix(_source_path, root_dir+"/"), ".go")

	return filename
}

func getPkgName() string {
	_, file_path, _, _ := runtime.Caller(0)
	file, _ := os.Open(file_path)
	r := bufio.NewReader(file)
	line, _, _ := r.ReadLine()
	pkg_name := bytes.TrimPrefix(line, []byte("package "))

	return string(pkg_name)
}

func isDirExist(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
