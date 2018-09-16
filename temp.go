package file

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

// WriteByteTemp write byte to temp file
func WriteByteTemp(pattern string, content []byte) (path string, len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"pattern":     pattern,
				"raw_content": content,
			}),
		).Debugln("WriteByteTemp to file . . .")
	}

	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot create file")
		}
		return
	}
	path = file.Name()
	defer file.Close()

	len, err = file.Write(content)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":           path,
				"content":        content,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("WriteByteTemp to file")
	}

	return
}

// WriteTemp write to temp file
func WriteTemp(pattern string, content ...interface{}) (path string, len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"pattern":     pattern,
				"raw_content": content,
			}),
		).Debugln("WriteTemp to file . . .")
	}

	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot create file")
		}
		return
	}
	path = file.Name()
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
		).Debugln("WriteTemp to file")
	}

	return
}

// WritefTemp write to temp file with format
func WritefTemp(pattern string, format string, content ...interface{}) (path string, len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"pattern":     pattern,
				"format":      format,
				"raw_content": content,
			}),
		).Debugln("WritefTemp to file . . .")
	}

	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot create file")
		}
		return
	}
	path = file.Name()
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
		).Debugln("WritefTemp to file")
	}

	return
}

// WritelnTemp write to temp file with newline
func WritelnTemp(pattern string, content ...interface{}) (path string, len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"pattern":     pattern,
				"raw_content": content,
			}),
		).Debugln("WritelnTemp to file . . .")
	}

	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot create file")
		}
		return
	}
	path = file.Name()
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
		).Debugln("WritelnTemp to file")
	}

	return
}
