package server

import (
	"log"
	"os"
)

// initLog creates directory '${HOME}/.GoChat/log/' if it doesn't exist.
func initLog() {
	if err := os.MkdirAll(ProjectLogDir(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
