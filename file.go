// Package file provides various single line function for file content operating
package file

import (
	"fmt"
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

// Append append to file
func Append(path string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprint(file, content...)

	return
}

// Appendf append to file with format
func Appendf(path string, format string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintf(file, format, content...)

	return
}

// Appendln append to file with newline
func Appendln(path string, content ...interface{}) (len int, err error) {
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	len, err = fmt.Fprintln(file, content...)

	return
}

// WriteByteTemp write byte to temp file
func WriteByteTemp(pattern string, content []byte) (path string, len int, err error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return
	}
	path = file.Name()
	defer file.Close()

	len, err = file.Write(content)

	return
}

// WriteTemp write to temp file
func WriteTemp(pattern string, content ...interface{}) (path string, len int, err error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return
	}
	path = file.Name()
	defer file.Close()

	len, err = fmt.Fprint(file, content...)

	return
}

// WritefTemp write to temp file with format
func WritefTemp(pattern string, format string, content ...interface{}) (path string, len int, err error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return
	}
	path = file.Name()
	defer file.Close()

	len, err = fmt.Fprintf(file, format, content...)

	return
}

// WritelnTemp write to temp file with newline
func WritelnTemp(pattern string, content ...interface{}) (path string, len int, err error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return
	}
	path = file.Name()
	defer file.Close()

	len, err = fmt.Fprintln(file, content...)

	return
}
