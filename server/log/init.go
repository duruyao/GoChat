package log

import (
	"github.com/duruyao/gochat/server/util"
	"log"
	"time"
)

func init() {
	if IsNotExist() {
		if err := createFiles(); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := openFiles(); err != nil {
			log.Fatalln(err)
		}
	}
	go goRefreshFiles()
}

//
func goRefreshFiles() {
	duration := util.Tomorrow().Sub(time.Now().Local())
	select {
	case <-util.Quit():
		if err := closeFiles(); err != nil {
			log.Println(err)
		}
	case <-time.After(duration):
		if err := closeFiles(); err != nil {
			log.Println(err)
		}
		if err := createFiles(); err != nil {
			log.Println(err)
		}
		refreshLogger()
	}
	go goRefreshFiles()
}
