package pocketlog

// Level represents an available logging level
type Level byte

const (
	// represents the lowest level of log
	LevelDebug Level = iota
	// contains information
	LevelInfo
	// warnings, safe to ignore
	LevelWarn
	// highest level, only used to trace errors
	LevelError
)