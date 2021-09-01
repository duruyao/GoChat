package log

import (
	"testing"
	"time"
)

func TestCreateFiles(t *testing.T) {
	if err := CreateFiles(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	DebugLn("i am DebugLn()")
	DebugF("i am %s\n", "DebugF()")
	InfoLn("i am InfoLn()")
	InfoF("i am %s\n", "InfoF()")
	ErrorLn("i am ErrorLn()")
	ErrorF("i am %s\n", "ErrorF()")
}

func TestOpenFiles(t *testing.T) {
	if err := OpenFiles(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseFiles(); err != nil {
			t.Fatal(err)
		}
	}()
	DebugLn("i am a new DebugLn()")
	DebugF("i am a new %s\n", "DebugF()")
	InfoLn("i am a new InfoLn()")
	InfoF("i am a new %s\n", "InfoF()")
	ErrorLn("i am a new ErrorLn()")
	ErrorF("i am a new %s\n", "ErrorF()")
}

func TestInit(t *testing.T) {
	quit := make(chan struct{})
	go func() {
		select {
		case <-time.After(1 * time.Hour):
			close(quit)
		}
	}()
loop:
	for {
		select {
		case <-time.After(time.Minute):
			timeStr := time.Now().Format("2006-01-02 03:04:05")
			DebugLn("Current time: " + timeStr)
		case <-quit:
			break loop
		}
	}
}
