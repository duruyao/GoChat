package server

import (
	"fmt"
	"log"
	"os"
)

type Db struct {
	Room  string `json:"room"`
	Uid   string `json:"admin_uid"`
	Token string `json:"token"`
}

// initDb creates directory '${HOME}/.GoChat/db' if it doesn't exist.
func initDb() {
	if err := os.MkdirAll((*Db)(nil).FileDir(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// FileDir returns '${HOME}/.GoChat/db'
func (d *Db) FileDir() string { return fmt.Sprintf(DbFileDirFmt, UserHomeDir) }
