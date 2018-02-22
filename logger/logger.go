package logger

const LOG_LEVEL_DEBUG = "debug"
const LOG_LEVEL_WARNING = "warning"
const LOG_LEVEL_ERROR = "error"

// normal log
func Debug(pkg string, info ...string) {
	logger := getLogWriter()
	logger.write(LOG_LEVEL_DEBUG, pkg, info)
}

// warning log
func Warning(pkg string, info ...string) {
	logger := getLogWriter()
	logger.write(LOG_LEVEL_WARNING, pkg, info)
}

// fatal error log
func Error(err string) {
	info := []string{err}
	logger := getLogWriter()
	logger.write(LOG_LEVEL_ERROR, "error", info)
}
