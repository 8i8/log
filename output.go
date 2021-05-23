package log

import (
	"fmt"
	"log"

	"github.com/8i8/term"
)

// colour is the default output colour that is returned to after text
// has been wrapped with a colour tag.
var colour = term.White

// User provides a standardised logging output.
func User(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "USER"
	write(3, id, lev, action, fname, event, args...)
}

// User provides a standardised logging output.
func (l Logger) User(id interface{}, action, fname, event string, args ...interface{}) {
	const lev = "USER"
	writelog(3, id, lev, action, fname, event, args...)
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

// writelog provides a standardised log output.
func writelog(depth int, id interface{}, lev, action, fname, event string, args ...interface{}) {
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
