package conf

import (
	"encoding/json"
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
	DbFileEnable     bool   `json:"db_file_enable"`
	LogFileEnable    bool   `json:"log_file_enable"`
	Admins           []User `json:"admins,omitempty"`
}

type Config struct {
	// rwMu sync.RWMutex // TODO: sync.RWMutex will be added in the future, temporarily read-only
	cfg *config
}

// DefaultConf returns a type Config without default value.
func NewConfig() *Config {
	return &Config{
		cfg: &config{
			Admins: []User{},
		},
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
			DbFileEnable:     true,
			LogFileEnable:    true,
			Admins:           []User{{Uid: "root", Pwd: "971213"}},
		},
	}
}

func (c *Config) Addr() string { return c.cfg.Addr }

func (c *Config) MaxUsers() int { return c.cfg.MaxUsers }

func (c *Config) MaxRooms() int { return c.cfg.MaxRooms }

func (c *Config) MaxUsersPreRoom() int { return c.cfg.MaxUsersPerRoom }

func (c *Config) MaxRoomsPerAdmin() int { return c.cfg.MaxRoomsPerAdmin }

func (c *Config) DbFileEnable() bool { return c.cfg.DbFileEnable }

func (c *Config) LogFileEnable() bool { return c.cfg.LogFileEnable }

func (c *Config) Admins() (backup []User) {
	copy(backup, c.cfg.Admins)
	return backup
}

// String returns JSON format string.
func (c *Config) String() string {
	if js, err := json.MarshalIndent(c.cfg, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Serialize returns JSON format string without indent.
func (c *Config) Serialize() string {
	if js, err := json.Marshal(c.cfg); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Parse parses type Config from type []byte in JSON format.
func (c *Config) Parse(js []byte) error {
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, c.cfg)
}

// AppendAdmins appends admins to c.Admins.
func (c *Config) AppendAdmins(admins ...User) { c.cfg.Admins = append(c.cfg.Admins, admins...) }
