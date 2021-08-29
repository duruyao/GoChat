package db

import (
	"database/sql"
	"fmt"
	mlog "github.com/duruyao/gochat/server/log"
	_ "github.com/mattn/go-sqlite3"
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
			mlog.FatalLn(err)
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

// CreateDB creates a new path '${HOME}/.GoChat/db/gochat.db'.
func CreateDB() (err error) {
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
	if err = file.Close(); err != nil {
		return err
	}
	if err = OpenDB(); err != nil {
		return err
	}
	defer func() { _ = CloseDB() }()
	if err = createAdminsTable(); err != nil {
		return err
	}
	if err = createRoomsTable(); err != nil {
		return err
	}
	return err
}

//
func OpenDB() (err error) {
	db, err = sql.Open("sqlite3", Path())
	return err
}

//
func CloseDB() error {
	return db.Close()
}
