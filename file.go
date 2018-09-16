// Package file provides various single line function for file content operating
package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Buffer can store content to a builder and write to path at once
type Buffer struct {
	path   string
	buffer strings.Builder
}

// NewBuffer create a buffer
func NewBuffer(path string) (buffer *Buffer) {
	buffer = new(Buffer)
	buffer.path = path
	return
}

// Reset buffer content
func (buffer *Buffer) Reset() {
	buffer.buffer.Reset()
	return
}

// Len returns the number of accumulated bytes
func (buffer *Buffer) Len() (len int) {
	len = buffer.buffer.Len()
	return
}

// String returns the accumulated string
func (buffer *Buffer) String() (content string) {
	content = buffer.buffer.String()
	return
}

func (buffer *Buffer) Write(p []byte) (len int, err error) {
	len, err = buffer.buffer.Write(p)
	return
}

func (buffer *Buffer) WriteByte(c byte) (err error) {
	err = buffer.buffer.WriteByte(c)
	return
}

func (buffer *Buffer) WriteRune(r rune) (len int, err error) {
	len, err = buffer.buffer.WriteRune(r)
	return
}

// WriteString appends the contents to buffer
func (buffer *Buffer) WriteString(content ...interface{}) (len int, err error) {
	len, err = buffer.buffer.WriteString(fmt.Sprint(content...))
	return
}

// WriteStringf appends the contents to buffer with format
func (buffer *Buffer) WriteStringf(format string, content ...interface{}) (len int, err error) {
	len, err = buffer.buffer.WriteString(fmt.Sprintf(format, content...))
	return
}

// WriteStringln appends the contents to buffer with newline
func (buffer *Buffer) WriteStringln(content ...interface{}) (len int, err error) {
	len, err = buffer.buffer.WriteString(fmt.Sprintln(content...))
	return
}

// Set set buffer content
func (buffer *Buffer) Set(content ...interface{}) (len int, err error) {
	buffer.Reset()
	len, err = buffer.WriteString(content...)
	return
}

// Setf set buffer content with format
func (buffer *Buffer) Setf(format string, content ...interface{}) (len int, err error) {
	buffer.Reset()
	len, err = buffer.WriteStringf(format, content...)
	return
}

// Setln set buffer content with newline
func (buffer *Buffer) Setln(content ...interface{}) (len int, err error) {
	buffer.Reset()
	len, err = buffer.WriteStringln(content...)
	return
}

// Commit buffer content to file
func (buffer *Buffer) Commit() (len int, err error) {
	len, err = Write(buffer.path, buffer.buffer.String())
	return
}

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
