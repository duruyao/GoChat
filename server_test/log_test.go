package server_test

import (
	"github.com/duruyao/gochat/server"
	mlog "github.com/duruyao/gochat/server/log"
	"testing"
	"time"
)

func TestLogCreate(t *testing.T) {
	if err := mlog.CreateFiles(); err != nil {
		t.Fatal(err)
	}
	Debug := mlog.DebugLogger().Println
	DebugF := mlog.DebugLogger().Printf
	Info := mlog.InfoLogger().Println
	InfoF := mlog.InfoLogger().Printf
	Error := mlog.ErrorLogger().Println
	ErrorF := mlog.ErrorLogger().Printf
	Fatal := mlog.FatalLogger().Println
	FatalF := mlog.FatalLogger().Printf
	defer func() {
		if err := mlog.CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	Debug("i am Debug()")
	DebugF("i am %s\n", "DebugF()")
	Info("i am Info()")
	InfoF("i am %s\n", "InfoF()")
	Error("i am Error()")
	ErrorF("i am %s\n", "ErrorF()")
	Fatal("i am Fatal()")
	FatalF("i am %s\n", "FatalF()")
}

func TestLogOpen(t *testing.T) {
	if err := mlog.OpenFiles(); err != nil {
		t.Fatal(err)
	}
	Debug := mlog.DebugLogger().Println
	DebugF := mlog.DebugLogger().Printf
	Info := mlog.InfoLogger().Println
	InfoF := mlog.InfoLogger().Printf
	Error := mlog.ErrorLogger().Println
	ErrorF := mlog.ErrorLogger().Printf
	Fatal := mlog.FatalLogger().Println
	FatalF := mlog.FatalLogger().Printf
	defer func() {
		if err := mlog.CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	Debug("i am a new Debug()")
	DebugF("i am a new %s\n", "DebugF()")
	Info("i am a new Info()")
	InfoF("i am a new %s\n", "InfoF()")
	Error("i am a new Error()")
	ErrorF("i am a new %s\n", "ErrorF()")
	Fatal("i am a new Fatal()")
	FatalF("i am a new %s\n", "FatalF()")
}

func TestWriteNewLogFiles(t *testing.T) {
	if err := mlog.OpenFiles(); err != nil {
		t.Fatal(err)
	}
	Debug := mlog.DebugLogger().Println
	Info := mlog.InfoLogger().Println
	Error := mlog.ErrorLogger().Println
	Fatal := mlog.FatalLogger().Println
	defer func() {
		server.BeforeQuit.Do()
		server.WantQuit()
	}()
loop:
	for {
		select {
		case <-time.After(time.Minute):
			timeStr := time.Now().Format("2006-01-02 03:04:05")
			Debug("Debug() shows current time: " + timeStr)
			Info("Info() shows current time: " + timeStr)
			Error("Error() shows current time: " + timeStr)
			Fatal("Fatal() shows current time: " + timeStr)
		case <-time.After(7 * time.Hour):
			break loop
		}
	}

	//duration := time.Second
	//select {
	//case <-time.After(duration):
	//	t.Log(time.Now().Format("2006-01-02 03:04:05"))
	//}
	//for {
	//	select {
	//	case <-time.After(4 * duration):
	//		t.Log(time.Now().Format("2006-01-02 03:04:05"))
	//	}
	//}
}
