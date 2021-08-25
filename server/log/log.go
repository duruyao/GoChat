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
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
	panicLogger *log.Logger
	fatalLogger *log.Logger
)

var refreshLoggerOnce sync.Once

func RefreshLogger() {
	debugLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["debug"]), "[debug] ", flags)
	infoLogger = log.New(io.MultiWriter(os.Stdout, files["all"], files["info"]), "[info ] ", flags)
	errorLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["error"]), "[error] ", flags)
	panicLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["panic"]), "[panic] ", flags)
	fatalLogger = log.New(io.MultiWriter(os.Stderr, files["all"], files["fatal"]), "[fatal] ", flags)
}

func Debug(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = debugLogger.Output(2, fmt.Sprint(v...))
}

func DebugF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = debugLogger.Output(2, fmt.Sprintf(format, v...))
}

func DebugLn(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = debugLogger.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = infoLogger.Output(2, fmt.Sprint(v...))
}

func InfoF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func InfoLn(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = infoLogger.Output(2, fmt.Sprintln(v...))
}

func Panic(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	s := fmt.Sprint(v...)
	_ = panicLogger.Output(2, s)
	panic(s)
}

func PanicF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	s := fmt.Sprintf(format, v...)
	_ = panicLogger.Output(2, s)
	panic(s)
}

func PanicLn(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	s := fmt.Sprintln(v...)
	_ = panicLogger.Output(2, s)
	panic(s)
}

func Error(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = errorLogger.Output(2, fmt.Sprint(v...))
}

func ErrorF(format string, v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = errorLogger.Output(2, fmt.Sprintf(format, v...))
}

func ErrorLn(v ...interface{}) {
	refreshLoggerOnce.Do(RefreshLogger)
	_ = errorLogger.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	_ = fatalLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func FatalF(format string, v ...interface{}) {
	_ = fatalLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func FatalLn(v ...interface{}) {
	_ = fatalLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
