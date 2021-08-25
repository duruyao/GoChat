package log_test

import (
	mlog "github.com/duruyao/gochat/server/log"
	"testing"
)

func TestLogCreate(t *testing.T) {
	if err := mlog.CreateFiles(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mlog.CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	mlog.DebugLn("i am DebugLn()")
	mlog.DebugF("i am %s\n", "DebugF()")
	mlog.InfoLn("i am InfoLn()")
	mlog.InfoF("i am %s\n", "InfoF()")
	mlog.ErrorLn("i am ErrorLn()")
	mlog.ErrorF("i am %s\n", "ErrorF()")
	//mlog.FatalLn("i am FatalLn()")
	//mlog.FatalF("i am %s\n", "FatalF()")
}

func TestLogOpen(t *testing.T) {
	if err := mlog.OpenFiles(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mlog.CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	mlog.DebugLn("i am a new DebugLn()")
	mlog.DebugF("i am a new %s\n", "DebugF()")
	mlog.InfoLn("i am a new InfoLn()")
	mlog.InfoF("i am a new %s\n", "InfoF()")
	mlog.ErrorLn("i am a new ErrorLn()")
	mlog.ErrorF("i am a new %s\n", "ErrorF()")
	//mlog.FatalLn("i am a new FatalLn()")
	//mlog.FatalF("i am a new %s\n", "FatalF()")
}
