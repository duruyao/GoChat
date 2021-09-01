package conf

import (
	"encoding/json"
	"sync"
)

type config struct {
	rwMu             sync.RWMutex
	Addr             string `json:"addr,omitempty"`
	MaxUsers         int    `json:"max_users,omitempty"`
	MaxRooms         int    `json:"max_rooms,omitempty"`
	MaxUsersPerRoom  int    `json:"max_users_per_room,omitempty"`
	MaxRoomsPerAdmin int    `json:"max_rooms_per_admin,omitempty"`
}

// String returns JSON format string.
func (c *config) String() string {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	if js, err := json.MarshalIndent(c, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Serialize returns JSON format string without indent.
func (c *config) Serialize() string {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	if js, err := json.Marshal(c); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Parse parses type Config from type []byte in JSON format.
func (c *config) Parse(js []byte) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, c)
}

var cfg = config{
	Addr:             "localhost:8080",
	MaxUsers:         20000,
	MaxRooms:         200,
	MaxUsersPerRoom:  100,
	MaxRoomsPerAdmin: 10,
}

//
func Addr() string {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.Addr
}

//
func MaxUsers() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxUsers
}

//
func MaxRooms() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxRooms
}

//
func MaxUsersPreRoom() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxUsersPerRoom
}

//
func MaxRoomsPerAdmin() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxRoomsPerAdmin
}

//
func SetAddr(addr string) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.Addr = addr
	return writeFile(&cfg)
}

//
func SetMaxUsers(maxUsers int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxUsers = maxUsers
	return writeFile(&cfg)
}

//
func SetMaxRooms(maxRooms int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxRooms = maxRooms
	return writeFile(&cfg)
}

//
func SetMaxUsersPreRoom(maxUsersPreRoom int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxUsersPerRoom = maxUsersPreRoom
	return writeFile(&cfg)
}

//
func SetMaxRoomsPerAdmin(maxRoomsPerAdmin int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxRoomsPerAdmin = maxRoomsPerAdmin
	return writeFile(&cfg)
}
