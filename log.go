package log

import (
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/google/uuid"
)

// Logger wraps a native log and a logLevel to simplfy the passing
// between packages.
type Logger struct {
	// log.Logger is the wrapped logger form the standard log
	// library.
	*log.Logger
	// level is the level of verbosity of the logger, this graduates
	// the log output into the folowing catagories:
	//
	// TRACE : The most verbose log level setting giving fine
	// detals of variable values and every step of the progams
	// control flow.
	//
	// DEBUG : displays information that is helpful for most
	// debugging cases.
	//
	// USER : logs error messages that are sent to the user as well
	// as general information, this is the default setting.
	//
	// INFO : displays the programs routines output.
	//
	// ERROR : is, or at least should always be output to the log,
	// this is the reccomended minimum setting.
	//
	// SYSTEM is a system level error, something really bad has
	// happened!
	//
	// NONE stops all logging, useful sometimes whilst testing.
	level Level
}

// Level stors the logging level of a logger.
type Level uint64

// Identifier defines an interface for gathering displaying user id.
type Identifier interface {
	URL() *url.URL
	IP() string
	SID() uuid.UUID
}

func init() {
	// Set the default global log setting.
	log.SetFlags(log.LstdFlags)
}

// New returns a LogLevel variable with your provided logger attached.
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
		level:  USER,
	}
}

// Stop is an empty function used to halt the delve debugger.
func Stop(fname ...string) {
	fmt.Println("stop:", fname)
}
