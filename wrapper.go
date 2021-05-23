package log

import (
	"fmt"
	"log"
	"os"
)

// SetFlags wraps the standard libraries logger SetFlags function..
func (l *Logger) SetFlags(f int) {
	l.Logger.SetFlags(f)
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
