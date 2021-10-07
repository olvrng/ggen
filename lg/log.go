// lg implements a simple verbosity logger interface.
//
// The default verbosity is 0, and can be changed by setting the environment variable GGEN_LOGGING. For example:
//
//     GGEN_LOGGING=1          : set the current verbosity to 1
//
//     logger.V(0).Printf(...) : print log, since V(0) >  GGEN_LOGGING
//     logger.V(1).Printf(...) : print log, since V(1) == GGEN_LOGGING
//     logger.V(2).Printf(...) : not print, since V(2) >  GGEN_LOGGING
//
// User of ggen package can replace the default logger with their own implementation by implementing the Logger
// interface.
package lg

import (
	"log"
	"os"
	"strconv"
)

type Logger interface {

	// Verbosed checks if the current verbosity level is equal or higher than the param.
	Verbosed(verbosity int) bool

	// V returns a VerbosedLogger, which only outputs log when the current verbosity level is equal or higher than the
	// log line verbosity.
	V(verbosity int) VerbosedLogger
}

type VerbosedLogger interface {
	Printf(format string, args ...interface{})
}

var New func() Logger

var verbosity int

func init() {
	New = newLogger
	verbosity, _ = strconv.Atoi(os.Getenv("GGEN_LOGGING"))
}

type logger int

func (l logger) V(verbosity int) VerbosedLogger {
	return logger(verbosity)
}

func (_ logger) Verbosed(v int) bool {
	return v <= verbosity
}

func (l logger) Printf(format string, args ...interface{}) {
	if int(l) <= verbosity {
		log.Printf(format, args...)
	}
}

func newLogger() Logger {
	return logger(0)
}
