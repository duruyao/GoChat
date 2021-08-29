package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const fileDirFmt = "%s/.GoChat/log/%s"

var files = map[string]*os.File{ //
	"all":   nil,
	"debug": nil,
	"info":  nil,
	"error": nil,
	"panic": nil,
	"fatal": nil,
}

var userHomeDirOnce sync.Once
var userHomeDir string

//
func homeDir() string {
	userHomeDirOnce.Do(func() {
		var err error
		userHomeDir, err = os.UserHomeDir()
		if err != nil {
			FatalLn(err)
		}
	})
	return userHomeDir
}

// Dir returns '${HOME}/.GoChat/log'.
func Dir() string {
	return fmt.Sprintf(fileDirFmt, homeDir(), time.Now().Format("2006-01-02"))
}

//
func Path(name string) string {
	return Dir() + "/" + name + ".log"
}

//
func AreExist() bool {
	return !AreNotExist()
}

//
func AreNotExist() bool {
	for name := range files {
		if _, err := os.Stat(Path(name)); os.IsNotExist(err) {
			return true
		}
	}
	return false
}

//
func CreateFiles() (err error) {
	if _, e := os.Stat(Dir()); os.IsNotExist(e) {
		if err = os.MkdirAll(Dir(), os.ModePerm); err != nil {
			return err
		}
	}
	for name := range files {
		files[name], err = os.Create(Path(name))
		if err != nil {
			return err
		}
	}
	return err
}

//
func OpenFiles() (err error) {
	for name, f := range files {
		if nil != f {
			continue
		}
		if files[name], err = os.OpenFile(Path(name), os.O_WRONLY|os.O_APPEND, 0666); err != nil {
			return err
		}
	}
	return nil
}

//
func CloseFiles() error {
	for _, f := range files {
		if f != nil {
			if err := f.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
