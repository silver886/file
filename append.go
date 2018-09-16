package file

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

// Append appends the contents to file
func Append(path string, content ...interface{}) (len int, err error) {
	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			})),
			logrus.DebugLevel,
			"Append to file . . .",
		)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			logger.LevelLog(intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				})),
				logrus.ErrorLevel,
				"Cannot open file",
			)
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprint(content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			})),
			logrus.DebugLevel,
			"Append to file",
		)
	}

	return
}

// Appendf appends the contents to file with format
func Appendf(path string, format string, content ...interface{}) (len int, err error) {
	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"format":      format,
				"raw_content": content,
			})),
			logrus.DebugLevel,
			"Appendf to file . . .",
		)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			logger.LevelLog(intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				})),
				logrus.ErrorLevel,
				"Cannot open file",
			)
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintf(format, content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			})),
			logrus.DebugLevel,
			"Appendf to file",
		)
	}

	return
}

// Appendln appends the contents to file with newline
func Appendln(path string, content ...interface{}) (len int, err error) {
	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			})),
			logrus.DebugLevel,
			"Appendln to file . . .",
		)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			logger.LevelLog(intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				})),
				logrus.ErrorLevel,
				"Cannot open file",
			)
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintln(content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		logger.LevelLog(intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			})),
			logrus.DebugLevel,
			"Appendln to file",
		)
	}

	return
}
