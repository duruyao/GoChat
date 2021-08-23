package log

import (
	"io"
	"log"
	"os"
	"sync"
)

const flags = log.Ltime | log.Llongfile

var (
	debugLoggerOnce sync.Once
	debugLogger     *log.Logger
)

func DebugLogger() *log.Logger {
	debugLoggerOnce.Do(func() {
		debugLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["debug"]), "[debug] ", flags)
	})
	return debugLogger
}

var (
	infoLoggerOnce sync.Once
	infoLogger     *log.Logger
)

func InfoLogger() *log.Logger {
	infoLoggerOnce.Do(func() {
		infoLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["info"]), "[info ] ", flags)
	})
	return infoLogger
}

var (
	errorLoggerOnce sync.Once
	errorLogger     *log.Logger
)

func ErrorLogger() *log.Logger {
	errorLoggerOnce.Do(func() {
		errorLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["error"]), "[error] ", flags)
	})
	return errorLogger
}

var (
	fatalLoggerOnce sync.Once
	fatalLogger     *log.Logger
)

func FatalLogger() *log.Logger {
	fatalLoggerOnce.Do(func() {
		fatalLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["fatal"]), "[fatal] ", flags)
	})
	return fatalLogger
}

func RefreshLogger() {
	debugLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["debug"]), "[debug] ", flags)
	infoLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["info"]), "[info ] ", flags)
	errorLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["error"]), "[error] ", flags)
	fatalLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["fatal"]), "[fatal] ", flags)
}
