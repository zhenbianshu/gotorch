package logger

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
)

type writer struct {
	file  string
	level string
}

func (logger *writer) write(data map[string]string, level string) {

}

func (logger *writer) setLevel(data map[string]string, level string) {

}

func (logger *writer) setLogFile() {
	// todo 加载全局配置中的日志目录
	file_name := getPkgName() + "/" + getFileName()
	logger.file = file_name
}

func (logger *writer) setLogLevel(level string) {
	logger.level = level
}

func Debug(info map[string]string) {

}

func Warning(info map[string]string) {

}

func Error(info map[string]string) {

}

func getFileName() string {
	_, file_path, _, _ := runtime.Caller(0)
	filename_full := path.Base(file_path)
	suffix := path.Ext(filename_full)
	filename := strings.TrimSuffix(filename_full, suffix)

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
