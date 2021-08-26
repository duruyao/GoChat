package server_test

import (
	"github.com/duruyao/gochat/server"
	mlog "github.com/duruyao/gochat/server/log"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	defer server.BeforeQuit.Do()
	defer server.WantQuit()
	mlog.InfoLn("Load configuration:\n" + server.MyConfig.String())
	mlog.InfoLn("Load database:")
	rooms, err := server.MyRoomsTable.Query()
	if err != nil {
		mlog.ErrorLn(err)
	}
	mlog.DebugLn(rooms)
	admins, err := server.MyAdminsTable.Query()
	if err != nil {
		mlog.ErrorLn(err)
	}
	mlog.DebugLn(admins)
	time.Sleep(3 * time.Second)
}

func TestRefreshLogFiles(t *testing.T) {
	defer func() {
		server.BeforeQuit.Do()
	}()
	go func() {
		select {
		case <-time.After(1 * time.Hour):
			server.WantQuit()
		}
	}()
loop:
	for {
		select {
		case <-time.After(time.Minute):
			timeStr := time.Now().Format("2006-01-02 03:04:05")
			mlog.DebugLn("Current time: " + timeStr)
		case <-server.Quit():
			break loop
		}
	}
}
