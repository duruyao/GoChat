package server

import (
	"log"
	"os"
)

type Db struct {
	Room  string `json:"room"`
	Uid   string `json:"admin_uid"`
	Token string `json:"token"`
}

// initDb creates directory '${HOME}/.GoChat/db/' if it doesn't exist.
func initDb() {
	if err := os.MkdirAll(ProjectDbDir(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
