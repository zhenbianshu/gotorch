package common

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strings"
)

func GetFileName() string {
	_, sourcePath, _, _ := runtime.Caller(4)
	rootDir, _ := os.Getwd()

	filename := strings.TrimSuffix(strings.TrimPrefix(sourcePath, rootDir+"/"), ".go")

	return filename
}

func GetPkgName() string {
	_, filePath, _, _ := runtime.Caller(0)
	file, _ := os.Open(filePath)
	r := bufio.NewReader(file)
	line, _, _ := r.ReadLine()
	pkgName := bytes.TrimPrefix(line, []byte("package "))

	return string(pkgName)
}

func IsDirExist(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}
