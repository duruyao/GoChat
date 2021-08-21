package log

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const fileDirFmt = "%s/.GoChat/log/%s"

var files = map[string]*os.File{
	"all":   nil,
	"debug": nil,
	"info":  nil,
	"error": nil,
	"fatal": nil,
}

var dirOnce sync.Once
var dir string

// Dir returns '${HOME}/.GoChat/log'.
func Dir() string {
	dirOnce.Do(func() {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		dir = fmt.Sprintf(fileDirFmt, userHomeDir, time.Now().Format("2006-01-02"))
	})
	return dir
}

//
func CreateFiles() (err error) {
	if _, e := os.Stat(Dir()); os.IsNotExist(e) {
		if err = os.MkdirAll(Dir(), os.ModePerm); err != nil {
			return err
		}
	}
	for name := range files {
		files[name], err = os.Create(Dir() + "/" + name + ".log")
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
		if files[name], err = os.OpenFile(Dir()+"/"+name+".log", os.O_WRONLY|os.O_APPEND, 0666); err != nil {
			return err
		}
	}
	return nil
}

//
func CloseFiles() error {
	for _, f := range files {
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}
