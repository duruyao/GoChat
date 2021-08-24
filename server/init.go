package server

import (
	"github.com/duruyao/gochat/server/conf"
	"github.com/duruyao/gochat/server/db"
	mlog "github.com/duruyao/gochat/server/log"
	"log"
)

var (
	Config = conf.DefaultConf()
	Table  = db.NewTable()
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
		BeforeQuit.Append(mlog.CloseFiles)
	}
	BindPrintFunc()
	go GoRefreshLogFiles()

	// init database file
	if Config.DbFileEnable {
		if db.IsNotExist() {
			if err := db.CreateFile(); err != nil {
				Fatal(err)
			}
		} else {
			if err := db.ReadFile(Table); err != nil {
				Fatal(err)
			}
		}
	}
}
