package log

import "sync/atomic"

//go:generate stringer -type=LogLevel
var level = USER

const (
	// TRACE is the most verbose log level setting giving fine
	// detals of variable values and every step of the progams
	// control flow.
	TRACE Level = 1 << iota
	// DEBUG displays information that is helpful for most debugging
	// cases.
	DEBUG
	// USER logs error messages that are sent to the user.
	USER
	// INFO is the standard default logging level setting, displying
	// the basic program routine.
	INFO
	// ERROR is, or at least should always be output to the log,
	// this is the reccomended minimum setting.
	ERROR
	// SYSTEM is a system level error, something really bad has
	// happened!
	SYSTEM
	// NONE stops all logging, useful sometimes whilst testing.
	NONE
)

// Is returns true of the given level is above the global
// loglevel.
func Is(v Level) bool {
	if uint64(v) >= atomic.LoadUint64((*uint64)(&level)) {
		return true
	}
	return false
}

// Level returns true of the given level is above the loggers
// loglevel.
func (l *Logger) Level(v Level) bool {
	if uint64(v) >= atomic.LoadUint64((*uint64)(&l.level)) {
		return true
	}
	return false
}

// GetLevel returns the global loggers current level.
func GetLevel() Level {
	return Level(atomic.LoadUint64((*uint64)(&level)))
}

// GetLevel returns the given loggers current level.
func (l Logger) GetLevel() Level {
	return Level(atomic.LoadUint64((*uint64)(&level)))
}

// SetLevel sets the packages global loggers logging level.
func SetLevel(l Level) Level {
	prev := level
	atomic.StoreUint64((*uint64)(&level), uint64(l))
	return prev
}

// SetLevel sets the loggers logging level.
func (l *Logger) SetLevel(level Level) Level {
	prev := l.level
	atomic.StoreUint64((*uint64)(&l.level), uint64(level))
	return prev
}
