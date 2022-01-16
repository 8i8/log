package log

import (
	"fmt"
	"log"
	"os"
)

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// With the exception of the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// SetFlags sets the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func SetFlags(flag int) {
	log.SetFlags(flag)
}

// SetFlags wraps the standard libraries logger SetFlags function..
func (l *Logger) SetFlags(f int) {
	l.Logger.SetFlags(f)
}

// Print wraps log.Print.
func Print(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
}

// Fatal wraps log.Fatal
func Fatal(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf wraps log.Fatalf
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
