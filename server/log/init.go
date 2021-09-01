package log

import (
	"github.com/duruyao/gochat/server/util"
	"log"
	"time"
)

func init() {
	if IsNotExist() {
		if err := CreateFiles(); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := OpenFiles(); err != nil {
			log.Fatalln(err)
		}
	}
	go goRefreshFiles()
}

//
func goRefreshFiles() {
	duration := util.Tomorrow().Sub(time.Now())
	select {
	// TODO: add case <-quit: CloseFiles()
	case <-time.After(duration):
		if err := CloseFiles(); err != nil {
			log.Fatalln(err)
		}
		if err := CreateFiles(); err != nil {
			log.Fatalln(err)
		}
		refreshLogger()
	}
	go goRefreshFiles()
}
