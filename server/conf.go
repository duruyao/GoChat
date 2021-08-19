package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Admin struct {
	Uid string `json:"uid,omitempty"`
	Pwd string `json:"pwd,omitempty"`
}

type Conf struct {
	Addr             string  `json:"addr,omitempty"`
	MaxUsers         int     `json:"max_users,omitempty"`
	MaxRooms         int     `json:"max_rooms,omitempty"`
	MaxUsersPerRoom  int     `json:"max_users_per_room,omitempty"`
	MaxRoomsPerAdmin int     `json:"max_rooms_per_admin,omitempty"`
	DbEnable         bool    `json:"db_enable"`
	LogEnable        bool    `json:"log_enable"`
	Admins           []Admin `json:"admins,omitempty"`
}

const confDefault = `{
  "addr": "localhost:1213",
  "max_users": 20000,
  "max_rooms": 200,
  "max_users_per_room": 100,
  "max_rooms_per_admin": 10,
  "db_enable": true,
  "log_enable": true,
  "admins": [
    {
      "uid": "root",
      "pwd": "971213"
    }
  ]
}`

// init creates file '${HOME}/.GoChat/gochat.conf' if it doesn't exist.
func initConf() {
	// mkdir ${HOME}/.GoChat/
	path := (*Conf)(nil).FilePath()
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

// NewConf returns address of type Conf with default values.
// The default and first admin in Admins is root, and root's default password is '971213'.
func NewConf() *Conf {
	return &Conf{
		Addr:             "localhost:1213",
		MaxUsers:         20000,
		MaxRooms:         200,
		MaxUsersPerRoom:  100,
		MaxRoomsPerAdmin: 10,
		DbEnable:         true,
		LogEnable:        true,
		Admins:           []Admin{{Uid: "root", Pwd: "971213"}},
	}
}

// Load reads content of configuration from confFile
func (c *Conf) Load() {
	data, err := ioutil.ReadFile(c.FilePath())
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, c); err != nil {
		log.Fatal(err)
	}
}

// Load writes content of configuration to confFile
func (c *Conf) Save() {
	file, err := os.Create(c.FilePath())
	if err != nil {
		log.Fatal(err)
	}
	_, _ = fmt.Fprintln(file, c.String())
	defer func() { _ = file.Close() }()
}

// String returns type string in Json format
func (c *Conf) String() string {
	if data, err := json.MarshalIndent(*c, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(data)
	}
}

// AddAdmin appends a type Admin to c.Admins
func (c *Conf) AddAdmin(uid string, pwd string) {
	a := Admin{
		Uid: uid,
		Pwd: pwd,
	}
	c.Admins = append(c.Admins, a)
}

// FilePath returns '${HOME}/.GoChat/gochat.conf'.
func (c *Conf) FilePath() string { return fmt.Sprintf(ConfFilePathFmt, UserHomeDir) }
