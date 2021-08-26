package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const filePathFmt = "%s/.GoChat/db/gochat.db"

var pathOnce sync.Once
var path string

// Path returns '${HOME}/.GoChat/db/gochat.db'.
func Path() string {
	pathOnce.Do(func() {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path = fmt.Sprintf(filePathFmt, userHomeDir)
	})
	return path
}

// Dir returns '${HOME}/.GoChat/db'.
func Dir() string { return filepath.Dir(Path()) }

// IsExist returns true if the file '${HOME}/.GoChat/db/gochat.db' exists, otherwise false.
func IsExist() bool { return !IsNotExist() }

// IsNotExist returns true if the file '${HOME}/.GoChat/db/gochat.db' doesn't exists, otherwise false.
func IsNotExist() bool {
	if _, err := os.Stat(Path()); os.IsNotExist(err) {
		return true
	}
	return false
}

// CreateFile creates a new path '${HOME}/.GoChat/db/gochat.db'.
func CreateFile() (err error) {
	var file *os.File
	if _, e := os.Stat(Dir()); os.IsNotExist(e) {
		if err = os.MkdirAll(Dir(), os.ModePerm); err != nil {
			return err
		}
	}
	file, err = os.Create(Path())
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()
	return err
}

// ReadFile reads a type RoomTable from the file '${HOME}/.GoChat/db/gochat.db' to conf.
func ReadFile(r *RoomTable) error {
	data, err := ioutil.ReadFile(Path())
	if err != nil {
		return err
	}
	return r.Parse(data)
}

// WriteFile writes a type RoomTable from conf to the new file '${HOME}/.GoChat/db/gochat.db'.
func WriteFile(r *RoomTable) error {
	return ioutil.WriteFile(Path(), []byte(r.String()), 0666)
}
