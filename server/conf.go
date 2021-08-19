package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
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
	DbFileEnable     bool    `json:"db_file_enable"`
	LogFileEnable    bool    `json:"log_file_enable"`
	Admins           []Admin `json:"admins,omitempty"`
}

const confDefault = `{
  "addr": "localhost:1213",
  "max_users": 20000,
  "max_rooms": 200,
  "max_users_per_room": 100,
  "max_rooms_per_admin": 10,
  "db_file_enable": true,
  "log_file_enable": true,
  "admins": [
    {
      "uid": "root",
      "pwd": "971213"
    }
  ]
}`

// Conf returns type string in JSON format
func (c *Conf) String() string {
	if data, err := json.MarshalIndent(c, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(data)
	}
}

//
func (c *Conf) AddAdmin(admins ...Admin) {
	for _, admin := range admins {
		c.Admins = append(c.Admins, admin)
	}
}

var cfgOnce sync.Once
var cfgInstance *Configurator

type Configurator struct {
	rwConf sync.RWMutex
	conf   Conf
	rwPath sync.RWMutex
	path   string
}

// GetConfigurator gets address of type Configurator with default values as return.
// The default and first admin in conf.Admins is root, and root's default password is '971213'.
func GetConfigurator() *Configurator {
	cfgOnce.Do(func() {
		cfgInstance = &Configurator{
			conf: Conf{
				Addr:             "localhost:1213",
				MaxUsers:         20000,
				MaxRooms:         200,
				MaxUsersPerRoom:  100,
				MaxRoomsPerAdmin: 10,
				DbFileEnable:     true,
				LogFileEnable:    true,
				Admins:           []Admin{{Uid: "root", Pwd: "971213"}},
			},
			path: fmt.Sprintf(ConfFilePathFmt, UserHomeDir)}
	})
	return cfgInstance
}

//
func (cfg *Configurator) String() string {
	cfg.rwConf.RLock()
	defer cfg.rwConf.RUnlock()
	ret := cfg.conf.String()
	return ret
}

// Conf gets cfg.conf as return.
func (cfg *Configurator) Conf() (conf Conf) {
	cfg.rwConf.RLock()
	defer cfg.rwConf.RUnlock()
	conf = cfg.conf
	return conf
}

// SetConf sets cfg.Conf as param conf.
func (cfg *Configurator) SetConf(conf *Conf) {
	cfg.rwConf.Lock()
	defer cfg.rwConf.Unlock()
	cfg.conf = *conf
}

// LoadConf reads content of configuration from cfg.path.
func (cfg *Configurator) LoadConf() {
	// read file
	cfg.rwPath.RLock()
	data, err := ioutil.ReadFile(cfg.path)
	if err != nil {
		log.Fatal(err)
	}
	cfg.rwPath.RUnlock()
	// update conf
	cfg.rwConf.Lock()
	defer cfg.rwConf.Unlock()
	if err := json.Unmarshal(data, &(cfg.conf)); err != nil {
		log.Fatal(err)
	}
}

// LoadConf writes content of configuration to cfg.path.
func (cfg *Configurator) SaveConf() {
	cfg.rwPath.Lock()
	defer cfg.rwPath.Unlock()
	file, err := os.Create(cfg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	_, _ = file.Write([]byte(cfg.String()))
}
