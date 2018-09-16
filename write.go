package file

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

// WriteByte write byte to file
func WriteByte(path string, content []byte) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			}),
		).Debugln("WriteByte to file . . .")
	}

	file, err := os.Create(path)
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
		).Debugln("WriteByte to file")
	}

	return
}

// Write write to file
func Write(path string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			}),
		).Debugln("Write to file . . .")
	}

	file, err := os.Create(path)
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
		).Debugln("Write to file")
	}

	return
}

// Writef write to file with format
func Writef(path string, format string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"format":      format,
				"raw_content": content,
			}),
		).Debugln("Writef to file . . .")
	}

	file, err := os.Create(path)
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
		).Debugln("Writef to file")
	}

	return
}

// Writeln write to file with newline
func Writeln(path string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"path":        path,
				"raw_content": content,
			}),
		).Debugln("Writeln to file . . .")
	}

	file, err := os.Create(path)
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
		).Debugln("Writeln to file")
	}

	return
}
