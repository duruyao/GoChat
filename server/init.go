package server

import (
	"github.com/duruyao/gochat/server/conf"
	"github.com/duruyao/gochat/server/db"
	"github.com/duruyao/gochat/server/key"
	mlog "github.com/duruyao/gochat/server/log"
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
			mlog.FatalLn(err)
		}
		if err := conf.WriteFile(MyConfig); err != nil {
			mlog.FatalLn(err)
		}
	} else {
		if err := conf.ReadFile(MyConfig); err != nil {
			mlog.FatalLn(err)
		}
	}

	// init log files
	if MyConfig.LogFileEnable() {
		if mlog.AreNotExist() {
			if err := mlog.CreateFiles(); err != nil {
				mlog.FatalLn(err)
			}
		} else {
			if err := mlog.OpenFiles(); err != nil {
				mlog.FatalLn(err)
			}
		}
		BeforeQuit.Append(mlog.CloseFiles)
		go GoRefreshLogFiles()
	}

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

	// init https keys
	if MyConfig.HttpsEnable() {
		if key.AreNotExist() {
			if err := key.CreateKeys(); err != nil {
				mlog.FatalLn(err)
			}
		}
	}
}
