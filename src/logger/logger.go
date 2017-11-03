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

func Error(err string) {
	info := map[string]string{"error": err}
	logger := getLogWriter()
	logger.write(info, LOG_LEVEL_ERROR)
}
