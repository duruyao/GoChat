package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"log"
	"sync"
	"time"
)

// functions for printing logs
var (
	Debug func(...interface{})
	Info  func(...interface{})
	Error func(...interface{})
	Fatal func(...interface{})

	DebugF func(string, ...interface{})
	InfoF  func(string, ...interface{})
	ErrorF func(string, ...interface{})
	FatalF func(string, ...interface{})
)

//
func BindPrintFunc() {
	Debug = mlog.DebugLogger().Println
	Info = mlog.InfoLogger().Println
	Error = mlog.ErrorLogger().Println
	Fatal = mlog.FatalLogger().Println

	DebugF = mlog.DebugLogger().Printf
	InfoF = mlog.InfoLogger().Printf
	ErrorF = mlog.ErrorLogger().Printf
	FatalF = mlog.FatalLogger().Printf
}

var goRefreshLogFilesOnce sync.Once

//
func GoRefreshLogFiles() {
	goRefreshLogFilesOnce.Do(func() {
		duration := Tomorrow().Sub(time.Now())
		select {
		case <-time.After(duration):
			if err := mlog.CloseFiles(); err != nil {
				log.Fatalln(err)
			}
			if err := mlog.CreateFiles(); err != nil {
				log.Fatalln(err)
			}
			mlog.RefreshLogger()
			BindPrintFunc()
		}
	loop:
		for {
			select {
			case <-time.After(24 * time.Hour):
				if err := mlog.CloseFiles(); err != nil {
					log.Fatalln(err)
				}
				if err := mlog.CreateFiles(); err != nil {
					log.Fatalln(err)
				}
				mlog.RefreshLogger()
				BindPrintFunc()
			case <-Quit():
				break loop
			}
		}
	})
}
