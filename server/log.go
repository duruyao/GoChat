package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var loggerOnce sync.Once
var loggerInstance *Logger

type Logger struct {
	file *os.File
	path string
}

//
func GetLogger(fileEnable bool) *Logger {
	loggerOnce.Do(func() {
		if fileEnable {
			loggerInstance = &Logger{
				path: fmt.Sprintf(LogFileDirFmt, UserHomeDir) + "/" + time.Now().Format("2006-01-02") + ".log",
			}
			loggerInstance.OpenFile()
		}
	})
	return loggerInstance
}

//
func (l *Logger) OpenFile() {
	var err error
	l.file, err = os.OpenFile(l.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, l.file)
	log.SetOutput(mw)
	log.Printf("Open %s\n", l.path)
}

//
func (l *Logger) CloseFile() {
	log.Printf("Close %s\n", l.path)
	_ = l.file.Close()
}
