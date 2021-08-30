package log

import "testing"

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
