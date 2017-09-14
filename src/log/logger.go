package log

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
)

type logger struct {
	file  string
	level string
}

func (logger *logger) write(data map[string]string, level string) {

}

func (logger *logger) setLevel(data map[string]string, level string) {

}

func (logger *logger) setLogFile() {
	// todo 加载全局配置中的日志目录
	file_name := getPkgName() + "/" + getFileName()
	logger.file = file_name
}

func (logger *logger) setLogLevel(level string) {
	logger.level = level
}

func Debug() {

}

func Warning() {

}

func Error() {

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
