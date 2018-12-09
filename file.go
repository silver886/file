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
)

// Exist detect the existence of a path
func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Read read a file
func Read(path string) (string, error) {
	contentByte, err := ioutil.ReadFile(path)
	return string(contentByte), err
}

// Copy copy files
func Copy(dest string, src ...string) error {
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
			errorMsg.WriteString("Cannot find file: ")
			errorMsg.WriteString(pattern)
			errorMsg.WriteString(", reason: ")
			errorMsg.WriteString(err.Error())
			errorMsg.WriteString("; ")
		} else {
			fullFileList = append(fullFileList, files...)
		}
	}

	for _, val := range fullFileList {
		systemsavPath := strings.Split(val, string(os.PathSeparator))
		dest := dest + "/" + systemsavPath[len(systemsavPath)-1]
		if err := copy.Copy(val, dest); err != nil {
			errorMsg.WriteString("Cannot copy file: ")
			errorMsg.WriteString(val)
			errorMsg.WriteString(", reason: ")
			errorMsg.WriteString(err.Error())
			errorMsg.WriteString("; ")
		}
	}

	if errorMsg.Len() != 0 {
		return errors.New(errorMsg.String())
	}

	return nil
}
