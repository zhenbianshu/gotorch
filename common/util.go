package common

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strings"
)

// get current file name
func GetFileName() string {
	_, sourcePath, _, _ := runtime.Caller(4)
	rootDir, _ := os.Getwd()

	filename := strings.TrimSuffix(strings.TrimPrefix(sourcePath, rootDir+"/"), ".go")

	return filename
}

// get a package name
func GetPkgName() string {
	_, filePath, _, _ := runtime.Caller(0)
	file, _ := os.Open(filePath)
	r := bufio.NewReader(file)
	line, _, _ := r.ReadLine()
	pkgName := bytes.TrimPrefix(line, []byte("package "))

	return string(pkgName)
}

// check directory exist
func IsDirExist(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

// check file exist
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// implode the strings
func Join(data []string, sep string) string {
	if len(data) <= 0 {
		return ""
	} else if len(data) == 1 {
		return data[0]
	}

	str := data[0]
	data = data[1:]
	for _, words := range data {
		str += sep + words
	}

	return str
}
