package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	defer BeforeQuit.Do()
	defer WantQuit()
	mlog.InfoLn("Load configuration:\n" + MyConfig.String())
	mlog.InfoLn("Load database:")
	rooms, err := MyRoomsTable.Query()
	if err != nil {
		mlog.ErrorLn(err)
	}
	mlog.DebugLn(rooms)
	admins, err := MyAdminsTable.Query()
	if err != nil {
		mlog.ErrorLn(err)
	}
	mlog.DebugLn(admins)
	time.Sleep(3 * time.Second)
}

func TestGoRefreshLogFiles(t *testing.T) {
	defer func() {
		BeforeQuit.Do()
	}()
	go func() {
		select {
		case <-time.After(1 * time.Hour):
			WantQuit()
		}
	}()
loop:
	for {
		select {
		case <-time.After(time.Minute):
			timeStr := time.Now().Format("2006-01-02 03:04:05")
			mlog.DebugLn("Current time: " + timeStr)
		case <-Quit():
			break loop
		}
	}
}
