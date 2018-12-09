package file

import (
	"fmt"
	"io/ioutil"
)

// WriteByteTemp write byte to temp file
func WriteByteTemp(pattern string, content []byte) (string, int, error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	len, err := file.Write(content)
	return file.Name(), len, err
}

// WriteTemp write to temp file
func WriteTemp(pattern string, content ...interface{}) (string, int, error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	len, err := fmt.Fprint(file, content...)
	return file.Name(), len, err
}

// WritefTemp write to temp file with format
func WritefTemp(pattern string, format string, content ...interface{}) (string, int, error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	len, err := fmt.Fprintf(file, format, content...)
	return file.Name(), len, err
}

// WritelnTemp write to temp file with newline
func WritelnTemp(pattern string, content ...interface{}) (path string, len int, err error) {
	file, err := ioutil.TempFile("", pattern)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	len, err = fmt.Fprintln(file, content...)
	return file.Name(), len, err
}
