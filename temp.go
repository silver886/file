package file

import (
	"fmt"
	"io/ioutil"
)

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
