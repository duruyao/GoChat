package log

import (
	"fmt"
	"github.com/duruyao/gochat/server/util"
	"os"
	"time"
)

var files = map[string]*os.File{ //
	"all":   nil,
	"debug": nil,
	"info":  nil,
	"error": nil,
	"panic": nil,
	"fatal": nil,
}

// Dir returns "$HOME/.GoChat/log/yyyy-mm-dd".
func Dir() string {
	return fmt.Sprintf("%s/.GoChat/log/%s", util.UserHomeDir(), time.Now().Local().Format("2006-01-02"))
}

// Path returns "$HOME/.GoChat/log/yyyy-mm-dd/name.log".
func Path(name string) string {
	return Dir() + "/" + name + ".log"
}

func IsNotExist() bool {
	for name := range files {
		if _, err := os.Stat(Path(name)); os.IsNotExist(err) {
			return true
		}
	}
	return false
}

func createFiles() (err error) {
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

func openFiles() (err error) {
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

func closeFiles() error {
	for _, f := range files {
		if f != nil {
			if err := f.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
