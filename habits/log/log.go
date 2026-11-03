package log

import (
	"io"
	"log"
	"sync"
)

// A logger that can log messages
type Logger struct {
	mutex sync.Mutex
	Logger *log.Logger
}

// New returns a logger
func New(output io.Writer) *Logger {
	return &Logger{
		Logger: log.New(output, "", log.Ldate | log.Ltime),
	}
}

// Logf sends a message to the log if the severity is high enough
func (l *Logger) Logf(format string, args ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.Logger.Printf(format, args...)
}