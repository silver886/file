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

	len, err = fmt.Fprint(file, content...)

	return
}

// Writef write to file with format
func Writef(path string, format string, content ...interface{}) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintf(file, format, content...)

	return
}

// Writeln write to file with newline
func Writeln(path string, content ...interface{}) (len int, err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintln(file, content...)

	return
}
