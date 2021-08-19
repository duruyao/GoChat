package server

import (
	"fmt"
	"log"
	"os"
)

type Log struct {
}

// initLog creates directory '${HOME}/.GoChat/log' if it doesn't exist.
func initLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := os.MkdirAll((*Log)(nil).FileDir(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// File returns '${HOME}/.GoChat/log'.
func (l *Log) FileDir() string { return fmt.Sprintf(LogFileDirFmt, UserHomeDir) }
