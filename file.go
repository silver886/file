package file

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

var (
	intLog    bool
	intLogger *logger.Logger
)

// SetLogger set internal logger for logging
func SetLogger(extLogger *logger.Logger) {
	intLogger = extLogger
	intLog = true
}

// ResetLogger reset internal logger
func ResetLogger() {
	intLogger = nil
	intLog = false
}

// Exist detect the existence of a path
func Exist(path string) (result bool) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"file_path": path,
			}),
		).Debugln("Check path existence . . .")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		result = false
	} else {
		result = true
	}

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"file_path": path,
				"result":    result,
			}),
		).Debugln("Check path existence")
	}

	return
}

// Read read a file
func Read(path string) (content string, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"file_path": path,
			}),
		).Debugln("Read from file . . .")
	}

	contentByte, err := ioutil.ReadFile(path)
	content = string(contentByte)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"file_path":   path,
				"raw_content": contentByte,
			}),
		).Debugln("Read from file")
	}

	return
}
