package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold != LevelInfo {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Warnf(format string, args ...any) {
		if l.threshold != LevelWarn {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold != LevelError {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// logf prints the message to the output
func (l *Logger) logf(format string, args ...any) {
		// making sure we can safely write tothe output
	if l.output == nil {
		l.output = os.Stdout
	}
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}