package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"log"
	"sync"
	"time"
)

var refreshLogFilesOnce sync.Once

//
func GoRefreshLogFiles() {
	refreshLogFilesOnce.Do(refreshLogFiles)
	duration := Tomorrow().Sub(time.Now())
	select {
	case <-time.After(duration):
		refreshLogFiles()
	case <-Quit():
		return
	}
	go GoRefreshLogFiles()
}

//
func refreshLogFiles() {
	if err := mlog.CloseFiles(); err != nil {
		log.Fatalln(err)
	}
	if err := mlog.CreateFiles(); err != nil {
		log.Fatalln(err)
	}
	mlog.RefreshLogger()
}
