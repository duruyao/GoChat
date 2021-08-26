package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const filePathFmt = "%s/.GoChat/conf/gochat.conf"

var pathOnce sync.Once
var path string

// Path returns '${HOME}/.GoChat/conf/gochat.conf'.
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

// Dir returns '${HOME}/.GoChat'.
func Dir() string { return filepath.Dir(Path()) }

// IsExist returns true if the file '${HOME}/.GoChat/conf/gochat.conf' exists, otherwise false.
func IsExist() bool {
	if _, err := os.Stat(Path()); os.IsExist(err) {
		return true
	}
	return false
}

// IsNotExist returns true if the file '${HOME}/.GoChat/conf/gochat.conf' doesn't exists, otherwise false.
func IsNotExist() bool {
	return !IsExist()
}

// CreateFile creates a new path '${HOME}/.GoChat/conf/gochat.conf'.
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
	defer func() { _ = file.Close() }()
	return err
}

// ReadFile reads a type Config from the file '${HOME}/.GoChat/conf/gochat.conf' to conf.
func ReadFile(c *Config) error {
	data, err := ioutil.ReadFile(Path())
	if err != nil {
		return err
	}
	return c.Parse(data)
}

// WriteFile writes a type Config from conf to the new file '${HOME}/.GoChat/conf/gochat.conf'.
func WriteFile(c *Config) error {
	return ioutil.WriteFile(Path(), []byte(c.String()), 0666)
}
