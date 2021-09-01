package conf

import (
	"fmt"
	"github.com/duruyao/gochat/server/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Path returns '$HOME/.GoChat/conf/gochat.conf'.
func Path() string { return fmt.Sprintf("%s/.GoChat/conf/gochat.conf", util.UserHomeDir()) }

// Dir returns '$HOME/.GoChat/conf'.
func Dir() string { return filepath.Dir(Path()) }

// IsNotExist returns true if the file '$HOME/.GoChat/conf/gochat.conf' doesn't exists, otherwise false.
func IsNotExist() bool {
	if _, err := os.Stat(Path()); os.IsNotExist(err) {
		return true
	}
	return false
}

// createFile creates a new path '$HOME/.GoChat/conf/gochat.conf'.
func createFile() (err error) {
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

// readFile reads a type Config from the file '$HOME/.GoChat/conf/gochat.conf' to conf.
func readFile(c *config) error {
	data, err := ioutil.ReadFile(Path())
	if err != nil {
		return err
	}
	return c.Parse(data)
}

// writeFile writes a type Config from conf to the new file '$HOME/.GoChat/conf/gochat.conf'.
func writeFile(c *config) error {
	return ioutil.WriteFile(Path(), []byte(c.String()), 0666)
}
