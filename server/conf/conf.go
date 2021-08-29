package conf

import (
	"encoding/json"
	"sync"
)

type User struct {
	Uid string `json:"uid,omitempty"`
	Pwd string `json:"pwd,omitempty"`
}

type config struct {
	Addr             string `json:"addr,omitempty"`
	MaxUsers         int    `json:"max_users,omitempty"`
	MaxRooms         int    `json:"max_rooms,omitempty"`
	MaxUsersPerRoom  int    `json:"max_users_per_room,omitempty"`
	MaxRoomsPerAdmin int    `json:"max_rooms_per_admin,omitempty"`
	HttpsEnable      bool   `json:"https_enable"`
	LogFileEnable    bool   `json:"log_file_enable"`
}

type Config struct {
	rwMu sync.RWMutex
	cfg  *config
}

// DefaultConf returns a type Config without default value.
func NewConfig() *Config {
	return &Config{
		cfg: &config{},
	}
}

// DefaultConf returns a type Config with default value.
func NewDefaultConfig() *Config {
	return &Config{
		cfg: &config{
			Addr:             "localhost:1213",
			MaxUsers:         20000,
			MaxRooms:         200,
			MaxUsersPerRoom:  100,
			MaxRoomsPerAdmin: 10,
			HttpsEnable:      true,
			LogFileEnable:    true,
		},
	}
}

//
func (c *Config) Addr() string {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.Addr
}

//
func (c *Config) MaxUsers() int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.MaxUsers
}

//
func (c *Config) MaxRooms() int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.MaxRooms
}

//
func (c *Config) MaxUsersPreRoom() int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.MaxUsersPerRoom
}

//
func (c *Config) MaxRoomsPerAdmin() int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.MaxRoomsPerAdmin
}

//
func (c *Config) LogFileEnable() bool {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.LogFileEnable
}

//
func (c *Config) HttpsEnable() bool {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.cfg.HttpsEnable
}

//
func (c *Config) SetAddr(addr string) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.Addr = addr
}

//
func (c *Config) SetMaxUsers(maxUsers int) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.MaxUsers = maxUsers
}
//
func (c *Config) SetMaxRooms(maxRooms int) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.MaxRooms = maxRooms
}

//
func (c *Config) SetMaxUsersPreRoom(maxUsersPreRoom int) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.MaxUsersPerRoom = maxUsersPreRoom
}

//
func (c *Config) SetMaxRoomsPerAdmin(maxRoomsPerAdmin int) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.MaxRoomsPerAdmin = maxRoomsPerAdmin
}

//
func (c *Config) SetLogFileEnable(logFileEnable bool) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.LogFileEnable = logFileEnable
}

//
func (c *Config) SetHttpsEnable(httpsEnable bool) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.cfg.HttpsEnable = httpsEnable
}

// String returns JSON format string.
func (c *Config) String() string {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	if js, err := json.MarshalIndent(c.cfg, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Serialize returns JSON format string without indent.
func (c *Config) Serialize() string {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	if js, err := json.Marshal(c.cfg); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Parse parses type Config from type []byte in JSON format.
func (c *Config) Parse(js []byte) error {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, c.cfg)
}
