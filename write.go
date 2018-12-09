package file

import (
	"fmt"
	"os"
)

// WriteByte write byte to file
func WriteByte(path string, content []byte) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = file.Write(content)

	return
}

// Write write to file
func Write(path string, content ...interface{}) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	contentStr := fmt.Sprint(content...)
	len, err = fmt.Fprint(file, contentStr)

	return
}

// Writef write to file with format
func Writef(path string, format string, content ...interface{}) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintf(format, content...)
	len, err = fmt.Fprint(file, contentStr)

	return
}

// Writeln write to file with newline
func Writeln(path string, content ...interface{}) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	contentStr := fmt.Sprintln(content...)
	len, err = fmt.Fprint(file, contentStr)

	return
}
