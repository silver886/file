package file

import (
	"fmt"
	"os"
)

// Append appends the contents to file
func Append(path string, content ...interface{}) (int, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprint(file, fmt.Sprint(content...))
}

// Appendf appends the contents to file with format
func Appendf(path string, format string, content ...interface{}) (int, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprint(file, fmt.Sprintf(format, content...))
}

// Appendln appends the contents to file with newline
func Appendln(path string, content ...interface{}) (int, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644|os.ModeAppend)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprint(file, fmt.Sprintln(content...))
}
