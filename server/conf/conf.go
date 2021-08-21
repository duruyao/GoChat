package conf

import (
	"encoding/json"
)

type User struct {
	Uid string `json:"uid,omitempty"`
	Pwd string `json:"pwd,omitempty"`
}

type Conf struct {
	Addr             string `json:"addr,omitempty"`
	MaxUsers         int    `json:"max_users,omitempty"`
	MaxRooms         int    `json:"max_rooms,omitempty"`
	MaxUsersPerRoom  int    `json:"max_users_per_room,omitempty"`
	MaxRoomsPerAdmin int    `json:"max_rooms_per_admin,omitempty"`
	DbFileEnable     bool   `json:"db_file_enable"`
	LogFileEnable    bool   `json:"log_file_enable"`
	Admins           []User `json:"admins,omitempty"`
}

var defaultConf = Conf{
	Addr:             "localhost:1213",
	MaxUsers:         20000,
	MaxRooms:         200,
	MaxUsersPerRoom:  100,
	MaxRoomsPerAdmin: 10,
	DbFileEnable:     true,
	LogFileEnable:    true,
	Admins:           []User{{Uid: "root", Pwd: "971213"}},
}

// String returns type string in JSON format.
func (c *Conf) String() string {
	if js, err := json.MarshalIndent(c, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// String returns type []byte in JSON format.
func (c *Conf) Bytes() []byte {
	if js, err := json.MarshalIndent(c, "", "    "); err != nil {
		return []byte(err.Error())
	} else {
		return js
	}
}

// Parse parses type Conf from type []byte in JSON format.
func (c *Conf) Parse(js []byte) error {
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, c)
}

// AppendAdmins appends admins to c.Admins.
func (c *Conf) AppendAdmins(admins ...User) { c.Admins = append(c.Admins, admins...) }

// DefaultConf returns a type Conf with default value.
func DefaultConf() Conf { return defaultConf }
