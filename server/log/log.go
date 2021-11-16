package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const flags = log.Ltime | log.Llongfile

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	PanicLogger *log.Logger
	FatalLogger *log.Logger
)

var refreshLoggerOnce sync.Once

func refreshLogger() {
	DebugLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["debug"]), "[debug] ", flags)
	InfoLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["info"]), "[info ] ", flags)
	ErrorLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["error"]), "[error] ", flags)
	PanicLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["panic"]), "[panic] ", flags)
	FatalLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["fatal"]), "[fatal] ", flags)
}

func Debug(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = DebugLogger.Output(2, fmt.Sprint(v...))
}

func DebugF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = DebugLogger.Output(2, fmt.Sprintf(format, v...))
}

func DebugLn(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = DebugLogger.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = InfoLogger.Output(2, fmt.Sprint(v...))
}

func InfoF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

func InfoLn(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = InfoLogger.Output(2, fmt.Sprintln(v...))
}

func Panic(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	s := fmt.Sprint(v...)
	_ = PanicLogger.Output(2, s)
	panic(s)
}

func PanicF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	s := fmt.Sprintf(format, v...)
	_ = PanicLogger.Output(2, s)
	panic(s)
}

func PanicLn(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	s := fmt.Sprintln(v...)
	_ = PanicLogger.Output(2, s)
	panic(s)
}

func Error(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = ErrorLogger.Output(2, fmt.Sprint(v...))
}

func ErrorF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = ErrorLogger.Output(2, fmt.Sprintf(format, v...))
}

func ErrorLn(v ...interface{}) {
	refreshLoggerOnce.Do(refreshLogger)
	_ = ErrorLogger.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	_ = FatalLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func FatalF(format string, v ...interface{}) {
	_ = FatalLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func FatalLn(v ...interface{}) {
	_ = FatalLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
