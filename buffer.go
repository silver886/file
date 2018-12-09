package file

import (
	"fmt"
	"strings"
)

// Buffer can store contents to a buffer and write them to path at one time
type Buffer struct {
	path   string
	buffer strings.Builder
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

// Write appends the contents to buffer
func (buffer *Buffer) Write(content []byte) (len int, err error) {
	len, err = buffer.buffer.Write(content)

	return
}

// WriteByte appends the contents to buffer
func (buffer *Buffer) WriteByte(content byte) (err error) {
	err = buffer.buffer.WriteByte(content)

	return
}

// WriteRune appends the UTF-8 encoding of Unicode code content to buffer
func (buffer *Buffer) WriteRune(content rune) (len int, err error) {
	len, err = buffer.buffer.WriteRune(content)

	return
}

// WriteString appends the contents to buffer
func (buffer *Buffer) WriteString(content ...interface{}) (len int, err error) {
	contentStr := fmt.Sprint(content...)
	len, err = buffer.buffer.WriteString(contentStr)

	return
}

// WriteStringf appends the contents to buffer with format
func (buffer *Buffer) WriteStringf(format string, content ...interface{}) (len int, err error) {
	contentStr := fmt.Sprintf(format, content...)
	len, err = buffer.buffer.WriteString(contentStr)

	return
}

// WriteStringln appends the contents to buffer with newline
func (buffer *Buffer) WriteStringln(content ...interface{}) (len int, err error) {

	contentStr := fmt.Sprintln(content...)
	len, err = buffer.buffer.WriteString(contentStr)

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

// NewBuffer create a buffer
func NewBuffer(path string) (buffer *Buffer) {
	buffer = new(Buffer)
	buffer.path = path

	return
}
