package log

import (
	"fmt"
	"os"
)

type logfunc func() string

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelStr = []string{
	"TRACE",
	"DEBUG",
	"INFO ",
	"WARN ",
	"ERROR",
	"FATAL",
}

const consoleFlags = Ldate | Ltime | Lmicroseconds | Llevel | Lfile

var (
	console  = NewLogger(os.Stderr, "Console", consoleFlags)
	logLevel = INFO
)

func SetConsoleLogLevel(level int) {
	logLevel = level
}

func SetConsolePrefix(prefix string) {
	console.SetPrefix(prefix)
}

func Trace(a ...interface{}) {
	if logLevel <= TRACE {
		console.Output(TRACE, fmt.Sprint(a...))
	}
}

func Tracef(format string, a ...interface{}) {
	if logLevel <= TRACE {
		console.Output(TRACE, fmt.Sprintf(format, a...))
	}
}

func Tracefn(fn logfunc) {
	if logLevel <= TRACE {
		console.Output(TRACE, fn())
	}
}

func Debug(a ...interface{}) {
	if logLevel <= DEBUG {
		console.Output(DEBUG, fmt.Sprint(a...))
	}
}

func Debugf(format string, a ...interface{}) {
	if logLevel <= DEBUG {
		console.Output(DEBUG, fmt.Sprintf(format, a...))
	}
}

func Debugfn(fn logfunc) {
	if logLevel <= DEBUG {
		console.Output(DEBUG, fn())
	}
}

func Info(a ...interface{}) {
	if logLevel <= INFO {
		console.Output(INFO, fmt.Sprint(a...))
	}
}

func Infof(format string, a ...interface{}) {
	if logLevel <= INFO {
		console.Output(INFO, fmt.Sprintf(format, a...))
	}
}

func Infofn(fn logfunc) {
	if logLevel <= INFO {
		console.Output(INFO, fn())
	}
}

func Warn(a ...interface{}) {
	if logLevel <= WARN {
		console.Output(WARN, fmt.Sprint(a...))
	}
}

func Warnf(format string, a ...interface{}) {
	if logLevel <= WARN {
		console.Output(WARN, fmt.Sprintf(format, a...))
	}
}

func Warnfn(fn logfunc) {
	if logLevel <= WARN {
		console.Output(WARN, fn())
	}
}

func WarnOnError(err error) {
	if logLevel <= WARN && err != nil {
		console.Output(WARN, fmt.Sprint(err))
	}
}

func Error(a ...interface{}) {
	if logLevel <= ERROR {
		console.Output(ERROR, fmt.Sprint(a...))
	}
}

func Errorf(format string, a ...interface{}) {
	if logLevel <= ERROR {
		console.Output(ERROR, fmt.Sprintf(format, a...))
	}
}

func Errorfn(fn logfunc) {
	if logLevel <= ERROR {
		console.Output(ERROR, fn())
	}
}

func ErrorOnError(err error) {
	if logLevel <= ERROR && err != nil {
		console.Output(ERROR, fmt.Sprint(err))
	}
}

func Fatal(a ...interface{}) {
	console.Output(FATAL, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	console.Output(FATAL, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func Fatalfn(fn logfunc) {
	console.Output(FATAL, fn())
	os.Exit(1)
}

func FatalOnError(err error) {
	if err != nil {
		console.Output(FATAL, fmt.Sprint(err))
		os.Exit(1)
	}
}
