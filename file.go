package file

import (
	"io/ioutil"
	"os"
)

// Exist detect the existence of a path
func Exist(path string) (result bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		result = false
	} else {
		result = true
	}
	return
}

// Read read a file
func Read(path string) (content string, err error) {
	contentByte, err := ioutil.ReadFile(path)
	content = string(contentByte)
	return
}
