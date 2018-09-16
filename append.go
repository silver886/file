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
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			}),
		).Debugln("Append to file . . .")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot open file")
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprint(content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("Append to file")

	}

	return
}

// Appendf appends the contents to file with format
func Appendf(path string, format string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"format":      format,
				"raw_content": content,
			}),
		).Debugln("Appendf to file . . .")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot open file")
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintf(format, content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("Appendf to file")
	}

	return
}

// Appendln appends the contents to file with newline
func Appendln(path string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			}),
		).Debugln("Appendln to file . . .")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot open file")
		}
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintln(content...)
	len, err = fmt.Fprint(file, contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("Appendln to file")
	}

	return
}
