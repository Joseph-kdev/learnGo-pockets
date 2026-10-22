package pocketlog

type Logger struct {
	threshold Level
}

//New returns you a logger, ready to log at a specific threshold level
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func(l *Logger) Debugf(format string, args ...any) {

}

func (l *Logger) Infof(format string, args ...any) {

}

func (l *Logger) Warnf(format string, args ...any) {

}
func (l *Logger) Errorf(format string, args ...any) {

}