package log

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"sync/atomic"

	"github.com/8i8/term"
	"github.com/google/uuid"
)

func init() {
	SetLevel(INFO)
	log.SetFlags(log.Lshortfile)
}

// Identifier defines an interface for gathering displaying user id.
type Identifier interface {
	URL() *url.URL
	IP() net.IP
	SID() uuid.UUID
}

type LogLevel uint64

//go:generate stringer -type=LogLevel
var level LogLevel

var (
	colour = term.White
)

const (
	// TRACE is the most verbose log level setting giving fine
	// detals of variable values and every step of the progams
	// control flow.
	TRACE LogLevel = 1 << iota
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

// Level returns true of the given level is above the package
// loglevel.
func Level(v LogLevel) bool {
	if uint64(v) >= atomic.LoadUint64((*uint64)(&level)) {
		return true
	}
	return false
}

// GetLevel returns the given loggers current level.
func GetLevel() LogLevel {
	return LogLevel(atomic.LoadUint64((*uint64)(&level)))
}

// SetLevel sets the package logging level, accepting a string as input
// to simplify use from the command line.
func SetLevel(l LogLevel) LogLevel {
	prev := level
	atomic.StoreUint64((*uint64)(&level), uint64(l))
	return prev
}

// Print wraps log.Print.
func Print(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
}

// Fatal wrap log.Fatal
func Fatal(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatal wrap log.Fatalf
func Fatalf(format string, v ...interface{}) {
	log.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Printf wraps log.Printf.
func Printf(format string, v ...interface{}) {
	log.Output(2, fmt.Sprintf(format, v...))
}

// Println wraps log.Println.
func Println(v ...interface{}) {
	log.Output(2, fmt.Sprintln(v...))
}

// User provides a standardised logging output.
func User(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "USER"
	write(3, id, lev, action, fname, event, args...)
}

// Info provides a standardised logging output.
func Info(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "INFO"
	write(3, id, lev, action, fname, event, args...)
}

// Debug provides a standardised logging output.
func Debug(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "DEBUG"
	write(3, id, lev, action, fname, event, args...)
}

// DebugDepth provides a standardised logging output for a nested
// debugging log call.
func DebugDepth(d int, id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "DEBUG"
	write(d, id, lev, action, fname, event, args...)
}

// Err provides a standardised logging output.
func Err(id interface{}, action, fname, event string, args ...interface{}) {
	var lev = colour.Red("ERROR")
	write(3, id, lev, action, fname, colour.Red(event), args...)
}

// ErrDepth provides a standardised logging output for a nested error
// call.
func ErrDepth(d int, id interface{}, action, fname, event string, args ...interface{}) {
	var lev = colour.Red("ERROR")
	write(3+d, id, lev, action, fname, event, args...)
}

// Trace provides a standardised logging output.
func Trace(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "TRACE"
	write(3, id, lev, action, fname, event, args...)
}

// Sys provides a standardised logging output.
func Sys(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "SYSTEM"
	write(3, id, lev, action, fname, event, args...)
}

// write provides a standardised log output.
func write(depth int, id interface{}, lev, action, fname, event string, args ...interface{}) {
	if len(args)&1 > 0 {
		log.Output(3, fmt.Sprintf(
			"%s: need pairs of arguments %q", lev, args))
	}

	var ok bool
	var ident Identifier
	if id != nil {
		switch id.(type) {
		case Identifier:
			ok = true
			ident = id.(Identifier)
		}
	}
	if ok {
		path, host, ip := ident.URL().Path, ident.URL().Host, ident.IP()
		switch len(args) {
		case 0:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s]",
				colour.White(lev), host, path, ip, action, fname, event))
		case 2:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v]",
				colour.White(lev), host, path, ip, action, fname, event, args[0], args[1]))
		case 4:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
				colour.White(lev), host, path, ip, action, fname, event, args[0], args[1],
				args[2], args[3]))
		case 6:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
				colour.White(lev), host, path, ip, action, fname, event, args[0], args[1],
				args[2], args[3], args[4], args[5]))
		default:
			log.Output(2, fmt.Sprint("default reached in log write() ok"))
		}
		return
	}

	switch len(args) {
	case 0:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s]",
			colour.White(lev), action, fname, event))
	case 2:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v]",
			colour.White(lev), action, fname, event, args[0], args[1]))
	case 4:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
			colour.White(lev), action, fname, event, args[0], args[1],
			args[2], args[3]))
	case 6:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
			colour.White(lev), action, fname, event, args[0], args[1],
			args[2], args[3], args[4], args[5]))
	default:
		log.Output(depth, fmt.Sprint("default reached reached in log write()"))
	}
}
