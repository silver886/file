package file

import (
	"fmt"
	"os"
)

// Append appends the contents to file
func Append(path string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprint(file, content...)

	return
}

// Appendf appends the contents to file with format
func Appendf(path string, format string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintf(file, format, content...)

	return
}

// Appendln appends the contents to file with newline
func Appendln(path string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintln(file, content...)

	return
}
