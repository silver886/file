package file

import (
	"fmt"
	"strings"
)

// Buffer can store contents to a buffer and write them to path at one time
type Buffer struct {
	*strings.Builder
	path string
}

// WriteString appends the contents to buffer
func (buffer *Buffer) WriteString(content ...interface{}) (int, error) {
	return buffer.Builder.WriteString(fmt.Sprint(content...))
}

// WriteStringf appends the contents to buffer with format
func (buffer *Buffer) WriteStringf(format string, content ...interface{}) (int, error) {
	return buffer.Builder.WriteString(fmt.Sprintf(format, content...))
}

// WriteStringln appends the contents to buffer with newline
func (buffer *Buffer) WriteStringln(content ...interface{}) (int, error) {
	return buffer.Builder.WriteString(fmt.Sprintln(content...))
}

// Set set buffer content
func (buffer *Buffer) Set(content ...interface{}) (int, error) {
	buffer.Reset()
	return buffer.WriteString(content...)
}

// Setf set buffer content with format
func (buffer *Buffer) Setf(format string, content ...interface{}) (int, error) {
	buffer.Reset()
	return buffer.WriteStringf(format, content...)
}

// Setln set buffer content with newline
func (buffer *Buffer) Setln(content ...interface{}) (int, error) {
	buffer.Reset()
	return buffer.WriteStringln(content...)
}

// Commit buffer content to file
func (buffer *Buffer) Commit() (int, error) {
	return Write(buffer.path, buffer.String())
}

// NewBuffer create a buffer
func NewBuffer(path string) *Buffer {
	return &Buffer{
		Builder: new(strings.Builder),
		path:    path,
	}
}
