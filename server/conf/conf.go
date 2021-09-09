package conf

import (
	"encoding/json"
	"sync"
)

type config struct {
	rwMu              sync.RWMutex
	Addr              string `json:"addr,omitempty"`
	MaxUsers          int    `json:"max_users,omitempty"`
	MaxGroups         int    `json:"max_groups,omitempty"`
	MaxUsersPerGroup  int    `json:"max_users_per_group,omitempty"`
	MaxGroupsPerAdmin int    `json:"max_groups_per_admin,omitempty"`
	HttpsEnable       bool   `json:"https_enable"`
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
	Addr:              "localhost:8080",
	MaxUsers:          20000,
	MaxGroups:         200,
	MaxUsersPerGroup:  100,
	MaxGroupsPerAdmin: 10,
	HttpsEnable:       false,
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
func MaxGroups() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxGroups
}

//
func MaxUsersPreGroup() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxUsersPerGroup
}

//
func MaxGroupsPerAdmin() int {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.MaxGroupsPerAdmin
}

//
func HttpsEnable() bool {
	cfg.rwMu.RLock()
	defer cfg.rwMu.RUnlock()
	return cfg.HttpsEnable
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
func SetMaxGroups(MaxGroups int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxGroups = MaxGroups
	return writeFile(&cfg)
}

//
func SetMaxUsersPreGroup(maxUsersPreGroup int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxUsersPerGroup = maxUsersPreGroup
	return writeFile(&cfg)
}

//
func SetMaxGroupsPerAdmin(MaxGroupsPerAdmin int) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.MaxGroupsPerAdmin = MaxGroupsPerAdmin
	return writeFile(&cfg)
}

//
func SetHttpsEnable(httpsEnable bool) error {
	cfg.rwMu.Lock()
	defer cfg.rwMu.Unlock()
	cfg.HttpsEnable = httpsEnable
	return writeFile(&cfg)
}
