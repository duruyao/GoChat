package log

import (
	"log"
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
