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
	projectDb   = ".GoChat/db"
	projectLog  = ".GoChat/log"
	projectConf = ".GoChat/gochat.conf"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initDb()
	initLog()
	initConf() // make sure to call initConf() last
}

func ProjectDbDir() string { return UserHomeDir + "/" + projectDb }

func ProjectLogDir() string { return UserHomeDir + "/" + projectLog }

func ProjectConfDir() string { return UserHomeDir + "/" + projectConf }
