package file

import (
	"fmt"
	"os"
)

// WriteByte write byte to file
func WriteByte(path string, content []byte) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(content)
}

// Write write to file
func Write(path string, content ...interface{}) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprint(file, content...)
}

// Writef write to file with format
func Writef(path string, format string, content ...interface{}) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprintf(file, format, content...)
}

// Writeln write to file with newline
func Writeln(path string, content ...interface{}) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return fmt.Fprintln(file, content...)
}
