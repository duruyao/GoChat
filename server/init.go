package server

import (
	"github.com/duruyao/gochat/server/conf"
	"github.com/duruyao/gochat/server/db"
	mlog "github.com/duruyao/gochat/server/log"
	"log"
)

var (
	MyConfig      = conf.NewDefaultConfig()
	MyRoomsTable  = new(db.RoomsTable)
	MyAdminsTable = new(db.AdminsTable)
)

func init() {
	// init configuration file
	if conf.IsNotExist() {
		if err := conf.CreateFile(); err != nil {
			log.Fatalln(err)
		}
		if err := conf.WriteFile(MyConfig); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := conf.ReadFile(MyConfig); err != nil {
			log.Fatalln(err)
		}
	}

	// init log files
	if MyConfig.LogFileEnable() {
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
	go GoRefreshLogFiles()

	// init database file
	if db.IsNotExist() {
		if err := db.CreateDB(); err != nil {
			mlog.FatalLn(err)
		}
	}
	if err := db.OpenDB(); err != nil {
		mlog.FatalLn(err)
	}
	BeforeQuit.Append(db.CloseDB)
}
