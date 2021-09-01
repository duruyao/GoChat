package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Tomorrow() returns an object of type time.Time whose value is 00:00:00.000 of the next day.
func Tomorrow() time.Time {
	t := time.Now().AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

var projectDirOnce sync.Once
var projectDir string

// ProjectDir returns the directory of the current project, such as "/home/user/project/gochat*".
func ProjectDir() string {
	projectDirOnce.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		for {
			base := filepath.Base(dir)
			if strings.HasPrefix(base, "gochat") || strings.HasPrefix(base, "GoChat") {
				projectDir = dir
				return
			}
			if filepath.Dir(dir) == dir {
				break
			}
			dir = filepath.Dir(dir)
		}
		log.Fatalln("not found project directory")
	})
	return projectDir
}

var userHomeDirOnce sync.Once
var userHomeDir string

// ProjectDir returns the home directory of the current user, such as "/home/user" in Unix-like OS.
func UserHomeDir() string {
	userHomeDirOnce.Do(func() {
		var err error
		userHomeDir, err = os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
	})
	return userHomeDir
}

var quit = make(chan struct{})

//
func Quit() <-chan struct{} { return quit }

//
func IsQuit() bool {
	select {
	case <-quit:
		return true
	default:
		return false
	}
}

var quitOnce sync.Once

//
func SetQuit() { quitOnce.Do(func() { close(quit) }) }
