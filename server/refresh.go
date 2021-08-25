package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"log"
	"time"
)

//
func GoRefreshLogFiles() {
	refreshLogFiles()
	duration := Tomorrow().Sub(time.Now())
	select {
	case <-time.After(duration):
		refreshLogFiles()
	case <-Quit():
		return
	}
	go GoRefreshLogFiles()
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
