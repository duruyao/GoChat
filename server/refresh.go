package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"log"
	"sync"
	"time"
)

var goRefreshLogFilesOnce sync.Once

//
func GoRefreshLogFiles() {
	goRefreshLogFilesOnce.Do(func() {
		refreshLogFiles()
		duration := Tomorrow().Sub(time.Now())
		select {
		case <-time.After(duration):
			refreshLogFiles()
		}
	loop:
		for {
			select {
			case <-time.After(24 * time.Hour):
				refreshLogFiles()
			case <-Quit():
				break loop
			}
		}
	})
}

func refreshLogFiles() {
	if err := mlog.CloseFiles(); err != nil {
		log.Fatalln(err)
	}
	if err := mlog.CreateFiles(); err != nil {
		log.Fatalln(err)
	}
	mlog.RefreshLogger()
}
