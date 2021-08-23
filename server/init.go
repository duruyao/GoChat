package server

import (
	"github.com/duruyao/gochat/server/conf"
	"github.com/duruyao/gochat/server/db"
	mlog "github.com/duruyao/gochat/server/log"
	"log"
	"os"
	"sync"
	"time"
)

var userHomeDirOnce sync.Once
var userHomeDir string

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

var (
	Config = conf.DefaultConf()

	Debug func(...interface{})
	Info  func(...interface{})
	Error func(...interface{})
	Fatal func(...interface{})

	DebugF func(string, ...interface{})
	InfoF  func(string, ...interface{})
	ErrorF func(string, ...interface{})
	FatalF func(string, ...interface{})

	Table = db.NewTable()

	Quit = func() error { return nil } // mlog.CloseFiles()
)

func init() {
	// init configuration file
	if conf.IsNotExist() {
		if err := conf.CreateFile(); err != nil {
			log.Fatalln(err)
		}
		if err := conf.WriteFile(&Config); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := conf.ReadFile(&Config); err != nil {
			log.Fatalln(err)
		}
	}

	// init log files
	if Config.LogFileEnable {
		if mlog.AreNotExist() {
			if err := mlog.CreateFiles(); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := mlog.OpenFiles(); err != nil {
				log.Fatalln(err)
			}
		}
		Quit = mlog.CloseFiles
	}
	bindLoggerFunc()
	go goLogFilesRefresher()

	// init database file
	if Config.DbFileEnable {
		if db.IsNotExist() {
			if err := db.CreateFile(); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := db.ReadFile(Table); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

//
func goLogFilesRefresher() {
	duration := tomorrow().Sub(time.Now())
	for {
		select {
		case <-time.After(duration):
			if err := mlog.CloseFiles(); err != nil {
				log.Fatalln(err)
			}
			if err := mlog.CreateFiles(); err != nil {
				log.Fatalln(err)
			}
			mlog.RefreshLogger()
			bindLoggerFunc()
			duration = tomorrow().Sub(time.Now())
		}
	}
}

//
func bindLoggerFunc() {
	Debug = mlog.DebugLogger().Println
	Info = mlog.InfoLogger().Println
	Error = mlog.ErrorLogger().Println
	Fatal = mlog.FatalLogger().Println

	DebugF = mlog.DebugLogger().Printf
	InfoF = mlog.InfoLogger().Printf
	ErrorF = mlog.ErrorLogger().Printf
	FatalF = mlog.FatalLogger().Printf
}

// return tomorrow 00:00:00.000.
func tomorrow() time.Time {
	t := time.Now().AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
