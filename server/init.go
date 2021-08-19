package server

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	initConfPath()
	initLogDir()
	initDbDir()
}

// initConfPath creates file '${HOME}/.GoChat/gochat.conf' if it doesn't exist.
func initConfPath() {
	// mkdir ${HOME}/.GoChat/
	path := fmt.Sprintf(ConfFilePathFmt, UserHomeDir)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// echo ${confDefault} > ${HOME}/.GoChat/gochat.conf
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = file.Write([]byte(confDefault))
		defer func() { _ = file.Close() }()
	}
}

// initDbDir creates directory '${HOME}/.GoChat/db' if it doesn't exist.
func initDbDir() {
	dir := fmt.Sprintf(DbFileDirFmt, UserHomeDir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// initLogDir creates directory '${HOME}/.GoChat/log' if it doesn't exist.
func initLogDir() {
	dir := fmt.Sprintf(LogFileDirFmt, UserHomeDir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
