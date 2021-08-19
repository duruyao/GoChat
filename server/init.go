package server

import (
	"log"
	"os"
)

var UserHomeDir = func() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}()

const (
	DbFileDirFmt    = "%s/.GoChat/db"
	LogFileDirFmt   = "%s/.GoChat/log"
	ConfFilePathFmt = "%s/.GoChat/gochat.conf"
)

func init() {
	initConf()
	initLog()
	initDb()
}
