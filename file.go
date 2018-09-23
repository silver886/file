package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

var (
	intLog    bool
	intLogger *logger.Entry
)

// SetLogger set internal logger for logging
func SetLogger(extLogger *logger.Logger) {
	intLogger = extLogger.WithField("prefix", "file")
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

// Copy copy files
func Copy(dest string, src ...string) error {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"destination": dest,
				"source":      src,
			}),
		).Debugln("Copy file . . .")
	}

	var fullFileList []string
	var errorMsg strings.Builder

	for _, val := range src {

		var patternBuilder strings.Builder

		for _, val := range strings.Split(val, "/") {
			if strings.Contains(val, "*") {
				for _, r := range val {
					if unicode.IsLetter(r) {
						patternBuilder.WriteString(fmt.Sprintf("[%c%c]", unicode.ToLower(r), unicode.ToUpper(r)))
					} else {
						patternBuilder.WriteString(string(r))
					}
				}
			} else {
				patternBuilder.WriteString(val)
			}
			patternBuilder.WriteString("/")
		}
		pattern := strings.TrimRight(patternBuilder.String(), "/")

		if files, err := filepath.Glob(pattern); err != nil {
			if intLog {
				intLogger.WithFields(
					logger.DebugInfo(1, logrus.Fields{
						"pattern": pattern,
					}),
				).WithError(err).Warnln("Cannot find file")
				errorMsg.WriteString("Cannot find file: ")
				errorMsg.WriteString(pattern)
				errorMsg.WriteString(", reason: ")
				errorMsg.WriteString(err.Error())
				errorMsg.WriteString("; ")
			}
		} else {
			fullFileList = append(fullFileList, files...)
		}
	}

	for _, val := range fullFileList {
		systemsavPath := strings.Split(val, string(os.PathSeparator))
		dest := dest + "/" + systemsavPath[len(systemsavPath)-1]
		if err := copy.Copy(val, dest); err != nil {
			if intLog {
				intLogger.WithFields(
					logger.DebugInfo(1, logrus.Fields{
						"file_path": val,
					}),
				).WithError(err).Warnln("Cannot copy file")
				errorMsg.WriteString("Cannot copy file: ")
				errorMsg.WriteString(val)
				errorMsg.WriteString(", reason: ")
				errorMsg.WriteString(err.Error())
				errorMsg.WriteString("; ")
			}
		}
	}

	if errorMsg.Len() == 0 {
		return nil
	} else {
		return errors.New(errorMsg.String())
	}
}
