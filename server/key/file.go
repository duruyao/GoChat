package key

import (
	"fmt"
	mlog "github.com/duruyao/gochat/server/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const fileDirFmt = "%s/.GoChat/.key"
const keyPathFmt = "%s/.server.key"
const crtPathFmt = "%s/.server.crt"

var dirOnce sync.Once
var dir string

//
func Dir() string {
	dirOnce.Do(func() {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			mlog.FatalLn(err)
		}
		dir = fmt.Sprintf(fileDirFmt, userHomeDir)
	})
	return dir
}

//
func AllPaths() [2]string {
	return [2]string{fmt.Sprintf(keyPathFmt, Dir()), fmt.Sprintf(crtPathFmt, Dir())}
}

//
func Path(name string) string {
	if "crt" == name {
		return fmt.Sprintf(crtPathFmt, Dir())
	} else if "key" == name {
		return fmt.Sprintf(keyPathFmt, Dir())
	}
	return ""
}

//
func AreExist() bool {
	return !AreNotExist()
}

//
func AreNotExist() bool {
	if _, err := os.Stat(AllPaths()[0]); os.IsNotExist(err) {
		return true
	}
	if _, err := os.Stat(AllPaths()[1]); os.IsNotExist(err) {
		return true
	}
	return false
}

//
func projectDir() string {
	dir, err := os.Getwd()
	if err != nil {
		mlog.FatalLn(err)
	}
	for {
		base := filepath.Base(dir)
		if strings.HasPrefix(base, "gochat") || strings.HasPrefix(base, "GoChat") {
			return dir
		}
		if filepath.Dir(dir) == dir {
			break
		}
		dir = filepath.Dir(dir)
	}
	mlog.FatalLn("not found project directory")
	return ""
}

//
func CreateKeys() (err error) {
	if _, e := os.Stat(Dir()); os.IsNotExist(e) {
		if err = os.MkdirAll(Dir(), os.ModePerm); err != nil {
			return err
		}
	}
	return genKeys()
}
