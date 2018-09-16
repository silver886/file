package file

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"leoliu.io/logger"
)

// Buffer can store contents to a buffer and write them to path at one time
type Buffer struct {
	path   string
	buffer strings.Builder
}

// Reset buffer content
func (buffer *Buffer) Reset() {
	buffer.buffer.Reset()

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer": buffer,
			}),
		).Debugln("Reset buffer")
	}

	return
}

// Len returns the number of accumulated bytes
func (buffer *Buffer) Len() (len int) {
	len = buffer.buffer.Len()

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer": buffer,
				"length": len,
			}),
		).Debugln("Buffer length")
	}

	return
}

// String returns the accumulated string
func (buffer *Buffer) String() (content string) {
	content = buffer.buffer.String()

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer": buffer,
				"string": content,
			}),
		).Debugln("Buffer content")
	}

	return
}

// Write appends the contents to buffer
func (buffer *Buffer) Write(content []byte) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("Write to buffer . . .")
	}

	len, err = buffer.buffer.Write(content)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        content,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("Write to buffer")
	}
	return
}

// WriteByte appends the contents to buffer
func (buffer *Buffer) WriteByte(content byte) (err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("WriteByte to buffer . . .")
	}

	err = buffer.buffer.WriteByte(content)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        content,
				"internal_error": err,
			}),
		).Debugln("WriteByte to buffer")
	}

	return
}

// WriteRune appends the UTF-8 encoding of Unicode code content to buffer
func (buffer *Buffer) WriteRune(content rune) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("WriteRune to buffer . . .")
	}

	len, err = buffer.buffer.WriteRune(content)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        content,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("WriteRune to buffer")
	}

	return
}

// WriteString appends the contents to buffer
func (buffer *Buffer) WriteString(content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("WriteString to buffer . . .")
	}

	contentStr := fmt.Sprint(content...)
	len, err = buffer.buffer.WriteString(contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("WriteString to buffer")
	}

	return
}

// WriteStringf appends the contents to buffer with format
func (buffer *Buffer) WriteStringf(format string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"format":      format,
				"raw_content": content,
			}),
		).Debugln("WriteStringf to buffer . . .")
	}

	contentStr := fmt.Sprintf(format, content...)
	len, err = buffer.buffer.WriteString(contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("WriteStringf to buffer")
	}

	return
}

// WriteStringln appends the contents to buffer with newline
func (buffer *Buffer) WriteStringln(content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("WriteStringln to buffer . . .")
	}

	contentStr := fmt.Sprintln(content...)
	len, err = buffer.buffer.WriteString(contentStr)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":         buffer,
				"content":        contentStr,
				"length":         len,
				"internal_error": err,
			}),
		).Debugln("WriteStringln to buffer")
	}

	return
}

// Set set buffer content
func (buffer *Buffer) Set(content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("Set buffer . . .")
	}

	buffer.Reset()
	len, err = buffer.WriteString(content...)
	return
}

// Setf set buffer content with format
func (buffer *Buffer) Setf(format string, content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"format":      format,
				"raw_content": content,
			}),
		).Debugln("Setf buffer . . .")
	}

	buffer.Reset()
	len, err = buffer.WriteStringf(format, content...)
	return
}

// Setln set buffer content with newline
func (buffer *Buffer) Setln(content ...interface{}) (len int, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer":      buffer,
				"raw_content": content,
			}),
		).Debugln("Setln buffer . . .")
	}

	buffer.Reset()
	len, err = buffer.WriteStringln(content...)
	return
}

// Commit buffer content to file
func (buffer *Buffer) Commit() (len int, err error) {
	len, err = Write(buffer.path, buffer.buffer.String())

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer": buffer,
			}),
		).Debugln("Commit buffer")
	}

	return
}

// NewBuffer create a buffer
func NewBuffer(path string) (buffer *Buffer) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"file_path": path,
			}),
		).Debugln("New buffer . . .")
	}

	buffer = new(Buffer)
	buffer.path = path

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"buffer": buffer,
			}),
		).Debugln("New buffer")
	}

	return
}
