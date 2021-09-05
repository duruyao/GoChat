package conf

import (
	"fmt"
	"github.com/duruyao/gochat/server/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Path() string {
	return fmt.Sprintf("%s/app/gochat.conf", util.GoChatDir())
}

func Dir() string {
	return filepath.Dir(Path())
}

func IsNotExist() bool {
	if _, err := os.Stat(Path()); os.IsNotExist(err) {
		return true
	}
	return false
}

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

func readFile(c *config) error {
	data, err := ioutil.ReadFile(Path())
	if err != nil {
		return err
	}
	return c.Parse(data)
}

func writeFile(c *config) error {
	return ioutil.WriteFile(Path(), []byte(c.String()), 0666)
}
