package log

import (
	"fmt"
	"log"

	"github.com/8i8/term"
)

// colour holds a colour from the term library.
var colour term.Colour

var ColourErrors bool

func what(v interface{}) string {
	return fmt.Sprintf("!(%T, %v)", v, v)
}

// toString returns a string when given either an error or a string.
func toString(v interface{}) (str string) {
	switch t := v.(type) {
	case fmt.Stringer:
		str = t.String()
	case error:
		str = t.Error()
	case string:
		str = t
	default:
		panic("unknown type: "+what(v))
	}
	return
}

// User provides a standardised logging output.
func User(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "USER"
	write(3, id, lev, action, fname, event, args...)
}

// User provides a standardised logging output.
func (l Logger) User(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "USER"
	writelog(3, id, lev, action, fname, event, args...)
}

// Info provides a standardised logging output.
func Info(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "INFO"
	write(3, id, lev, action, fname, event, args...)
}

// Debug provides a standardised logging output.
func Debug(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "DEBUG"
	write(3, id, lev, action, fname, event, args...)
}

// DebugDepth provides a standardised logging output for a nested
// debugging log call.
func DebugDepth(d int, id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "DEBUG"
	write(d, id, lev, action, fname, event, args...)
}

// Err provides a standardised logging output.
func Err(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	var lev = "ERROR"
	evnt := toString(event)
	if ColourErrors {
		lev = colour.Red(lev)
		event = colour.Red(evnt)
	}
	write(3, id, lev, action, fname, evnt, args...)
}

// ErrDepth provides a standardised logging output for a nested error
// call.
func ErrDepth(d int, id interface{}, action, fname string, event interface{}, args ...interface{}) {
	var lev = "ERROR"
	evnt := toString(event)
	if ColourErrors {
		lev = colour.Red("ERROR")
		evnt = colour.Red(evnt)
	}
	write(3+d, id, lev, action, fname, evnt, args...)
}

// Trace provides a standardised logging output.
func Trace(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "TRACE"
	write(3, id, lev, action, fname, event, args...)
}

// Sys provides a standardised logging output.
func Sys(id interface{}, action, fname string, event interface{}, args ...interface{}) {
	const lev = "SYSTEM"
	write(3, id, lev, action, fname, event, args...)
}

// write provides a standardised log output.
func write(depth int, id interface{}, lev, action, fname string, v interface{}, args ...interface{}) {
	event := toString(v)
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
				lev, host, path, ip, action, fname, event))
		case 2:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1]))
		case 4:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1],
				args[2], args[3]))
		case 6:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1],
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
			lev, action, fname, event))
	case 2:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v]",
			lev, action, fname, event, args[0], args[1]))
	case 4:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
			lev, action, fname, event, args[0], args[1],
			args[2], args[3]))
	case 6:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
			lev, action, fname, event, args[0], args[1],
			args[2], args[3], args[4], args[5]))
	default:
		log.Output(depth, fmt.Sprint("default reached reached in log write()"))
	}
}

// writelog provides a standardised log output.
func writelog(depth int, id interface{}, lev, action, fname string, event interface{}, args ...interface{}) {
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
				lev, host, path, ip, action, fname, event))
		case 2:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1]))
		case 4:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1],
				args[2], args[3]))
		case 6:
			log.Output(depth, fmt.Sprintf(
				"%s:[host:%s][path:%s][ip:%s][action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
				lev, host, path, ip, action, fname, event, args[0], args[1],
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
			lev, action, fname, event))
	case 2:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v]",
			lev, action, fname, event, args[0], args[1]))
	case 4:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v]",
			lev, action, fname, event, args[0], args[1],
			args[2], args[3]))
	case 6:
		log.Output(depth, fmt.Sprintf(
			"%s:[action:%s][fname:%s][event:%s][%s:%v][%s:%v][%s:%v]",
			lev, action, fname, event, args[0], args[1],
			args[2], args[3], args[4], args[5]))
	default:
		log.Output(depth, fmt.Sprint("default reached reached in log write()"))
	}
}
